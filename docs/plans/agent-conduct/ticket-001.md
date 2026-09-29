---
id: ticket-001
title: Add ask-before-deciding conduct rules
plan: agent-conduct
status: done
type: docs
priority: high
estimate: XS
depends_on: []
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, agents]
---

# ticket-001: Add ask-before-deciding conduct rules

## Context

Agents have been making architecture and tooling choices autonomously. The user wants agents to ask
before important decisions, to follow existing convention by default, and to ask when in uncharted
territory rather than guessing.

## Scope

- Add the conduct rules to the `## Invariants` section of `AGENTS.md`:
  - Ask before important architecture or tooling decisions.
  - Follow existing convention; when no convention or ADR covers the case, stop and ask.

## Out of scope

- Changing workflow mechanics.

## Definition of Done

- [ ] `AGENTS.md` states both rules.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` passes
- [ ] Commit made as `docs(agents): add ask-before-deciding conduct rules (ticket-001)`
