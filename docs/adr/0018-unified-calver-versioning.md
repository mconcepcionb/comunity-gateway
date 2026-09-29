---
id: ADR-0018
title: Unified CalVer versioning
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [workflow, release, versioning]
---

# ADR-0018: Unified CalVer versioning

## Context

The repository had no versioning strategy. Community Gateway is a self-hosted server, and the
monorepo ships backend and clients together (ADR-0014). Its two audiences are operators who deploy
it and HTTP API clients (the console, a future bot, third-party integrations).

SemVer is designed for libraries: for a server, "breaking" is fuzzy (irreversible data migrations,
config-key renames, dropped endpoints), and the project is pre-stable, so major/minor debates would
add cost without signalling much. At the same time, the HTTP API is a contract that needs an
explicit compatibility boundary.

## Decision

- Use a **single unified version for the whole monorepo**: CalVer `YYYY.MM.PATCH` with a `v` prefix
  (e.g. `v2026.09.0`). `PATCH` increments per release within a month and resets to `0` on a new
  month.
- Releases are **ad-hoc and do not map to plans**; `meta` plans never produce a release.
- **Git tags are the source of truth.** The version is embedded into the binary at build time
  (ldflags) and exposed through an endpoint; there is no `VERSION` file to keep in sync.
- The **HTTP API is versioned independently in the path** (`/api/v1`). A breaking API change creates
  a new path version.
- **No product-level compatibility promise** for now: the API path version is the compatibility
  boundary. Breaking changes are called out in the changelog but do not alter the numbering scheme.
- `task release` computes the next CalVer from the latest tag, requires a clean tree and green
  gates, creates an annotated (signed when a key is configured) tag, regenerates the changelog, and
  pushes the tag when a remote exists.

## Consequences

**Positive**

- Version ordering is trivial and monotonic; no major/minor debates.
- One version matches the joint release model of the monorepo.
- The API compatibility boundary is explicit and decoupled from release cadence.

**Negative / trade-offs**

- CalVer carries no compatibility signal; clients must rely on the API version.
- Releases within the same month differ only by a counter, which is less expressive than SemVer.
- Until an API deprecation policy exists, external integrations have no stability promise.

## Alternatives considered

- **SemVer from `0.x`** — rejected: server breakage is fuzzy and the numbering invites debate.
- **SemVer from `1.0.0`** — rejected: an early compatibility promise the project cannot keep.
- **Independent per-component versions** — rejected: ADR-0014 ships components together.
- **Coupling the API version to the product version** — rejected: different lifecycles.
