#!/usr/bin/env bash
set -euo pipefail

root="$(git rev-parse --show-toplevel)"
title="${1:-}"

if [ -z "$title" ]; then
  echo "usage: new-adr.sh \"Title\"" >&2
  exit 1
fi

slug="$(printf '%s' "$title" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//')"
dir="$root/docs/adr"
mkdir -p "$dir"

last="$(find "$dir" -maxdepth 1 -name '[0-9][0-9][0-9][0-9]-*.md' -printf '%f\n' 2>/dev/null | sort | tail -n1 || true)"
n=1
if [ -n "$last" ]; then
  n=$(( 10#${last%%-*} + 1 ))
fi
num="$(printf '%04d' "$n")"
date="$(date +%F)"
file="$dir/$num-$slug.md"

sed -e "s/ADR-0000/ADR-$num/g" \
    -e "s|<short decision title>|$title|g" \
    -e "s/YYYY-MM-DD/$date/g" \
    "$root/docs/templates/adr.md" > "$file"

echo "$file"
