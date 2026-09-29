---
id: ticket-002
title: "Platform: configuration, logging and database"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: M
depends_on: [ticket-001]
related_adr: [ADR-0003]
created: 2026-09-29
updated: 2026-09-29
tags: [platform, database]
---

# ticket-002: Platform — configuration, logging and database

## Context

Every module needs configuration, structured logging and a database handle. This ticket provides
those foundations plus migrations.

## Scope

- `internal/platform/config`: load configuration from environment variables with validation.
- `internal/platform/logging`: `log/slog` JSON logger.
- `internal/platform/database`: connection pool (`pgx`) and a `golang-migrate` runner.
- `migrations/` directory with the initial (empty or baseline) migration.
- `deploy/docker-compose.yml` with PostgreSQL for local development.
- `/healthz` reports database connectivity.

## Out of scope

- Domain schemas (owned by their modules).

## Definition of Done

- [ ] Configuration loads and validates from env; unit tested.
- [ ] Logger wired through the composition root.
- [ ] Migrations run automatically at startup (or via a task) against compose Postgres.
- [ ] `/healthz` reflects database health.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] No secrets committed (only `.env.example`)
- [ ] Commit made as `feat(platform): config, logging and database bootstrap (ticket-002)`
