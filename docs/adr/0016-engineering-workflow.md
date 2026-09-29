---
id: ADR-0016
title: Repository engineering workflow
status: superseded
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by: ADR-0017
tags: [process, workflow]
---

# ADR-0016: Repository engineering workflow

## Context

The project needs a repeatable process that records decisions, structures work, guarantees quality
and preserves a chronological memory across sessions.

## Decision

Adopt the following workflow, specified in `docs/WORKFLOW.md` and enforced for agents through
`AGENTS.md`:

- **ADRs** for every non-trivial decision; globally numbered, MADR-lite, immutable once accepted.
- **Plans** under `docs/plans/<plan-name>/README.MD` with a live ticket index.
- **Atomic tickets** (`ticket-NNN.md`) with frontmatter metadata, a Definition of Done and Quality
  Gates.
- **Trunk-based, one commit per ticket**, Conventional Commits referencing the ticket id.
- **`make check`** (fmt + vet + lint + test) as the quality gate before every commit.
- A **per-session diary** under `docs/diary/`.
- Scaffolding through Makefile targets backed by `scripts/`.

## Consequences

**Positive**

- Traceability from decision to code; consistent quality; durable memory across sessions.

**Negative / trade-offs**

- Some ceremony per change; mitigated by templates and Make targets.

## Alternatives considered

- **Ad-hoc development** — no traceability and no durable memory.
- **Heavyweight process (e.g. full RFC + release trains)** — overkill for a solo project.
