FROM golang:1.12.7-alpine
RUN apk add --update --no-cache ca-certificates git && \
  apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

WORKDIR /work
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o app
ENTRYPOINT ./app