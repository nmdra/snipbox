FROM golang:alpine AS base

# Development stage 
FROM base AS development 

WORKDIR /app

# Install the air CLI for auto-reloading
RUN go install github.com/air-verse/air@latest

COPY go.mod /app/

RUN go mod download 

ENV PORT=3000

EXPOSE ${PORT}

CMD ["air", "-c", ".air.toml"]