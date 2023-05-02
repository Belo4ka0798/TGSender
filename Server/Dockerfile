FROM golang:latest AS builder
WORKDIR /build
COPY ./ ./
RUN CGO_ENABLED=0 go build -o tgsender cmd/app/main.go

FROM ubuntu:latest
WORKDIR /root
COPY --from=builder build/tgsender ./
COPY --from=builder build/config/config.yml ./config/
COPY --from=builder build/internal/static ./static/
CMD ["./tgsender"]
