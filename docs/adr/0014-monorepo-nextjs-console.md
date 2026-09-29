---
id: ADR-0014
title: Monorepo with backend and apps; Next.js admin console first
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [repo, frontend, console]
---

# ADR-0014: Monorepo with backend and apps; Next.js admin console first

## Context

The gateway ships a Go backend plus reference clients (admin console, Discord bot, web). The admin
console is the first client to build (it drives the MVP workflows).

## Decision

Use a single monorepo: the Go backend at the root (`cmd/`, `internal/`), clients under `apps/`, with
the admin console built in Next.js + React + TypeScript. Cross-cutting changes land atomically in
one commit.

## Consequences

**Positive**

- Atomic changes across backend and clients; one CI pipeline.
- The console can consume the generated OpenAPI client directly.

**Negative / trade-offs**

- Multiple toolchains coexist in one repo.
- Frontend and backend release together (acceptable here).

## Alternatives considered

- **Polyrepo** — synchronization overhead for a solo project.
- **Backend-only monorepo** — splits tooling and history.
- **SvelteKit / Vue** — viable, but React/Next has the larger ecosystem.
