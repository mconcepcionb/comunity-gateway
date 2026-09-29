---
id: ticket-002
title: Workflow validator and index generation
plan: workflow-hardening
status: done
type: chore
priority: high
estimate: L
depends_on: [ticket-001]
related_adr: [ADR-0016, ADR-0017]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, tooling]
---

# ticket-002: Workflow validator and index generation

## Context

The workflow rules are prose and prone to drift. This ticket makes them machine-checked and
generates the indexes so status lives in exactly one place (ticket frontmatter).

## Scope

- A standalone Go program at `tools/workflow` (stdlib only) with subcommands:
  - `verify` — frontmatter completeness and enums, placeholder leftovers, dependency existence and
    cycles, spike findings, and plan/ADR index sync.
  - `plan-sync` — regenerate each plan's ticket table from ticket frontmatter.
  - `adr-index` — regenerate `docs/adr/README.md`.
- Task targets `verify`, `plan:sync`, `adr:index`.
- Normalise the existing plans so `verify` passes.

## Out of scope

- Ticket lifecycle helpers (ticket-003).

## Definition of Done

- [ ] `task verify` reports all violations and exits non-zero on failure.
- [ ] `task plan:sync` and `task adr:index` regenerate their indexes deterministically.
- [ ] The whole repository passes `task verify`.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task check` pass
- [ ] Commit made as `chore(workflow): add validator and index generation (ticket-002)`
