---
id: entry-2026-09-29-workflow-bootstrap
date: 2026-09-29
session: workflow-bootstrap
plan: mvp-foundation
tickets: []
status: closed
---

# Diary — 2026-09-29 — Project definition and workflow bootstrap

## Intent

Review the README, run a grilling session to pin down how the project should be implemented, and
formalize the engineering workflow for the repository.

## What happened

- Rewrote `README.MD` in English as a working spec, capturing the decisions below.
- Ran four rounds of questions covering tenancy, MVP scope, modularity, persistence, identity,
  game integration, the trust model, gamification, clients, API style, widgets and testing.
- Defined the workflow: ADRs, plans, atomic tickets, one commit per ticket, quality gates and a
  per-session diary.
- Created `docs/WORKFLOW.md`, `AGENTS.md`, templates (`docs/templates/`), the `Taskfile` and the
  scaffolding scripts in `scripts/`.
- Bootstrapped 16 ADRs (ADR-0001..0016) from the session's decisions.
- Created the plan `docs/plans/mvp-foundation/` with 19 tickets covering the v1 vertical slice.

## Decisions / ADRs

ADR-0001 modular monolith · ADR-0002 single-community self-hosted · ADR-0003 Postgres + outbox ·
ADR-0004 stdlib HTTP + contract-first OpenAPI · ADR-0005 gateway-issued JWT over upstream OAuth ·
ADR-0006 Discord-primary identity · ADR-0007 AzerothCore first behind a port · ADR-0008 in-game
verification code · ADR-0009 declarative gamification rules · ADR-0010 moderation/billing ports
only · ADR-0011 widgets via HTML-to-PNG · ADR-0012 testing strategy · ADR-0013 showcase scope ·
ADR-0014 monorepo + Next.js console · ADR-0015 English everywhere · ADR-0016 engineering workflow.

## Blockers

- None for the definition work. The MVP has a known risk: capturing whispers on AzerothCore
  (ticket-003 spike), tracked in ADR-0008 and the plan's risks.

## Next

- Decide the Go module path (ticket-001).
- Start ticket-001 (repository scaffold and quality tooling).
- Run the ticket-003 spike early, before investing in the AzerothCore adapter, because it gates
  ticket-010 and ticket-011.

## Update — task runner

Switched the task runner from GNU Make to Task (`Taskfile.yml`, `task check`, `task new-*`),
removing the `Makefile`. Recorded as ADR-0017, which supersedes the task-runner portion of
ADR-0016; all other workflow decisions are unchanged. Living docs, templates, tickets and the plan
were updated to `task`.
