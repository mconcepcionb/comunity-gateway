---
id: entry-2026-09-29-workflow-hardening
date: 2026-09-29
session: workflow-hardening
plan: workflow-hardening
tickets: [ticket-001, ticket-002, ticket-003, ticket-004, ticket-005, ticket-006]
status: closed
---

# Diary — 2026-09-29 — Workflow hardening

## Intent

Turn the freshly defined workflow from prose into something enforced and self-maintaining, by
implementing the eight improvements selected earlier this session.

## What happened

- Added the `kind: product | meta` plan field and the spike output rule; created the `meta` plan
  `workflow-hardening` with one ticket per improvement (ticket-001).
- Built `tools/workflow` (stdlib-only Go, own module) with `verify`, `plan-sync` and `adr-index`;
  plan ticket tables and `docs/adr/README.md` are now generated (ticket-002).
- Added `task ticket:start` / `task ticket:close` lifecycle helpers (ticket-003).
- Added the `commit-msg` hook, `task hooks:install`, `task verify:commits` and the git-cliff
  changelog config with a committed `CHANGELOG.md` (ticket-004).
- Strengthened the gates: `-race`, `govulncheck`, a coverage floor, and a path-filtered console
  task (ticket-005).
- Slimmed `AGENTS.md` to pointers and adopted the generated ADR index (ticket-006).

## Decisions / ADRs

No new ADRs. This work elaborates ADR-0016/ADR-0017; the meta-plan policy and spike rule are
operational workflow details recorded in `docs/WORKFLOW.md`.

## Blockers

- `git-cliff` is not installed, so `task changelog` cannot run yet; a hand-written `CHANGELOG.md`
  placeholder is committed and will be regenerated once the tool is available.
- `task check` still fails by design: there is no root `go.mod` yet. It lands with `mvp-foundation`
  ticket-001, which also decides the module path.

## Next

- Start `mvp-foundation` ticket-001 (repository scaffold; decide the Go module path).
- Run the `mvp-foundation` ticket-003 spike (AzerothCore whisper capture) before the game adapter.
