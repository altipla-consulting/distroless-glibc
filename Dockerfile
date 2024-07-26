
FROM golang:1.22 AS builder

WORKDIR /workdir

COPY go.mod go.mod

RUN go mod download

COPY cmd cmd

RUN go build -v -o healthcheck ./cmd/healthcheck

# ==============================================================================

FROM gcr.io/distroless/base-debian12

COPY --from=builder /workdir/healthcheck /opt/healthcheck
COPY --from=gcr.io/distroless/base-debian12:debug /busybox/sh /bin/sh
COPY --from=gcr.io/distroless/base-debian12:debug /busybox/ls /bin/ls
COPY --from=gcr.io/distroless/base-debian12:debug /busybox/cat /bin/cat

COPY healthcheck /bin/healthcheck
