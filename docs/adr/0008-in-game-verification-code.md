---
id: ADR-0008
title: In-game verification code whispered to a bot account
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [games, trust, identity]
---

# ADR-0008: In-game verification code whispered to a bot account

## Context

Linking a game character to a gateway profile requires proving that the member actually owns the
character. Reading the game database directly proves existence, not ownership. This is the hardest
trust problem in the README and it was not addressed there.

## Decision

The member generates a short-lived verification code in the web console. In-game, they whisper the
code to a designated bot account. A minimal server-side capture persists the whisper and the gateway
reads it, matches the code, pairs the character with the profile and emits `character.verified`.

## Consequences

**Positive**

- Proves ownership without sharing credentials.
- The pattern is game-agnostic and demonstrable.

**Negative / trade-offs**

- AzerothCore does not persist whispers by default, so a server-side capture (C++ module, Eluna/Lua
  script or chat-log tailer) is required. This is **the open risk of the MVP** and must be spiked
  early.
- The bot account must be online in-game.

## Alternatives considered

- **Direct game-DB reads** — no proof of ownership.
- **OAuth per game** — only available for platform accounts, not in-game characters.
- **Manual admin approval** — does not scale.
