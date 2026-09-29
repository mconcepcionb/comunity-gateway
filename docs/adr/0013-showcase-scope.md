---
id: ADR-0013
title: Portfolio showcase scope — one polished vertical slice
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [scope, goals]
---

# ADR-0013: Portfolio showcase scope — one polished vertical slice

## Context

Community Gateway is intended as a portfolio piece demonstrating modular and scalable systems
design, not as an all-features product.

## Decision

Optimize for architectural quality and a single end-to-end vertical: Discord login → verify a WoW
character in-game → earn points/achievements via declarative rules → manage it from the admin
console. Breadth (more games, more clients, moderation, billing) is deferred and tracked in the
roadmap.

## Consequences

**Positive**

- A clear, demonstrable result that exercises the whole architecture.
- Quality is spent where it is visible.

**Negative / trade-offs**

- The delivered capability set is deliberately incomplete.

## Alternatives considered

- **Breadth-first** — many shallow features, weak demonstration.
- **Product-first** — different priorities and quality bar.
