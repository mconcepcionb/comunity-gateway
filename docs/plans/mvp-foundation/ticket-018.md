---
id: ticket-018
title: "Integration test: real AzerothCore end to end"
plan: mvp-foundation
status: todo
type: test
priority: high
estimate: M
depends_on: [ticket-010, ticket-011, ticket-013]
related_adr: [ADR-0012]
created: 2026-09-29
updated: 2026-09-29
tags: [testing, e2e, azerothcore]
---

# ticket-018: Integration test — real AzerothCore end to end

## Context

The vertical slice must be proven against a real AzerothCore, not only fakes (ADR-0012). This is the
authenticity check for the whole MVP.

## Scope

- A `docker-compose` environment with gateway, PostgreSQL and AzerothCore.
- An automated scenario: authenticate a member, generate a verification code, whisper it in-game,
  verify the character, and assert points/achievements are awarded.
- Wire the test to run in CI on a schedule or in a dedicated job.

## Out of scope

- Console interactions (the API is exercised directly).

## Definition of Done

- [ ] The end-to-end scenario passes against a real AzerothCore.
- [ ] The test is isolated and repeatable, with cleanup.
- [ ] Documented how to run it locally and in CI.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] E2E job passes in CI (scheduled if too heavy for every push)
- [ ] Commit made as `test(e2e): verify full slice against real azerothcore (ticket-018)`
