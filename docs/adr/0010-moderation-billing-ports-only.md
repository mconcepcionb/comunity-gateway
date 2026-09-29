---
id: ADR-0010
title: Moderation and billing as ports only in v1
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [scope, moderation, billing]
---

# ADR-0010: Moderation and billing as ports only in v1

## Context

The README lists five capability pillars, but the MVP is the vertical identity + one game +
gamification. Implementing all five at once would dilute the showcase.

## Decision

Define interfaces (ports) and the data model for moderation and billing in v1, but implement no
features. Real implementations arrive in v2, each with its own ADR.

## Consequences

**Positive**

- The architecture stays honest: every pillar has a boundary.
- The MVP stays focused on one polished vertical.

**Negative / trade-offs**

- The ports may drift from real needs; they must be revisited when implemented.

## Alternatives considered

- **Implement basic moderation/billing now** — dilutes the MVP.
- **Omit them entirely** — weakens the modular, multi-pillar story.
