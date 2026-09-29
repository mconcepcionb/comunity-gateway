---
id: ADR-0015
title: English for code and documentation
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [conventions, docs]
---

# ADR-0015: English for code and documentation

## Context

The initial README was written in Spanish. The project may be shared publicly and the Go ecosystem
conventions are English.

## Decision

All code, identifiers, comments, commit messages and documentation are written in English,
including the README, ADRs, plans, tickets and diary.

## Consequences

**Positive**

- Consistent with Go and open-source conventions; accessible to a global audience.

**Negative / trade-offs**

- Spanish-speaking collaborators lose native-language docs.

## Alternatives considered

- **Spanish docs** — smaller audience, inconsistent with the codebase.
- **Bilingual** — double the maintenance for little gain.
