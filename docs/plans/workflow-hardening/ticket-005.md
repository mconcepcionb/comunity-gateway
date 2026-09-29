---
id: ticket-005
title: Stronger quality gates
plan: workflow-hardening
status: done
type: chore
priority: medium
estimate: S
depends_on: [ticket-002]
related_adr: [ADR-0012]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, ci, testing]
---

# ticket-005: Stronger quality gates

## Context

The quality gate is fmt + vet + lint + test. It misses race conditions, known vulnerabilities,
coverage and conditional frontend checks.

## Scope

- Run tests with `-race`.
- Add `govulncheck` and a `coverage` target with a configurable floor.
- Add a path-filtered `console` task (no-op until `apps/console` exists).
- Add `task check:ci` running the full set for CI.

## Out of scope

- Performance benchmarks.

## Definition of Done

- [ ] `task test` runs with `-race`.
- [ ] `task coverage` enforces a floor; `task govulncheck` scans dependencies.
- [ ] `task console` no-ops when the console is absent and lints/typechecks it when present.
- [ ] `task check:ci` composes them; documented in `docs/WORKFLOW.md`.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task check` pass
- [ ] Commit made as `chore(workflow): strengthen quality gates (ticket-005)`
