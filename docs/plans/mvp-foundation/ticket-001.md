---
id: ticket-001
title: Repository scaffold and quality tooling
plan: mvp-foundation
status: todo
type: chore
priority: high
estimate: M
depends_on: []
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [tooling, ci]
---

# ticket-001: Repository scaffold and quality tooling

## Context

The repository currently holds only the workflow, docs and scripts. A Go module, the package layout
and a working `task check` in CI are needed before any feature work.

## Scope

- Initialise the Go module (decide and document the module path).
- Create the package skeleton: `cmd/gateway`, `internal/platform`, `internal/events`,
  `internal/identity`, `internal/games`, `internal/gamification`, `internal/moderation`,
  `internal/billing`, `internal/widgets`, `internal/httpapi`.
- A runnable `cmd/gateway` that starts an HTTP server with a `/healthz` endpoint.
- A GitHub Actions workflow running `task check`.
- An `.editorconfig` consistent with Go and TypeScript.

## Out of scope

- Any real domain functionality (covered by later tickets).

## Definition of Done

- [ ] `go.mod` exists with the chosen module path documented in the README or an ADR.
- [ ] `task build`, `task run` and `task check` succeed.
- [ ] CI workflow runs `task check` on push and pull request.
- [ ] Package skeleton committed with doc comments on each package.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] No secrets committed
- [ ] Commit made as `chore(repo): scaffold module and tooling (ticket-001)`

## Notes

The module path must be decided here. Default candidate: `github.com/marcelinocb/comunity-gateway`.
