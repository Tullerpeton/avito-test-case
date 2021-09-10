FROM golang:1.15 as balance-service-build
WORKDIR /project
COPY go.mod .
RUN go mod download
COPY . /project
RUN go build -o bin/balance_service -v ./cmd/balance_service

FROM ubuntu:latest as balance-service
RUN apt update && apt install ca-certificates -y && rm -rf /var/cache/apt/*
COPY --from=balance-service-build /project/bin/balance_service /
CMD ["./balance_service"]
