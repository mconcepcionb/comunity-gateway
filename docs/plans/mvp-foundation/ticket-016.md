---
id: ticket-016
title: "Console: members, links and game servers"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-015, ticket-011]
related_adr: [ADR-0014]
created: 2026-09-29
updated: 2026-09-29
tags: [console, frontend, games]
---

# ticket-016: Console — members, links and game servers

## Context

Administrators manage members and game servers, and members verify their characters — all from the
console.

## Scope

- Members list and detail: linked identities, characters, points.
- Character verification UI driven by the ticket-011 flow (generate code, show status).
- Game server CRUD (register an AzerothCore server and its connection settings).
- Optimistic/cached data fetching and error states.

## Out of scope

- Rule editing (ticket-017).

## Definition of Done

- [ ] Admins can create, view and remove game servers.
- [ ] Members can start and complete a character verification from the console.
- [ ] Members and their links are listed and searchable.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test + console eslint/tsc)
- [ ] Commit made as `feat(console): add members, links and servers (ticket-016)`
