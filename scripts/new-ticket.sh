#!/usr/bin/env bash
set -euo pipefail

root="$(git rev-parse --show-toplevel)"
plan="${1:-}"
title="${2:-}"

if [ -z "$plan" ] || [ -z "$title" ]; then
  echo "usage: new-ticket.sh <plan-name> \"Title\"" >&2
  exit 1
fi

dir="$root/docs/plans/$plan"
if [ ! -d "$dir" ]; then
  echo "plan not found: $plan" >&2
  exit 1
fi

last="$(find "$dir" -maxdepth 1 -name 'ticket-[0-9][0-9][0-9].md' -printf '%f\n' 2>/dev/null | sort | tail -n1 || true)"
n=1
if [ -n "$last" ]; then
  base="${last#ticket-}"
  base="${base%.md}"
  n=$(( 10#$base + 1 ))
fi
num="$(printf '%03d' "$n")"
date="$(date +%F)"
file="$dir/ticket-$num.md"

sed -e "s/ticket-000/ticket-$num/g" \
    -e "s|<Ticket title>|$title|g" \
    -e "s|<plan-name>|$plan|g" \
    -e "s/YYYY-MM-DD/$date/g" \
    "$root/docs/templates/ticket.md" > "$file"

echo "$file"
