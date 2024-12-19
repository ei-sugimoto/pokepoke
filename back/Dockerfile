FROM golang:1.23.4-alpine AS base

FROM base AS development
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]