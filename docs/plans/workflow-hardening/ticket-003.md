---
id: ticket-003
title: Ticket lifecycle helpers
plan: workflow-hardening
status: done
type: feat
priority: medium
estimate: M
depends_on: [ticket-002]
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, tooling]
---

# ticket-003: Ticket lifecycle helpers

## Context

Starting and closing a ticket means editing frontmatter, syncing the plan index and crafting a
commit by hand. This is error-prone and easy to skip.

## Scope

- `ticket-start <id> [plan]` — refuse unless `todo`/`blocked` and all `depends_on` are `done`; set
  `status: in_progress` and `updated`.
- `ticket-close <id> [plan]` — set `status: done`, run `plan-sync`, and print the exact commit
  command.
- Task targets `ticket:start` and `ticket:close`.

## Out of scope

- Automatic committing (commits stay explicit).

## Definition of Done

- [ ] Starting a blocked ticket fails with a clear message.
- [ ] Closing a ticket updates frontmatter and the plan index and prints the commit line.
- [ ] Both are covered by the tool's tests.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task check` pass
- [ ] Commit made as `feat(workflow): add ticket lifecycle helpers (ticket-003)`
