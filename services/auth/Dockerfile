# Download libraries
FROM alpine as downloader

RUN apk add wget tar

WORKDIR /tmp

RUN wget https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz
RUN tar -xvzf litestream-v0.3.8-linux-amd64-static.tar.gz

RUN wget https://github.com/pressly/goose/releases/download/v3.5.3/goose_linux_x86_64 && mv goose_linux_x86_64 goose && chmod +x goose

# Build binary
FROM golang:1.18 as builder

WORKDIR /app

COPY . /app

RUN go build \
  -tags netgo \
  -ldflags '-w -extldflags "-static"' \
  -o /app/main \
  main.go

# Final image
FROM alpine

RUN apk add bash
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=downloader /tmp/litestream /bin/litestream
COPY --from=downloader /tmp/goose /bin/goose
COPY ./etc/docker-entrypoint.sh /etc/docker-entrypoint.sh
COPY ./etc/litestream.yml /etc/litestream.yml
COPY ./migrations/ /app/migrations/

EXPOSE 8080

ENTRYPOINT [ "/bin/bash", "/etc/docker-entrypoint.sh" ]
