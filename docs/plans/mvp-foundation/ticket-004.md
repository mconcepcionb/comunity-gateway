---
id: ticket-004
title: "Events: in-process bus and transactional outbox"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-002]
related_adr: [ADR-0003, ADR-0009]
created: 2026-09-29
updated: 2026-09-29
tags: [events, reliability]
---

# ticket-004: Events — in-process bus and transactional outbox

## Context

Gamification reacts to domain events. Events must commit atomically with the state that produces
them and be delivered reliably (ADR-0003).

## Scope

- `internal/events`: `DomainEvent` type, `Publisher` and `Subscriber` interfaces.
- Transactional outbox table and repository (`sqlc`).
- Relay built on watermill that polls the outbox and dispatches to in-process subscribers.
- A small typed helper so modules publish events inside a database transaction.
- Migrations and tests, including idempotency and restart safety.

## Out of scope

- Concrete event types (defined by the modules that emit them).

## Definition of Done

- [ ] Events can be published within a transaction and are delivered at least once.
- [ ] Subscribers can be registered per event type.
- [ ] Integration test covers publish → relay → deliver and a simulated crash/restart.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(events): add bus and transactional outbox (ticket-004)`
