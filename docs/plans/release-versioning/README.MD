---
id: release-versioning
title: Release and Versioning
kind: meta
status: done
owner: marce
created: 2026-09-29
updated: 2026-09-29
related_adrs: [ADR-0018]
---

# Plan: Release and Versioning

## Goal

Adopt a unified CalVer versioning scheme for the monorepo and provide tooling to cut releases,
embed the version in the binary and document the strategy.

## Scope

**In scope**

- ADR-0018 recording the versioning strategy.
- `workflow version` and `workflow release` subcommands (CalVer computation, gates, tag, changelog,
  push).
- Task targets: `task version`, `task version:next`, `task release`, `task release:dryrun`.
- Version injection into the build via ldflags; exposure through an endpoint (wired in the app).
- Documentation in `docs/WORKFLOW.md`.

**Out of scope**

- Cutting an actual release (no application exists yet).
- API version negotiation beyond the `/api/v1` path prefix.

## Definition of Done (plan level)

- [ ] ADR-0018 accepted.
- [ ] `task version` and `task release:dryrun` work; `task release` is gated.
- [ ] `docs/WORKFLOW.md` documents the strategy.
- [ ] All tickets `done`.
- [ ] `task verify` green on `main`.

## Tickets

| ID | Title | Status | Depends on |
|----|-------|--------|------------|
| ticket-001 | Adopt CalVer and add release tooling | done |  |

## Risks

- CalVer gives no compatibility signal; mitigated by the independent `/api/v1` version.

## Diary links

- `docs/diary/entry-2026-09-29-release-versioning.md`
