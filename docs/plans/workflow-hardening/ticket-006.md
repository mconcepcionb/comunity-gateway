---
id: ticket-006
title: Slim AGENTS.md and adopt ADR index
plan: workflow-hardening
status: done
type: docs
priority: low
estimate: S
depends_on: [ticket-002]
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, docs]
---

# ticket-006: Slim AGENTS.md and adopt ADR index

## Context

`AGENTS.md` and `docs/WORKFLOW.md` restate the same rules and will drift. There is also no ADR
index for navigation.

## Scope

- Reduce `AGENTS.md` to pointers and the minimum invariants; `docs/WORKFLOW.md` is the single
  source of truth.
- Adopt the generated `docs/adr/README.md` index (produced by ticket-002).
- Add guidance to avoid volatile tool names inside ADR bodies so they do not go stale.

## Out of scope

- Rewriting existing accepted ADRs.

## Definition of Done

- [ ] `AGENTS.md` no longer duplicates the workflow rules.
- [ ] `docs/adr/README.md` is committed and referenced from the docs.
- [ ] `docs/WORKFLOW.md` warns against volatile tool names in ADRs.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task check` pass
- [ ] Commit made as `docs(workflow): slim AGENTS.md and add ADR index (ticket-006)`
