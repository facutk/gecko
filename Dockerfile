FROM golang:alpine as builder
# RUN apk update && apk add --no-cache git libc6-compat

WORKDIR $GOPATH/src/facutk/gecko/
COPY . .
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
# RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata
# RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o /go/bin/gecko gecko.go

FROM scratch
COPY --from=builder /go/bin/gecko /go/bin/gecko
ENTRYPOINT ["/go/bin/gecko"]
