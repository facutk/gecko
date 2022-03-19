FROM golang:alpine as builder
# RUN apk update && apk add --no-cache git libc6-compat

WORKDIR $GOPATH/src/facutk/gecko/
COPY . .
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
# RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata
# RUN go get -d -v
RUN go build -ldflags "-linkmode external -extldflags -static" -tags osusergo,netgo,sqlite_omit_load_extension -o /go/bin/gecko main.go

FROM scratch

COPY --from=builder /go/bin/gecko /go/bin/gecko

EXPOSE 8080

ENTRYPOINT ["/go/bin/gecko", "-dsn", "/var/lib/litestream/db"]
