#!/usr/bin/env bash
set -euo pipefail

root="$(git rev-parse --show-toplevel)"
name="${1:-}"
title="${2:-$name}"

if [ -z "$name" ]; then
  echo "usage: new-plan.sh <name> [title]" >&2
  exit 1
fi

slug="$(printf '%s' "$name" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//')"
dir="$root/docs/plans/$slug"
mkdir -p "$dir"
file="$dir/README.MD"
date="$(date +%F)"

if [ -e "$file" ]; then
  echo "plan already exists: $file" >&2
  exit 1
fi

sed -e "s|<plan-name>|$slug|g" \
    -e "s|<Plan title>|$title|g" \
    -e "s/YYYY-MM-DD/$date/g" \
    "$root/docs/templates/plan.md" > "$file"

echo "$file"
