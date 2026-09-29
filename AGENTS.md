# AGENTS.md

Instructions for AI agents working in this repository. Follow these rules on every task.

## Read first

1. `docs/WORKFLOW.md` — the full engineering workflow. Obey it.
2. The relevant ADRs in `docs/adr/`.
3. The active plan in `docs/plans/<plan-name>/README.MD` and its tickets.

## Non-negotiable rules

- **Every non-trivial decision produces an ADR** (`docs/adr/NNNN-<slug>.md`, MADR-lite). Generate
  it with `task new-adr title="..."`. Accepted ADRs are immutable; supersede, never edit.
- **All work happens inside a plan.** Never implement outside a plan. If none exists, create one
  with `task new-plan name=... title="..."`.
- **Plans split into atomic tickets** (`ticket-NNN.md`) with frontmatter metadata, an explicit
  Definition of Done, and Quality Gates. Use `task new-ticket plan=... title="..."`.
- **Do not start a ticket until its `depends_on` tickets are `done`.**
- **One commit per ticket**, trunk-based, Conventional Commits referencing the ticket id, e.g.
  `feat(games): add azerothcore adapter (ticket-010)`.
- **Never commit red.** Run `task check` (fmt + vet + lint + test) and confirm it passes before
  committing. Frontend gates (`eslint`, `tsc`) when the console is touched.
- **Set the ticket's status to `done` and update the plan's ticket index in the same commit.**
- **Write a diary entry** per session (`docs/diary/entry-YYYY-MM-DD-<slug>.md`) and commit it with
  the session's final ticket.
- **Never commit secrets.** Keep credentials out of the repo.
- Commits are made only as part of completing a ticket, or when the user explicitly asks.

## Quality bar

- Clean, idiomatic, formatted Go; no dead code; no TODOs without a ticket.
- Tests accompany behavior changes. Prefer table-driven tests and fakes at module boundaries.
- Update docs affected by the change (README, ADRs, plan).

## Tooling

`task check`, `task fmt`, `task vet`, `task lint`, `task test`, `task build`, `task run`,
`task new-adr`, `task new-plan`, `task new-ticket`, `task new-diary`.
