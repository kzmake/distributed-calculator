package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kzmake/distributed-calculator/common/health"
	grpc_zerolog "github.com/philip-bui/grpc-zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"

	pb "github.com/kzmake/distributed-calculator/microservices/adder/api/adder/v1"
	"github.com/kzmake/distributed-calculator/microservices/adder/handler"
)

const (
	serviceAddress = "0.0.0.0:4000"
	healthAddress  = "0.0.0.0:5000"
)

func init() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}

func newGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zerolog.NewUnaryServerInterceptorWithLogger(&log.Logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterAdderServer(s, handler.NewAdder())

	return s
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	grpcS := newGRPCServer()
	g.Go(func() error {
		lis, err := net.Listen("tcp", serviceAddress)
		if err != nil {
			return xerrors.Errorf("failed to listen: %w", err)
		}

		health.Serving()

		return grpcS.Serve(lis)
	})

	healthS := health.NewHealthServer()
	g.Go(func() error {
		lis, err := net.Listen("tcp", healthAddress)
		if err != nil {
			return xerrors.Errorf("failed to listen: %w", err)
		}

		return healthS.Serve(lis)
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	cancel()

	log.Info().Msg("Shutting down server...")

	_, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	grpcS.GracefulStop()

	if err := g.Wait(); err != nil {
		return xerrors.Errorf("failed to shutdown: %w", err)
	}

	log.Info().Msgf("Server exiting")

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Msgf("Failed to run server: %v", err)
	}
}
