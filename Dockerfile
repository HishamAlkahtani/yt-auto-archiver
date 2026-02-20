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

RUN go build -o ./bin/migrate ./cmd/migrate
RUN chmod +x ./bin/migrate


FROM alpine:3.23

# TODO: While all the other dependencies are optional, ffmpeg, ffprobe, yt-dlp-ejs and a supported JavaScript runtime/engine are highly recommended

RUN apk add --no-cache python3 py3-pip



RUN wget "https://github.com/yt-dlp/yt-dlp/releases/download/2026.02.04/yt-dlp" -O /usr/local/bin/yt-dlp
RUN chmod a+rx /usr/local/bin/yt-dlp

RUN yt-dlp --version

COPY --from=builder /app/bin/monitorer ./monitorer

COPY --from=builder /app/bin/archiver ./archiver

COPY --from=builder /app/bin/server ./server

COPY --from=builder /app/bin/verifier ./verifier

COPY --from=builder /app/bin/migrate ./migrate