---
id: ADR-0006
title: Discord-primary identity with Google and Steam linking
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [identity, oauth]
---

# ADR-0006: Discord-primary identity with Google and Steam linking

## Context

Gaming communities live on Discord, and members often have accounts on multiple platforms. The
README calls for login with Discord, Google, Steam and "other popular platforms", with multiple
accounts linked to one profile.

## Decision

Discord is the primary login. Google and Steam can be linked to the same profile, so one member may
hold many external identities. Steam is integrated via a dedicated **OpenID 2.0** adapter because it
does not offer OAuth 2.0.

## Consequences

**Positive**

- Frictionless, familiar login for the target audience.
- A single profile can aggregate a member's platform accounts.

**Negative / trade-offs**

- The identity model must represent many external accounts to one profile.
- Steam diverges from the OAuth path and needs its own adapter and tests.

## Alternatives considered

- **Email/password primary** — friction and a security burden.
- **Steam primary** — OpenID 2.0 quirks and narrower community reach.
- **No linking** — members would juggle separate accounts.
