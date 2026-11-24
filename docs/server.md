# Running Ploos Core

Run from the repository root:

    go run ./cmd/ploos-core

Default listener: 127.0.0.1:8080. Override with PLOOS_CORE_ADDR. For a future container deployment, explicitly bind to 0.0.0.0:8080 and place behind a trusted TLS reverse proxy.

    curl -i http://127.0.0.1:8080/healthz
    curl -i http://127.0.0.1:8080/readyz

Liveness returns 200; readiness deliberately returns 503 until persistent storage and required dependencies are wired. No customer API is exposed. The server has bounded HTTP timeouts and graceful SIGINT/SIGTERM shutdown.

Database configuration, migrations, authentication, Docker packaging and deployment remain follow-up work.
