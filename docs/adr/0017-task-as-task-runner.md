---
id: ADR-0017
title: Task (go-task) as the repository task runner
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes: ADR-0016
superseded_by:
tags: [process, workflow, tooling]
---

# ADR-0017: Task (go-task) as the repository task runner

## Context

ADR-0016 defined the engineering workflow and named a GNU Makefile as the scaffolding and quality
gate runner. Make is awkward on Windows, requires a POSIX shell and tab-sensitive syntax, and does
not model task dependencies or required variables cleanly. This supersedes **only the task-runner
choice**; every other decision in ADR-0016 (ADRs, plans, atomic tickets, one commit per ticket,
`check` quality gate, per-session diary) remains in force.

## Decision

Use [Task](https://taskfile.dev) (`go-task`) as the task runner, defined in `Taskfile.yml`. The
quality gate is `task check` (fmt + vet + lint + test). Scaffolding is provided by
`task new-adr`, `task new-plan`, `task new-ticket` and `task new-diary`, backed by the same
`scripts/` shell scripts. The Makefile is removed. Any `make` reference in earlier ADRs denotes the
current `task` equivalent.

## Consequences

**Positive**

- Cross-platform, readable YAML with no whitespace pitfalls.
- Structured dependencies and `requires.vars` give clearer errors and less duplication.
- The same runner works on Windows, Linux and in CI.

**Negative / trade-offs**

- Introduces a tool that must be installed (`go install
  github.com/go-task/task/v3/cmd/task@latest`, or scoop/brew/apt).
- Contributions must assume `task` rather than `make`.

## Alternatives considered

- **Keep GNU Make** (status quo of ADR-0016) — poor Windows ergonomics and tab-sensitivity.
- **Plain shell scripts** — no dependency graph, discovery or per-task descriptions.
- **just** — capable, but less ergonomic for required variables and cross-platform installs.
