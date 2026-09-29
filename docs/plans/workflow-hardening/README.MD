---
id: workflow-hardening
title: Workflow Hardening
kind: meta
status: active
owner: marce
created: 2026-09-29
updated: 2026-09-29
related_adrs: [ADR-0016, ADR-0017]
---

# Plan: Workflow Hardening

## Goal

Turn the engineering workflow from prose into something enforced and self-maintaining:
a validator, generated indexes, ticket lifecycle helpers, commit discipline, a stronger quality
gate, and a defined home for meta/tooling work.

## Scope

**In scope**

- A `meta` plan kind so tooling/workflow work has a legitimate home.
- A spike output rule (every spike ends in an ADR or a finding).
- `task verify`: frontmatter, placeholders, dependencies, cycles and index sync.
- Generated plan index and ADR index (`task plan:sync`, `task adr:index`).
- Ticket lifecycle helpers (`task ticket:start`, `task ticket:close`).
- Commit hooks and a changelog (`task hooks:install`, `task changelog`).
- A stronger quality gate (`-race`, `govulncheck`, coverage, path-filtered console).
- Slimmed `AGENTS.md` (single source of truth in `docs/WORKFLOW.md`).

**Out of scope**

- Application features (owned by the `mvp-foundation` plan).

## Definition of Done (plan level)

- [ ] `task verify` passes on the repository.
- [ ] Plan and ADR indexes are generated, never hand-edited.
- [ ] Ticket lifecycle and commit hooks are wired.
- [ ] `docs/WORKFLOW.md` documents all of the above.
- [ ] All tickets `done`.
- [ ] `task check` green on `main`.

## Tickets

| ID | Title | Status | Depends on |
|----|-------|--------|------------|
| ticket-001 | Add meta plan kind and spike output rule | done |  |
| ticket-002 | Workflow validator and index generation | done | ticket-001 |
| ticket-003 | Ticket lifecycle helpers | todo | ticket-002 |
| ticket-004 | Commit hooks and changelog | todo | ticket-002 |
| ticket-005 | Stronger quality gates | todo | ticket-002 |
| ticket-006 | Slim AGENTS.md and adopt ADR index | todo | ticket-002 |

## Risks

- The validator must tolerate real-world frontmatter quirks; test it against the whole repo.
- Nested tooling must run before `go.mod` exists; run it as a standalone program.

## Diary links

- `docs/diary/entry-2026-09-29-workflow-bootstrap.md`
