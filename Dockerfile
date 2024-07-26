
FROM golang:1.22 AS builder

WORKDIR /workdir

COPY go.mod go.mod

RUN go mod download

COPY cmd cmd

RUN go build -v -o healthcheck ./cmd/healthcheck

# ==============================================================================

FROM busybox:stable AS busybox

# ==============================================================================

FROM gcr.io/distroless/base-debian12

COPY --from=builder /workdir/healthcheck healthcheck
COPY --from=busybox /bin/busybox /busybox/busybox
RUN ["/busybox/busybox", "ln", "-sf", "/busybox/busybox", "/bin/sh"]
