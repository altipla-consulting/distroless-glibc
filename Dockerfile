
FROM busybox AS busybox

# ==============================================================================

FROM golang:1.22 AS builder

WORKDIR /workdir

COPY go.mod go.mod

RUN go mod download

COPY cmd cmd

RUN go build -v -o healthcheck ./cmd/healthcheck

# ==============================================================================

FROM gcr.io/distroless/base-debian12:latest

WORKDIR /bin

COPY --from=busybox /bin/busybox /bin/sh
COPY --from=builder /workdir/healthcheck healthcheck
