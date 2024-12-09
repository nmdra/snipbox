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

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/web

FROM scratch AS production

LABEL org.opencontainers.image.title="Snipbox"
LABEL org.opencontainers.image.description="A lightweight Go application."
LABEL org.opencontainers.image.authors="Nimendra <nimendraonline@gmail.com>"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/nmdra/snipbox"

COPY --from=builder /build/main ./

CMD ["./main"]