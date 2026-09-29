#!/usr/bin/env bash
set -euo pipefail

root="$(git rev-parse --show-toplevel)"
slug="${1:-session}"

slug="$(printf '%s' "$slug" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//')"
date="$(date +%F)"
file="$root/docs/diary/entry-$date-$slug.md"

if [ -e "$file" ]; then
  echo "diary entry already exists: $file" >&2
  exit 1
fi

mkdir -p "$root/docs/diary"
sed -e "s/entry-YYYY-MM-DD-slug/entry-$date-$slug/g" \
    -e "s/YYYY-MM-DD/$date/g" \
    -e "s|<session-slug>|$slug|g" \
    "$root/docs/templates/diary.md" > "$file"

echo "$file"
