# Mimir monitoring

The Forseti pilot listens on 127.0.0.1:18090. Both /healthz and /readyz must return HTTP 200. Readiness additionally checks database connectivity and schema compatibility.

Do not publish Core or PostgreSQL to the Internet for monitoring. First verify whether Forseti has an authenticated private monitoring path to Mimir. If not, establish one before configuring remote probes. A restricted private reverse proxy or local probe agent can expose only the two operational endpoints. Never expose customer API routes through that monitoring path.

Configure separate Blackbox Exporter HTTP probes for health and readiness using the verified private endpoint. Suggested Prometheus scrape interval: 30s, timeout: 10s. Suggested alert rules:

    - alert: PloosCoreProbeFailed
      expr: probe_success{job="ploos-core"} == 0
      for: 2m
      labels:
        severity: warning
      annotations:
        summary: "Ploos Core probe failed: {{ $labels.instance }}"
    - alert: PloosCoreMonitoringMissing
      expr: absent(probe_success{job="ploos-core",check="health"}) or absent(probe_success{job="ploos-core",check="readiness"})
      for: 5m
      labels:
        severity: warning
      annotations:
        summary: "Ploos Core monitoring series missing"

Set check labels to health and readiness respectively. Validate rules with promtool and test alert delivery before relying on them. Monitor host disk and memory using the existing node exporter architecture. Backup freshness requires a separate metric and alert; HTTP readiness does not establish backup health.

No private address, firewall rule, proxy or live Mimir configuration is assumed or changed by this document.
