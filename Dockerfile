
FROM ubuntu:23.04

RUN apt update && \
    apt install -y ca-certificates tzdata curl
