FROM golang:1.16 as builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG SERVICE_NAME=adder
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
COPY common common
COPY microservices/$SERVICE_NAME microservices/$SERVICE_NAME
RUN go build -o /go/bin/service -ldflags '-s -w' microservices/$SERVICE_NAME/cmd/service/main.go
RUN go build -o /go/bin/gateway -ldflags '-s -w' microservices/$SERVICE_NAME/cmd/gateway/main.go


FROM scratch as runner

COPY --from=builder /go/bin/service /service
COPY --from=builder /go/bin/gateway /gateway

CMD ["/service"]
