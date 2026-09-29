---
id: ticket-011
title: "Games: character verification flow"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-004, ticket-006, ticket-010]
related_adr: [ADR-0008]
created: 2026-09-29
updated: 2026-09-29
tags: [games, verification, trust]
---

# ticket-011: Games — character verification flow

## Context

This is the trust core of the MVP: a member proves ownership of a character by whispering a code to
a bot account, and the gateway pairs the character with the profile (ADR-0008).

## Scope

- Generate a short-lived verification code and persist the challenge.
- Read captured whispers from the AzerothCore adapter and match the pending code.
- Create the character link and emit `character.verified` via the event bus.
- Handle expiry, duplicate attempts and mismatches.
- Endpoint(s) to start a verification and to query its status.

## Out of scope

- Console UI for the flow (ticket-016).

## Definition of Done

- [ ] A code whispered in-game links the character and emits `character.verified`.
- [ ] Expired and invalid codes are handled and reported.
- [ ] Integration test with the fake game server covers the full flow.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(games): add character verification flow (ticket-011)`
