---
id: ticket-001
title: Add meta plan kind and spike output rule
plan: workflow-hardening
status: done
type: docs
priority: high
estimate: S
depends_on: []
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, docs]
---

# ticket-001: Add meta plan kind and spike output rule

## Context

The workflow says "never implement outside a plan", but tooling and workflow changes have no home.
Spikes also have no required output, so they can end without a decision.

## Scope

- Add a `kind: product | meta` field to plan frontmatter; update the plan template and the existing
  plans (`mvp-foundation` = product, `workflow-hardening` = meta).
- Define the meta-work policy in `docs/WORKFLOW.md`: meta plans cover tooling/workflow changes and
  are exempt from the product roadmap.
- Define the spike output rule: a `spike` ticket that reaches `done` must link a decision (ADR) or
  a finding document.
- Add `docs/templates/finding.md` and a `finding` field to the ticket template.

## Out of scope

- Validator enforcement (ticket-002).

## Definition of Done

- [ ] Plan template and both plans carry a `kind`.
- [ ] `docs/WORKFLOW.md` documents the meta policy and the spike rule.
- [ ] `docs/templates/finding.md` exists; ticket template has a `finding` field.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` passes (after ticket-002; run manually otherwise)
- [ ] Commit made as `docs(workflow): add meta plan kind and spike rule (ticket-001)`
