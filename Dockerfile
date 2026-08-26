# Stage 1: Build
FROM golang:1.26-alpine AS builder

COPY .env /.env

RUN apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/polaris/main.go

# Stage 2: Final image
FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /main .
COPY --from=builder /.env .

EXPOSE 8080

CMD ["./main"]