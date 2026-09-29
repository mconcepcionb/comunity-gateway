---
id: ticket-010
title: "Games: AzerothCore adapter"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-003, ticket-009]
related_adr: [ADR-0007, ADR-0008]
created: 2026-09-29
updated: 2026-09-29
tags: [games, azerothcore]
---

# ticket-010: Games — AzerothCore adapter

## Context

AzerothCore is the first game integration behind the `GameServer` port (ADR-0007). The whisper
capture mechanism is validated by the spike (ticket-003); this ticket implements the adapter.

## Scope

- Implement the `GameServer` port for AzerothCore using its REST/SOAP surface and read-only DB
  queries as needed.
- Implement the whisper-capture reader chosen in ticket-003 (module output or dedicated table).
- Configuration for connecting to an AzerothCore instance.
- Tests against the protocol fake and one test against a real AzerothCore in Docker.

## Out of scope

- The end-to-end verification flow (ticket-011).

## Definition of Done

- [ ] The adapter reads characters and captures whispers from a running AzerothCore.
- [ ] A fake backs fast unit tests; a real-server test validates authenticity.
- [ ] AzerothCore-specific logic is confined to the adapter.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Real-server integration test passes locally
- [ ] Commit made as `feat(games): add azerothcore adapter (ticket-010)`
