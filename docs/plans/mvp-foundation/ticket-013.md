---
id: ticket-013
title: "Gamification: event wiring and golden tests"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: M
depends_on: [ticket-011, ticket-012]
related_adr: [ADR-0009]
created: 2026-09-29
updated: 2026-09-29
tags: [gamification, events]
---

# ticket-013: Gamification — event wiring and golden tests

## Context

With the rules engine (ticket-012) and real domain events (ticket-011) in place, gamification must
consume events and apply awards.

## Scope

- Subscribe to domain events (e.g. `character.verified`, `member.joined`) on the event bus.
- Evaluate the applicable rules and apply awards idempotently.
- Emit `gamification.awarded` for downstream consumers (widgets, notifications).
- Golden tests covering event → award end to end.

## Out of scope

- Points display and rule editing UI (ticket-017).

## Definition of Done

- [ ] Events trigger awards idempotently under at-least-once delivery.
- [ ] `gamification.awarded` is published.
- [ ] Golden tests demonstrate the full event → award path.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(gamification): wire events to rules (ticket-013)`
