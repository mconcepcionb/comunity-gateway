---
id: ADR-0005
title: Gateway-issued JWT access and refresh tokens over upstream OAuth
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [auth, identity, security]
---

# ADR-0005: Gateway-issued JWT access and refresh tokens over upstream OAuth

## Context

Several clients (admin console, future Discord bot, API consumers) must authenticate against the
gateway. Upstream providers (Discord, Google, Steam) authenticate users but do not provide a
uniform, gateway-scoped session.

## Decision

The gateway acts as its own identity provider. It completes the upstream OAuth flow, then issues its
own short-lived JWT access tokens and rotating refresh tokens. The console and public API use these
tokens; bots use service credentials. Refresh tokens are stored and revocable server-side.

## Consequences

**Positive**

- One authentication model for every client.
- Stateless access tokens, revocable refresh tokens.
- Upstream providers can be added without changing clients.

**Negative / trade-offs**

- Refresh-token rotation and revocation must be implemented correctly.
- Token lifecycle (expiry, replay) needs care and tests.

## Alternatives considered

- **Cookie sessions** — awkward for bots and third-party API consumers.
- **Raw upstream tokens** — no unified identity, inconsistent expiry and scopes.
