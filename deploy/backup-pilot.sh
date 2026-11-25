#!/bin/bash
set -euo pipefail
umask 077
dir=/etc/ploos-apps/ploos-core
dest=/srv/ploos-apps/ploos-core/data/backups
install -d -m 0700 "$dest"
compose=(docker compose --project-name ploos-core --env-file "$dir/.env" -f "$dir/compose.yaml")
tmp=$(mktemp "$dest/.backup.XXXXXXXX")
trap 'rm -f "$tmp"' EXIT
"${compose[@]}" exec -T db pg_dump -U postgres -d ploos_core -Fc > "$tmp"
test -s "$tmp"
"${compose[@]}" exec -T db pg_restore --list < "$tmp" >/dev/null
final="$dest/core-$(date -u +%Y%m%dT%H%M%SZ)-${tmp##*.}.dump"
mv "$tmp" "$final"
echo "Backup created: $final"
