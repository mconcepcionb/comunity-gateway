---
id: ticket-017
title: "Console: points and gamification rules"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: M
depends_on: [ticket-015, ticket-013]
related_adr: [ADR-0009, ADR-0014]
created: 2026-09-29
updated: 2026-09-29
tags: [console, frontend, gamification]
---

# ticket-017: Console — points and gamification rules

## Context

The rule engine is only useful if administrators can inspect and edit rules and see awards.

## Scope

- Points and achievements view per member and across the community.
- Rules list with enable/disable, edit (YAML or structured form) and validation feedback.
- A preview of which rules match a given event.
- Audit of applied awards.

## Out of scope

- Billing or entitlements (v2).

## Definition of Done

- [ ] Admins can view, edit and validate rules; invalid rules are rejected with clear errors.
- [ ] Awards and balances are visible per member.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test + console eslint/tsc)
- [ ] Commit made as `feat(console): add points and rules management (ticket-017)`
