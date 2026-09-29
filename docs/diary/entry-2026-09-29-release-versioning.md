---
id: entry-2026-09-29-release-versioning
date: 2026-09-29
session: release-versioning
plan: release-versioning
tickets: [ticket-001]
status: closed
---

# Diary — 2026-09-29 — Release and versioning

## Intent

Define how the project is versioned and released, and wire tooling for it.

## What happened

- Ran a grilling session and chose a unified CalVer scheme over SemVer (ADR-0018).
- Added `workflow version [--next]` and `workflow release [--dry-run]`: CalVer computation from git
  tags, clean-tree and gate checks, annotated (or signed) tag, changelog regeneration and tag push.
- Added Task targets `version`, `version:next`, `release`, `release:dryrun`; `task build` now injects
  the version via ldflags.
- Documented the strategy in `docs/WORKFLOW.md`.

## Decisions / ADRs

- **ADR-0018: Unified CalVer versioning** — one `vYYYY.MM.PATCH` version for the monorepo, releases
  ad-hoc and not mapped to plans, git tags as the source of truth, the API versioned independently
  at `/api/v1`, and no product-level compatibility promise yet.

## Blockers

- No release can be cut until the application builds (`task release` requires green gates), i.e.
  until `mvp-foundation` ticket-001 lands.
- `git-cliff` is still not installed; `task release` skips changelog regeneration until it is.

## Next

- Start `mvp-foundation` ticket-001 (repository scaffold; decide the Go module path; expose the
  version endpoint).
