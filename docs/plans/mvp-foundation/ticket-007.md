---
id: ticket-007
title: "Identity: link Google and Steam identities"
plan: mvp-foundation
status: todo
type: feat
priority: medium
estimate: M
depends_on: [ticket-006]
related_adr: [ADR-0006]
created: 2026-09-29
updated: 2026-09-29
tags: [identity, oauth, steam]
---

# ticket-007: Identity — link Google and Steam identities

## Context

A profile can aggregate Google and Steam accounts alongside Discord (ADR-0006). Steam uses OpenID
2.0, not OAuth 2.0, so it needs a dedicated adapter.

## Scope

- Google OAuth 2.0 link flow (link to the authenticated profile, not a new login).
- Steam OpenID 2.0 adapter with its distinct protocol handling.
- Conflict handling: an external identity already linked to another profile.
- Unit tests for both adapters using fakes.

## Out of scope

- Primary login changes; Discord remains primary.

## Definition of Done

- [ ] A logged-in member can link Google and Steam to their profile.
- [ ] Conflicts are detected and reported clearly.
- [ ] Adapters are tested with fakes; Steam protocol quirks documented.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(identity): link google and steam identities (ticket-007)`
