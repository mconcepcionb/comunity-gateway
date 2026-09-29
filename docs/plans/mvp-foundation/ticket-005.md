---
id: ticket-005
title: "Identity: account and external-identity model"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: M
depends_on: [ticket-002]
related_adr: [ADR-0006]
created: 2026-09-29
updated: 2026-09-29
tags: [identity, database]
---

# ticket-005: Identity — account and external-identity model

## Context

The gateway aggregates a member's Discord, Google and Steam accounts into one profile (ADR-0006).
This ticket defines that model and its persistence.

## Scope

- Tables: `profiles`, `external_identities` (provider, external id, profile), and
  `refresh_tokens`.
- Domain types and a repository interface for the identity module.
- `sqlc` queries and migrations.
- Unit tests at the repository boundary.

## Out of scope

- OAuth flows and token issuance (ticket-006/007).
- Linking UX (ticket-007, console).

## Definition of Done

- [ ] Schema migrated and queries generated.
- [ ] Repository interface with a Postgres implementation.
- [ ] Integration test with testcontainers covering create profile, link identity, uniqueness.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(identity): add account and external-identity model (ticket-005)`
