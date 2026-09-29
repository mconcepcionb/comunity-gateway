---
id: ADR-0004
title: Standard-library HTTP with a contract-first OpenAPI document
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [api, http, openapi]
---

# ADR-0004: Standard-library HTTP with a contract-first OpenAPI document

## Context

The gateway exposes a REST API consumed by the admin console, a future Discord bot and third-party
integrations. We want typed handlers, accurate documentation and minimal dependencies.

## Decision

Use `net/http` (Go 1.22+ method and pattern matching) as the router. The hand-written
`api/openapi.yaml` is the source of truth; `oapi-codegen` generates typed handlers over the standard
library. The same spec is served as Scalar/Redoc documentation.

## Consequences

**Positive**

- No web-framework lock-in; the API contract is authoritative and versioned in the repo.
- Generated handlers keep transport code typed and repetitive boilerplate low.

**Negative / trade-offs**

- A code generation step is required in the build.
- Slightly more ceremony per endpoint than annotation-based frameworks.

## Alternatives considered

- **chi + huma** — convenient but adds dependencies and hides the routing model.
- **gin/echo** — heavier than needed; custom context types.
- **`swag` annotations** — documentation drifts from actual behavior.
