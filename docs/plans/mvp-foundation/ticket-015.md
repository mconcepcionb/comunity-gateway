---
id: ticket-015
title: "Console: Next.js scaffold with authentication"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-006, ticket-008]
related_adr: [ADR-0014, ADR-0005]
created: 2026-09-29
updated: 2026-09-29
tags: [console, frontend]
---

# ticket-015: Console — Next.js scaffold with authentication

## Context

The admin console is the first client (ADR-0014). It must authenticate members against the gateway
(ADR-0005).

## Scope

- Scaffold `apps/console` with Next.js + React + TypeScript.
- Consume the generated OpenAPI Go client (or a generated TS client).
- Discord login flow; store and refresh JWTs; protected routing.
- Base layout, navigation and a design-system starting point.
- Lint and typecheck wired into `task check` for the console.

## Out of scope

- Feature screens (ticket-016, ticket-017).

## Definition of Done

- [ ] A member can log in to the console and reach a protected area.
- [ ] Token refresh works transparently.
- [ ] `eslint` and `tsc` run as part of `task check` when the console is touched.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test + console eslint/tsc)
- [ ] Commit made as `feat(console): scaffold next.js app with auth (ticket-015)`
