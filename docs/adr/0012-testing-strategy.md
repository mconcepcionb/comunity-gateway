---
id: ADR-0012
title: Testing strategy — unit, testcontainers, protocol fake and a real server
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [testing, ci]
---

# ADR-0012: Testing strategy — unit, testcontainers, protocol fake and a real server

## Context

We need strong confidence without a slow, flaky CI. Game-server integration is the riskiest area,
and fakes can drift from real behavior.

## Decision

Four layers:

1. **Unit tests** per module, with fakes at module boundaries.
2. **Integration tests** against PostgreSQL via testcontainers.
3. A **protocol-level fake** of the game server for fast tests.
4. **One integration test against a real AzerothCore in Docker** for authenticity.

The gamification rules engine is covered by golden tests. CI runs `make check` (fmt, vet, lint,
test).

## Consequences

**Positive**

- Fast inner loop with fakes; an authentic outer check with the real server.
- The fake is exercised by the real-server test, limiting drift.

**Negative / trade-offs**

- The real-server test is heavy and should run selectively or on a schedule.
- Maintaining a protocol fake is ongoing work.

## Alternatives considered

- **Fakes only** — misses real-world protocol quirks.
- **Real server only** — slow and heavy for every change.
- **Mocks everywhere** — brittle and coupled to implementation.
