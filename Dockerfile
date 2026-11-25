FROM golang:1.24-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -o /out/ploos-core ./cmd/ploos-core

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl && rm -rf /var/lib/apt/lists/*
COPY --from=build /out/ploos-core /usr/local/bin/ploos-core
USER 65532:65532
ENV PLOOS_CORE_ADDR=0.0.0.0:8080
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s CMD curl --fail --silent http://127.0.0.1:8080/healthz || exit 1
ENTRYPOINT ["/usr/local/bin/ploos-core"]
