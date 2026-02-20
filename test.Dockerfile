FROM golang:1.24-alpine

WORKDIR /app

ENV CGO_ENABLED=0

RUN apk add --no-cache python3 py3-pip

RUN wget "https://github.com/yt-dlp/yt-dlp/releases/download/2026.02.04/yt-dlp" -O /usr/local/bin/yt-dlp
RUN chmod a+rx /usr/local/bin/yt-dlp

RUN yt-dlp --version