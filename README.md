# distributed-calculator
Inspired  by https://github.com/dapr/quickstarts/tree/master/distributed-calculator

## Init
```bash
docker-compose up -d --build
docker-compsoe exec dev make kind
```

## Run applications
```bash
docker-compsoe exec dev make dev
```

## Request http/gRPC API
```bash
docker-compose exec dev make http grpc
```

## For developer
```bash
docker-compose exec dev make
```
