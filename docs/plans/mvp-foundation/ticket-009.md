---
id: ticket-009
title: "Games: GameServer port and domain model"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: M
depends_on: [ticket-002]
related_adr: [ADR-0007]
created: 2026-09-29
updated: 2026-09-29
tags: [games, architecture]
---

# ticket-009: Games — GameServer port and domain model

## Context

Game integrations must be pluggable behind a `GameServer` port so new games do not touch the core
(ADR-0007). This ticket defines the port and the domain model.

## Scope

- Domain types: `GameServer` (configuration), `Character`, verification challenge.
- `GameServer` port interface: read a character, and any command execution the adapter exposes.
- `games` module service for registering servers and characters.
- Persistence for game servers and characters; migrations; repository interface.
- Unit tests with a fake `GameServer`.

## Out of scope

- The AzerothCore implementation (ticket-010).
- The verification flow (ticket-011).

## Definition of Done

- [ ] The port is defined and a fake exercises the service.
- [ ] Game servers and characters persist; uniqueness enforced.
- [ ] Unit tests cover the service with the fake.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(games): add game server port and domain model (ticket-009)`
