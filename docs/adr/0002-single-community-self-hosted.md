---
id: ADR-0002
title: Single-community, self-hosted deployment
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [architecture, tenancy]
---

# ADR-0002: Single-community, self-hosted deployment

## Context

Tenancy is a foundational fork: it determines the database schema, the authentication model, the
billing model and the operational story. The project targets communities that run their own game
servers.

## Decision

One deployment serves exactly one community. There is no `tenant_id`; community-wide settings live
in configuration. Multi-tenancy is explicitly deferred.

## Consequences

**Positive**

- Simple schema, simple auth, simple billing.
- Fits the self-hosted, privacy-friendly nature of game communities.
- Keeps the showcase focused.

**Negative / trade-offs**

- Cannot host many communities on a single instance.
- Introducing multi-tenancy later would be a significant migration requiring a new ADR.

## Alternatives considered

- **Multi-tenant SaaS** — requires row-level security, subdomains and per-tenant billing; out of scope.
- **Tenant-aware core from day one** — speculative complexity with no current consumer.
