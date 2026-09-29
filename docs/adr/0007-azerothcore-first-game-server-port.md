---
id: ADR-0007
title: AzerothCore as the first game integration behind a GameServer port
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [games, integration, azerothcore]
---

# ADR-0007: AzerothCore as the first game integration behind a GameServer port

## Context

Four games are targeted: World of Warcraft (AzerothCore), Project Zomboid, Conan Exiles and Rust.
Each exposes different protocols (SOAP/REST/DB for AzerothCore, RCON for the others). The gateway
must integrate games one at a time without leaking game specifics into the core.

## Decision

Implement World of Warcraft / AzerothCore first, behind a `GameServer` port that abstracts server
commands and character reads. AzerothCore specifics (SOAP/REST calls, DB reads) live only in its
adapter. Other games implement the same port later.

## Consequences

**Positive**

- Proves the abstraction against a rich, complex server.
- AzerothCore offers the most integration surface (REST/SOAP/DB) for a convincing showcase.

**Negative / trade-offs**

- The port risks being shaped by AzerothCore's quirks; game-specific logic must stay in the adapter.
- Other games remain unproven until implemented.

## Alternatives considered

- **Zomboid / Rust / Conan first** — simpler RCON, but far less integration surface.
- **All games at once** — breadth over depth, against the showcase goal.
