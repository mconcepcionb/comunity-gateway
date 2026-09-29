---
id: ticket-000
title: <Ticket title>
plan: <plan-name>
status: todo            # todo | in_progress | review | blocked | done | cancelled
type: feat              # feat | fix | refactor | docs | chore | test | spike
priority: medium        # low | medium | high | critical
estimate: S             # XS | S | M | L | XL
depends_on: []
related_adr: []
finding:                # spike only: path to docs/plans/<plan>/findings/<id>.md
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
---

# ticket-000: <Ticket title>

## Context

Why this ticket exists and what it unblocks.

## Scope

Exactly what is delivered. Keep it atomic.

## Out of scope

What is deliberately not done here (and which ticket covers it, if any).

## Definition of Done

- [ ] Implementation complete
- [ ] Tests added/updated
- [ ] Docs updated (README / ADR / plan)
- [ ] Ticket status set to `done` and plan index updated

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Frontend gates pass if the console was touched (`eslint`, `tsc`)
- [ ] No secrets committed
- [ ] Commit made as `type(scope): summary (ticket-NNN)`

## Notes

Implementation notes, discoveries, follow-ups.
