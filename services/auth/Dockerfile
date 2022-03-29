FROM golang:1.18 as builder

WORKDIR $GOPATH/src/auth/
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  go build -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo,sqlite_omit_load_extension -o /go/bin/auth main.go

FROM alpine

COPY --from=builder /go/bin/auth /go/bin/auth

EXPOSE 8080

ENTRYPOINT ["/go/bin/auth"]
