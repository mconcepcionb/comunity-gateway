# Engineering Workflow

This document is the single source of truth for how work is done in this repository. The agent
must follow it (see `AGENTS.md`), and humans are expected to as well.

## Principles

- **One decision, one ADR.** Any non-trivial decision is recorded before or alongside its
  implementation.
- **All work is planned.** Nothing is implemented outside a plan; plans decompose into tickets.
- **Tickets are atomic.** One deliverable each, independently verifiable, committed alone.
- **Every ticket ends in a commit.** Never commit red; `task check` must be green first.
- **Every session leaves a diary entry.** Chronological memory of the project.
- **Excellent tooling.** Code stays clean, formatted, linted and tested.

## Artifacts

| Artifact | Location                              | Template                 | Nature                        |
|----------|---------------------------------------|--------------------------|-------------------------------|
| ADR      | `docs/adr/NNNN-<slug>.md`             | `docs/templates/adr.md`  | Immutable once accepted       |
| Plan     | `docs/plans/<plan-name>/README.MD`    | `docs/templates/plan.md` | Living until the plan closes  |
| Ticket   | `docs/plans/<plan-name>/ticket-NNN.md`| `docs/templates/ticket.md`| Atomic, status-tracked       |
| Diary    | `docs/diary/entry-YYYY-MM-DD-<slug>.md`| `docs/templates/diary.md`| Append-only, one per session  |

## ADRs (Architecture Decision Records)

- Every non-trivial decision produces an ADR.
- Format: MADR-lite — **Context**, **Decision**, **Consequences**, **Alternatives considered**.
- Globally numbered, zero-padded to four digits (`0001`).
- **Accepted ADRs are immutable.** To change a decision, write a new ADR and set
  `supersedes` / `superseded_by` on both.
- Status lifecycle: `proposed` → `accepted` → (`superseded` | `deprecated`).
- Create one with `task new-adr title="..."`.

## Plans

- A plan is a bounded body of work with a goal and a plan-level Definition of Done.
- Lives at `docs/plans/<plan-name>/README.MD`; the plan name is kebab-case.
- Contains a **ticket index table** that must stay current as tickets change status.
- A plan is `draft` until reviewed, then `active`, then `done`.
- Create one with `task new-plan name=... title="..."`.

## Tickets

- **Atomic:** a single deliverable that can be reviewed and verified on its own.
- Frontmatter carries all metadata. Required fields:
  `id, title, plan, status, type, priority, estimate, depends_on, related_adr, created, updated, tags`.
- Body sections: **Context**, **Scope**, **Out of scope**, **Definition of Done** (checklist),
  **Quality Gates** (checklist), **Notes**.
- Status lifecycle: `todo` → `in_progress` → `review` → `done` (or `blocked` / `cancelled`).
- A ticket is **not started until its `depends_on` tickets are `done`**.
- Create one with `task new-ticket plan=<plan-name> title="..."`.

## Commits

- **Trunk-based.** Work lands on `main`. Short-lived branches are allowed but never long-lived.
- **Exactly one commit per ticket**, using Conventional Commits and referencing the ticket id:

  ```
  feat(games): add azerothcore adapter (ticket-010)
  ```

- The commit contains: the code, its tests, the ticket file (status set to `done`), the plan index
  update, any ADR the ticket produced, and the diary entry if it is the session's last ticket.
- Commit only when the ticket's Definition of Done and Quality Gates are checked.

## Diary

- One entry per work session: `docs/diary/entry-YYYY-MM-DD-<session-slug>.md`.
- Records **Intent**, **What happened**, **Decisions / ADRs**, **Blockers**, **Next**.
- Append-only: never rewrite a past entry.
- Create one with `task new-diary slug=...`; commit it with the session's final ticket commit.

## Quality gates

`task check` runs `fmt` + `vet` + `lint` + `test`. When the console is touched, frontend gates
(`eslint`, `tsc`) run too. Every ticket's Definition of Done must include: tests added/updated,
docs updated, and `task check` green.

## Tooling

Go 1.27+, `golangci-lint` v2, [Task](https://taskfile.dev) (go-task), Node 26+, bash.

```
task               # list tasks
task check         # fmt + vet + lint + test
task fmt vet lint test
task build run tidy
task new-adr       title="..."
task new-plan      name=... title="..."
task new-ticket    plan=... title="..."
task new-diary     slug=...
```

CI mirrors `task check`.

## Directory conventions

```
docs/
  WORKFLOW.md
  adr/                 NNNN-<slug>.md
  plans/
    <plan-name>/
      README.MD
      ticket-NNN.md
  diary/
    entry-YYYY-MM-DD-<session-slug>.md
  templates/           adr.md plan.md ticket.md diary.md
scripts/               new-adr.sh new-plan.sh new-ticket.sh new-diary.sh
```
