
FROM golang:1.22 AS builder

WORKDIR /workdir

COPY go.mod go.mod

RUN go mod download

COPY cmd cmd

RUN go build -v -o healthcheck ./cmd/healthcheck

# ==============================================================================

FROM gcr.io/distroless/base-debian12:debug

COPY --from=builder /workdir/healthcheck healthcheck
