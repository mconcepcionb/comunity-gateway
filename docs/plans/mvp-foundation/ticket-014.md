---
id: ticket-014
title: "Widgets: HTML-to-PNG activity widget"
plan: mvp-foundation
status: todo
type: feat
priority: medium
estimate: M
depends_on: [ticket-004]
related_adr: [ADR-0011]
created: 2026-09-29
updated: 2026-09-29
tags: [widgets, rendering]
---

# ticket-014: Widgets — HTML-to-PNG activity widget

## Context

The README promises shareable widgets. The first is a recent-activity widget rendered as an image
(ADR-0011).

## Scope

- A renderer interface with a headless-Chrome (`chromedp`) implementation.
- An HTML/CSS template for the recent-activity widget.
- The data query that feeds it, driven by domain events.
- A cached endpoint returning a PNG.
- Unit tests with a fake renderer plus one real render test.

## Out of scope

- Stats images and the "wrapped" (v2).
- Design-system reuse from the console (can come later).

## Definition of Done

- [ ] The activity widget renders to PNG and is served from a cached endpoint.
- [ ] The renderer is isolated behind an interface and faked in fast tests.
- [ ] One real headless-Chrome render test passes.
- [ ] Docs updated, ticket `done`, plan index updated.

## Quality Gates

- [ ] `task check` passes (fmt + vet + lint + test)
- [ ] Commit made as `feat(widgets): add activity widget via chromedp (ticket-014)`
