---
id: ticket-001
title: Adopt CalVer and add release tooling
plan: release-versioning
status: done
type: chore
priority: high
estimate: M
depends_on: []
related_adr: [ADR-0018]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, tooling, release]
---

# ticket-001: Adopt CalVer and add release tooling

## Context

The repository has no versioning strategy. The decision (see ADR-0018) is a unified CalVer
(`vYYYY.MM.PATCH`) for the monorepo, an independently versioned HTTP API (`/api/v1`), git tags as
the source of truth, and tooling to cut releases.

## Scope

- Write ADR-0018 with the strategy.
- Add `workflow version [--next]` and `workflow release [--dry-run]` (CalVer computation, clean-tree
  and gate checks, annotated/signed tag, changelog regeneration, tag push when a remote exists).
- Add Task targets `version`, `version:next`, `release`, `release:dryrun`; inject the version into
  `task build` via ldflags.
- Tests for the CalVer computation.
- Document the strategy in `docs/WORKFLOW.md`.

## Out of scope

- Cutting a real release (blocked until the application builds).
- Runtime exposure endpoint (owned by `mvp-foundation` ticket-001).

## Definition of Done

- [ ] ADR-0018 accepted and indexed.
- [ ] `task version` and `task version:next` print the expected versions.
- [ ] `task release:dryrun` previews; `task release` enforces a clean tree and green gates.
- [ ] CalVer computation is unit tested.
- [ ] `docs/WORKFLOW.md` documents the strategy.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task test:workflow` pass
- [ ] Commit made as `chore(workflow): adopt CalVer and release tooling (ticket-001)`
