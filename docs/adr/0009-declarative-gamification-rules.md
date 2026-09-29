---
id: ADR-0009
title: Declarative gamification rules over domain events
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [gamification, events, extensibility]
---

# ADR-0009: Declarative gamification rules over domain events

## Context

The README requires a flexible system where modules expose hooks or events that grant points,
achievements and rewards dynamically, ideally without changing the core for each new rule.

## Decision

Modules emit domain events. Declarative rules (YAML shipped with the application, overridable in the
database) map events to awards — points, achievements and rewards — and are evaluated by a small
rules engine. A module may register a Go handler as an escape hatch for logic too complex to be
declarative.

## Consequences

**Positive**

- New rules require no core code change and, for DB-overridden rules, no redeploy.
- The event bus from ADR-0003 feeds gamification naturally.

**Negative / trade-offs**

- The rule schema must be versioned and validated.
- Overly complex conditions do not fit a declarative model and fall back to Go handlers.

## Alternatives considered

- **Hardcoded handlers per module** — inflexible, contradicts the README.
- **CEL expressions** — more powerful but a heavier concept to design and document.
- **External rules engine** — an extra dependency and a separate mental model.
