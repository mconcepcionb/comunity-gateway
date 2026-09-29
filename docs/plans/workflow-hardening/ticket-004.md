---
id: ticket-004
title: Commit hooks and changelog
plan: workflow-hardening
status: done
type: chore
priority: medium
estimate: M
depends_on: [ticket-002]
related_adr: [ADR-0016]
created: 2026-09-29
updated: 2026-09-29
tags: [workflow, git]
---

# ticket-004: Commit hooks and changelog

## Context

The commit convention (Conventional Commits + ticket id) is not enforced, and there is no release
history.

## Scope

- A `commit-msg` hook validating Conventional Commits and a ticket reference, with an explicit
  `[no-ticket]` escape hatch.
- `task hooks:install` setting `core.hooksPath` to the repo's hooks.
- `cliff.toml` and `task changelog` generating `CHANGELOG.md` with `git-cliff`.
- A CI step validating the pushed commit messages.

## Out of scope

- Releases and version tags.

## Definition of Done

- [ ] The hook rejects malformed messages and accepts valid ones.
- [ ] `task hooks:install` wires the hook.
- [ ] `task changelog` regenerates `CHANGELOG.md` (documented dependency on git-cliff).
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task verify` and `task check` pass
- [ ] Commit made as `chore(workflow): add commit hooks and changelog (ticket-004)`
