---
id: ticket-019
title: "Docs: publish OpenAPI and finalize README"
plan: mvp-foundation
status: todo
type: docs
priority: medium
estimate: S
depends_on: [ticket-008, ticket-017]
related_adr: [ADR-0004, ADR-0013]
created: 2026-09-29
updated: 2026-09-29
tags: [docs]
---

# ticket-019: Docs — publish OpenAPI and finalize README

## Context

With the slice complete, the documentation must reflect the final state so the project reads well as
a showcase.

## Scope

- Ensure the OpenAPI document is complete and served.
- Update the README: final architecture diagram, run instructions, configuration, screenshots.
- Add a short "how the vertical slice works" walkthrough.
- Close the plan and record its outcome.

## Out of scope

- New functionality.

## Definition of Done

- [ ] OpenAPI published and accurate for all endpoints.
- [ ] README reflects the implemented system with run instructions and screenshots.
- [ ] Plan marked `done` with the ticket index fully closed.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `docs: publish openapi and finalize readme (ticket-019)`
