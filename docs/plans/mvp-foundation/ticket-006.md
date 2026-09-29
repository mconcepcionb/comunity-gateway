---
id: ticket-006
title: "Identity: Discord login and JWT issuance"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-005]
related_adr: [ADR-0005, ADR-0006]
created: 2026-09-29
updated: 2026-09-29
tags: [identity, oauth, auth]
---

# ticket-006: Identity — Discord login and JWT issuance

## Context

Discord is the primary identity (ADR-0006). The gateway completes the upstream OAuth flow and issues
its own JWT access and refresh tokens (ADR-0005).

## Scope

- Discord OAuth 2.0 authorization-code flow with state and PKCE.
- Profile resolution: link to an existing profile or create a new one.
- Token service: JWT access tokens, rotating refresh tokens persisted and revocable.
- Refresh and logout endpoints; revocation of refresh tokens.
- Unit and integration tests, including rotation and replay rejection.

## Out of scope

- Google and Steam (ticket-007).
- HTTP routing and middleware (ticket-008 wires these handlers).

## Definition of Done

- [ ] An OAuth login with Discord yields a profile and a valid token pair.
- [ ] Refresh rotates and revokes correctly; replay is rejected.
- [ ] Tests cover the happy path and failure modes (invalid state, expired refresh).
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] No secrets committed; credentials read from config
- [ ] Commit made as `feat(identity): add discord login and jwt issuance (ticket-006)`
