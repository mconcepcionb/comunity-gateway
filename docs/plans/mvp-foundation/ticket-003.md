---
id: ticket-003
title: "Spike: AzerothCore whisper capture"
plan: mvp-foundation
status: todo
type: spike
priority: critical
estimate: M
depends_on: [ticket-001]
related_adr: [ADR-0008]
created: 2026-09-29
updated: 2026-09-29
tags: [games, azerothcore, spike, risk]
---

# ticket-003: Spike — AzerothCore whisper capture

## Context

ADR-0008 depends on capturing a whisper sent to a bot account; AzerothCore does not persist whispers
by default. This is the biggest unknown in the plan and must be resolved before investing in the
game adapter (ticket-010) and verification flow (ticket-011).

## Scope

- Stand up a local AzerothCore in Docker via `deploy/docker-compose.yml`.
- Evaluate the capture options: a minimal C++ module, an Eluna/Lua script, or tailing the worldserver
  chat log.
- Produce a working proof for the most viable option (a whisper to the bot account becomes readable
  by an external process).
- Write the finding and the recommended approach; create a follow-up ADR if the choice is
  non-trivial.

## Out of scope

- Integrating the capture with the gateway (ticket-010/011).

## Definition of Done

- [ ] AzerothCore runs locally in Docker and a character can log in.
- [ ] A whisper to the bot account is captured by the chosen mechanism.
- [ ] A short written finding documents the options, trade-offs and recommendation.
- [ ] Follow-up ADR created if needed.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Docker assets committed under `deploy/` and `game-mods/azerothcore/`
- [ ] Commit made as `spike(games): validate azerothcore whisper capture (ticket-003)`

## Notes

This spike gates the riskiest part of the MVP. Do not skip it.
