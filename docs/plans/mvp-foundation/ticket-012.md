---
id: ticket-012
title: "Gamification: points, achievements and rules engine"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-004]
related_adr: [ADR-0009]
created: 2026-09-29
updated: 2026-09-29
tags: [gamification, rules]
---

# ticket-012: Gamification — points, achievements and rules engine

## Context

Gamification must be flexible: rules map domain events to awards without core code changes
(ADR-0009).

## Scope

- Domain model: a points ledger, achievements and awards.
- Rule schema (YAML) with a loader and validation; rules overridable in the database.
- A rules engine that evaluates an event against rules and produces awards.
- Persistence for rules, awards and balances; migrations.
- Unit and golden tests for the engine.

## Out of scope

- Subscribing to real events (ticket-013).

## Definition of Done

- [ ] Rules are loaded, validated and evaluated against events.
- [ ] Points ledger is append-only and balances are derived.
- [ ] Golden tests cover representative rule sets.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(gamification): add rules engine and ledger (ticket-012)`
