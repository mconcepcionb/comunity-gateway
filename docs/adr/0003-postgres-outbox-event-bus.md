---
id: ADR-0003
title: PostgreSQL with a transactional outbox and in-process event bus
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [persistence, events, reliability]
---

# ADR-0003: PostgreSQL with a transactional outbox and in-process event bus

## Context

Gamification reacts to domain events, and events must not be lost when a process crashes. The
project also values a small operational footprint. The README asks for a flexible event/hook model.

## Decision

Use PostgreSQL as the single datastore. Domain events are persisted through a transactional outbox
in the same transaction as the state change, then relayed by watermill; consumers run in-process.
Queries use `sqlc`; schema changes use `golang-migrate`.

## Consequences

**Positive**

- Events are never lost: they commit atomically with the state they describe.
- One datastore to operate, back up and reason about.
- Consumers are plain Go code, easy to test.

**Negative / trade-offs**

- At-least-once delivery: consumers must be idempotent.
- In-process consumers do not scale horizontally without partitioning or locking (acceptable here).

## Alternatives considered

- **Kafka / NATS** — operational overkill for a single-node community deployment.
- **Synchronous module calls** — couples modules and loses events on failure.
- **Naive in-memory bus** — events vanish on crash, breaking gamification guarantees.
