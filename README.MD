# Community Gateway

A modular backend that helps online gaming communities manage their members end to end:
identity, game-server integration, gamification, moderation and monetization — plus reference
clients (admin console, Discord bot, web) and shareable widgets.

This project is a **portfolio showcase**. The goal is architectural quality and **one polished
vertical slice**, not feature breadth. Correct boundaries, clean code, tests and docs matter more
than shipping every capability at once.

---

## Goals and non-goals

**Goals**

- Demonstrate a clean **modular monolith in Go** with explicit module boundaries and a manual DI
  composition root.
- Ship one **end-to-end vertical**: Discord identity → link a WoW character via in-game
  verification → earn points and achievements from real server activity.
- Keep extension cheap: modules communicate through **domain events** and gamification is defined
  by **declarative rules**, so new behavior rarely touches core code.
- Be observable, tested and documented.

**Non-goals (v1)**

- Moderation and monetization **features** — only ports/interfaces and data model are defined.
- Multi-tenancy — one deployment serves **one community**.
- Games beyond AzerothCore — the integration port is prepared, only WoW is implemented.
- Runtime plugins — modules are compile-time.

---

## Architecture

A single deployable **modular monolith**: compile-time Go modules wired by hand at the
composition root, over one Postgres database with a **transactional outbox** and an in-process
**event bus**.

Within each module we use a layered layout (domain → application → adapters) so the domain never
depends on transport or storage. Inter-module calls go through the owning module's public
interface; cross-cutting reactions go through the event bus.

```
        HTTP (stdlib)                 Discord bot (v2)         Web client (v2)
              |                              |                        |
              +------------------------------+------------------------+
                              httpapi module
                                    |
   +-----------+-----------+---------------+-----------+-----------+
   |           |           |               |           |           |
identity     games     gamification    moderation    billing     widgets
   |           |           |               |           |           |
   +-----------+-----------+---- events (bus + outbox) -----------+
                                    |
                                PostgreSQL
```

### Modules

| Module         | Responsibility                                                                 | v1 status        |
|----------------|--------------------------------------------------------------------------------|------------------|
| `events`       | In-process event bus + Postgres transactional outbox (watermill).              | implemented      |
| `identity`     | Accounts, OAuth providers, account linking, sessions/JWT issuance.             | implemented      |
| `games`        | Game-server integrations, characters, account verification. `GameServer` port. | implemented (WoW)|
| `gamification` | Points, achievements, rewards; declarative rules over events.                  | implemented      |
| `moderation`   | Warnings, bans, appeals.                                                       | port only        |
| `billing`      | Subscriptions, donations, entitlements.                                        | port only        |
| `widgets`      | Activity widgets, stat images, "wrapped" (HTML → PNG).                         | implemented      |
| `httpapi`      | HTTP layer, auth middleware, OpenAPI.                                          | implemented      |

### Decisions behind the architecture

- **Modular monolith, compile-time modules.** Each module is a Go package exposing an interface
  and registering itself in a manual DI container. No runtime plugin loading.
- **Postgres + transactional outbox.** Domain events are written in the same transaction as the
  state change, then relayed by watermill, so events are never lost.
- **No web framework.** Routing uses `net/http` (Go 1.22+ method/pattern matching). The API
  contract lives in a hand-written `api/openapi.yaml`; `oapi-codegen` generates typed handlers
  over the standard library, and we serve Scalar/Redoc from the same spec.
- **Data access via `sqlc`**, schema via `golang-migrate`, logging via `log/slog`.

---

## Identity and account linking

- **Discord is the primary identity.** Users log in with Discord via OAuth.
- **Google and Steam can be linked** to the same profile, so one member can hold many external
  accounts.
- The gateway is its **own identity provider**: it finishes the upstream OAuth dance and issues
  its own **JWT access + refresh** tokens. The admin console and public API use those JWTs; bots
  use service credentials.
- Note: Steam uses **OpenID 2.0**, not OAuth 2.0 — it needs a dedicated adapter.

## Game integration and the trust model

The hard problem is proving a member really owns the in-game character they claim. The flow for
AzerothCore is:

1. The user generates a **short-lived verification code** in the web console; the gateway stores
   it against the profile.
2. In-game, the user **whispers the code to a designated bot account**.
3. A small **server-side capture** persists the whisper; the gateway reads it from the game DB.
4. The gateway matches the code, pairs the character with the profile, and emits
   `character.verified`.

`games` hides all of this behind a `GameServer` port (SOAP/REST/DB read for AzerothCore). Adding a
new game means implementing the port, not touching other modules.

> **Open risk — whisper capture.** AzerothCore does not persist whispers by default, so step 3
> requires *something* on the server side (a minimal C++ module or an Eluna/Lua script that writes
> the bot account's whispers to a dedicated table), or tailing the worldserver chat log. This is
> the only part that touches the game server's code and must be validated with an early spike.

## Gamification

Modules publish **domain events** (e.g. `character.verified`, `player.session.completed`,
`member.joined`). **Declarative rules** (YAML shipped with the app, overridable in DB) map events
to awards — points, achievements, rewards — and are evaluated by a small rules engine. Adding a
rule requires no core code change, satisfying the "flexible hooks/events" requirement.

## Widgets

Activity feeds, participation stats and a shareable "wrapped" are rendered by **HTML/CSS templates
converted to PNG with headless Chrome** (chromedp), so the visuals can reuse the console's design
system. Served through dedicated cached endpoints.

---

## MVP (v1) vertical slice

1. Log in with Discord; link Google/Steam.
2. Connect an AzerothCore server; verify a character with the in-game code flow.
3. Ingest activity events and award points/achievements through declarative rules.
4. **Admin console (Next.js):** manage members, links, game servers, view points, edit rules.
5. Public read API with generated OpenAPI docs.

## Roadmap

- **v1** — vertical slice above.
- **v2** — Discord bot, moderation features, billing features, additional games, public web
  profile.

---

## Tech stack

| Concern            | Choice                                                        |
|--------------------|---------------------------------------------------------------|
| Language           | Go (backend), TypeScript (console)                            |
| HTTP               | `net/http` stdlib + `oapi-codegen` handlers (contract-first)  |
| API contract       | Hand-written `api/openapi.yaml`, served as Scalar/Redoc       |
| Database           | PostgreSQL; `sqlc` for queries, `golang-migrate` for schema   |
| Events             | In-process bus + transactional outbox via watermill           |
| Auth               | OAuth upstream (Discord/Google/Steam) → own JWT access+refresh|
| Logging / DI       | `log/slog`; manual composition root                           |
| Widgets            | HTML/CSS → PNG via chromedp                                   |
| Console            | Next.js + React + TypeScript                                  |
| Tests / CI         | `testing`, testcontainers (Postgres), fake game server, GH Actions |

## Repository layout

```
comunity-gateway/
  cmd/gateway/            # composition root / main
  internal/
    platform/             # config, logging, db, di, http helpers
    events/               # bus + transactional outbox
    identity/             # accounts, providers, linking, tokens
    games/                # GameServer port + azerothcore adapter
    gamification/         # points, achievements, rules engine
    moderation/           # ports/interfaces only (v1)
    billing/              # ports/interfaces only (v1)
    widgets/              # HTML→PNG rendering
    httpapi/              # routing, middleware, generated handlers
  api/openapi.yaml        # API contract (source of truth)
  migrations/             # golang-migrate
  apps/console/           # Next.js admin console
  game-mods/azerothcore/  # whisper-capture module / script
  deploy/                 # docker-compose (gateway, postgres, azerothcore)
```

## Testing

- Unit tests per module, with fakes at module boundaries.
- Integration tests against Postgres via testcontainers.
- A protocol-level fake of the game server for fast tests, plus **one real AzerothCore in
  Docker** for an authenticity integration test.
- Golden tests for the gamification rules engine.
- CI on GitHub Actions: lint, vet, unit, integration.

## Development workflow

This repository follows a documented engineering workflow: every non-trivial decision is an
**ADR** (`docs/adr/`), all work lives in a **plan** (`docs/plans/<plan-name>/`) split into atomic
**tickets**, each ticket ends in exactly one commit, `task check` is the quality gate, and every
session leaves a **diary** entry (`docs/diary/`). See `docs/WORKFLOW.md`.

The current plan is **MVP Foundation** — `docs/plans/mvp-foundation/README.MD`.

```
task             # list tasks        task check
task new-adr       title="..."
task new-plan      name=... title="..."
task new-ticket    plan=... title="..."
task new-diary     slug=...
```

## Open questions

- Whisper capture on AzerothCore: C++ module vs Eluna/Lua script vs chat-log tailer (spike
  ticket-003).
- The Go module path (decided in ticket-001).
- Outbox implementation: watermill's SQL backend vs a small custom relay.
- Whether widgets run inside the main API process or a separate worker.
- Steam OpenID 2.0 adapter details.
