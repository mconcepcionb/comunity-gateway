---
id: ADR-0001
title: Modular monolith with compile-time modules and manual DI
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [architecture, modularity]
---

# ADR-0001: Modular monolith with compile-time modules and manual DI

## Context

The product spans identity, game integration, gamification, moderation, monetization and widgets.
A core requirement is that capabilities can be added or removed without touching the core, but the
project is a solo portfolio piece that also needs a simple deployment and fast tests.

## Decision

Build a modular monolith. Each capability is a compile-time Go module (package) that exposes an
explicit interface and is wired in a manual composition root. Modules interact through their public
interfaces and through an in-process event bus; no module reaches into another's internals. There is
no runtime plugin loading.

## Consequences

**Positive**

- One deployable, one process, simple operations.
- Strong typing and compile-time safety across module boundaries.
- Fast, in-memory tests.

**Negative / trade-offs**

- Adding or changing a module requires recompiling the whole binary.
- Boundaries are enforced by convention and review, not by the runtime.

## Alternatives considered

- **Microservices** — operationally heavy and premature for a solo project.
- **Runtime plugins (`hashicorp/go-plugin`)** — fragile ABI, complex packaging, little benefit here.
- **Single unstructured package** — would forfeit the modularity the project is meant to showcase.
