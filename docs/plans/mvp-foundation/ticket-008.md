---
id: ticket-008
title: "HTTP API: router, OpenAPI contract and auth middleware"
plan: mvp-foundation
status: todo
type: feat
priority: high
estimate: L
depends_on: [ticket-006]
related_adr: [ADR-0004, ADR-0005]
created: 2026-09-29
updated: 2026-09-29
tags: [api, openapi, http]
---

# ticket-008: HTTP API — router, OpenAPI contract and auth middleware

## Context

The gateway exposes a REST API consumed by the console and future clients. The contract is a
hand-written `api/openapi.yaml` with generated typed handlers over the standard library (ADR-0004).

## Scope

- `api/openapi.yaml` with the initial endpoints (auth, health).
- `oapi-codegen` configuration and generation wired into the build.
- `internal/httpapi`: router using `net/http` patterns, auth middleware validating JWTs, error
  mapping, request logging.
- Serve Scalar/Redoc documentation from the spec.
- A generated Go client for the console to consume.

## Out of scope

- Domain endpoints (added by their modules as they land).

## Definition of Done

- [ ] Spec generates typed handlers and a Go client reproducibly.
- [ ] Auth middleware protects routes and maps failures to proper status codes.
- [ ] Docs UI is served; generation is part of `task` tooling.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(httpapi): router, openapi and auth middleware (ticket-008)`
