FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux
RUN apk add upx

WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -a --ldflags "-w -s" -gcflags=all="-l -B" -o hero .
RUN upx --best --lzma hero

#FROM alpine
#FROM gcr.io/distroless/static-debian12
FROM scratch
WORKDIR /zero
COPY --from=builder /build/hero /zero/hero
COPY --from=builder /build/database/ /zero/database/
COPY --from=builder /build/public/ /zero/public/
COPY --from=builder /build/storage/ /zero/storage/
COPY --from=builder /build/.env /zero/.env

ENTRYPOINT ["/zero/hero"]