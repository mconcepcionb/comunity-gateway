---
id: mvp-foundation
title: MVP Foundation
kind: product
status: active
owner: marce
created: 2026-09-29
updated: 2026-09-29
related_adrs:
  [ADR-0001, ADR-0002, ADR-0003, ADR-0004, ADR-0005, ADR-0006, ADR-0007,
   ADR-0008, ADR-0009, ADR-0010, ADR-0011, ADR-0012, ADR-0013, ADR-0014,
   ADR-0015, ADR-0016, ADR-0017]
---

# Plan: MVP Foundation

## Goal

Deliver the v1 vertical slice end to end: a member logs in with Discord, verifies a World of
Warcraft character in-game, earns points and achievements through declarative rules, and an
administrator manages members, links, game servers and rules from the admin console.

## Scope

**In scope**

- Repository scaffold, tooling and quality gates.
- `platform` (config, logging, database, migrations).
- `events` (in-process bus + transactional outbox).
- `identity` (accounts, Discord/Google/Steam, JWT access + refresh).
- `games` (`GameServer` port, AzerothCore adapter, in-game verification flow).
- `gamification` (points, achievements, declarative rules engine).
- `widgets` (one HTML-to-PNG activity widget).
- `httpapi` (stdlib router, OpenAPI contract, auth middleware).
- Admin console (Next.js) covering members, links, servers, points and rules.
- Integration test against a real AzerothCore in Docker.

**Out of scope**

- Moderation and billing features (ports only, see ADR-0010).
- Games other than AzerothCore (see ADR-0007).
- Discord bot and public web client.
- Multi-tenancy (see ADR-0002).

## Definition of Done (plan level)

- [ ] The full vertical slice works end to end against a real AzerothCore in Docker.
- [ ] OpenAPI document published and served.
- [ ] Admin console covers members, links, game servers, points and rules.
- [ ] All tickets `done`.
- [ ] `task check` green on `main`.

## Tickets

| ID | Title | Status | Depends on |
|----|-------|--------|------------|
| ticket-001 | Repository scaffold and quality tooling | done |  |
| ticket-002 | Platform: configuration, logging and database | todo | ticket-001 |
| ticket-003 | Spike: AzerothCore whisper capture | todo | ticket-001 |
| ticket-004 | Events: in-process bus and transactional outbox | todo | ticket-002 |
| ticket-005 | Identity: account and external-identity model | todo | ticket-002 |
| ticket-006 | Identity: Discord login and JWT issuance | todo | ticket-005 |
| ticket-007 | Identity: link Google and Steam identities | todo | ticket-006 |
| ticket-008 | HTTP API: router, OpenAPI contract and auth middleware | todo | ticket-006 |
| ticket-009 | Games: GameServer port and domain model | todo | ticket-002 |
| ticket-010 | Games: AzerothCore adapter | todo | ticket-003, ticket-009 |
| ticket-011 | Games: character verification flow | todo | ticket-004, ticket-006, ticket-010 |
| ticket-012 | Gamification: points, achievements and rules engine | todo | ticket-004 |
| ticket-013 | Gamification: event wiring and golden tests | todo | ticket-011, ticket-012 |
| ticket-014 | Widgets: HTML-to-PNG activity widget | todo | ticket-004 |
| ticket-015 | Console: Next.js scaffold with authentication | todo | ticket-006, ticket-008 |
| ticket-016 | Console: members, links and game servers | todo | ticket-015, ticket-011 |
| ticket-017 | Console: points and gamification rules | todo | ticket-015, ticket-013 |
| ticket-018 | Integration test: real AzerothCore end to end | todo | ticket-010, ticket-011, ticket-013 |
| ticket-019 | Docs: publish OpenAPI and finalize README | todo | ticket-008, ticket-017 |

## Risks

- **Whisper capture on AzerothCore** (ticket-003) — the single biggest unknown; it shapes
  ticket-010/011. Spike before investing in the adapter.
- **Headless Chrome for widgets** — heavy dependency; isolate behind the renderer interface.
- **Steam OpenID 2.0** — a non-standard flow; needs its own adapter and tests.

## Diary links

- `docs/diary/entry-2026-09-29-workflow-bootstrap.md`
