FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o ./bin/monitorer ./cmd/monitorer
RUN chmod +x ./bin/monitorer

RUN go build -o ./bin/archiver ./cmd/archiver
RUN chmod +x ./bin/archiver

RUN go build -o ./bin/server ./cmd/server
RUN chmod +x ./bin/server

RUN go build -o ./bin/verifier ./cmd/verifier
RUN chmod +x ./bin/verifier

FROM alpine:3.23

COPY --from=builder /app/bin/monitorer ./monitorer

COPY --from=builder /app/bin/archiver ./archiver

COPY --from=builder /app/bin/server ./server

COPY --from=builder /app/bin/verifier ./verifier