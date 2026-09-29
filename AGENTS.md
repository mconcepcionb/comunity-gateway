# AGENTS.md

Instructions for AI agents working in this repository. This file is a pointer, not a copy: the
rules live in `docs/WORKFLOW.md`, which is the single source of truth.

## Start here

1. `docs/WORKFLOW.md` — how work is done here. Obey it.
2. `docs/adr/README.md` — decision index; read the ADRs relevant to your task.
3. The active plan at `docs/plans/<plan-name>/README.MD` and its tickets.

## Invariants

- **Ask before deciding.** For important decisions about architecture or tooling, present options and
  wait for a decision instead of choosing unilaterally.
- **Follow convention, but don't guess in uncharted territory.** Apply existing conventions, ADRs
  and surrounding patterns; when nothing covers the case, stop and ask rather than invent.
- All work lives in a plan; tooling/workflow changes use a `kind: meta` plan.
- A non-trivial decision produces an ADR. Accepted ADRs are immutable; supersede, never edit.
- Tickets are atomic, with an explicit Definition of Done and Quality Gates. Do not start a ticket
  until its `depends_on` tickets are `done`.
- One commit per ticket, Conventional Commits referencing the ticket id.
- Never commit red or secrets: run `task check` and `task verify` first.
- Close tickets with `task ticket:close` and keep a per-session diary entry.

## Commands

`task` lists every task. Core: `task check`, `task verify`, `task plan:sync`, `task ticket:start`,
`task ticket:close`, `task new-adr`, `task new-plan`, `task new-ticket`, `task new-diary`.
