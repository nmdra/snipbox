FROM golang:alpine AS base

# Development stage 
FROM base AS development 

WORKDIR /app

# Install the air CLI for auto-reloading
RUN go install github.com/air-verse/air@latest

COPY go.mod /app/

RUN go mod download 

CMD ["air", "-c", ".air.toml"]

# Buider Stage

FROM base AS builder

WORKDIR /build 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/web

FROM scratch AS production

COPY --from=builder /build/main ./

CMD ["./main"]