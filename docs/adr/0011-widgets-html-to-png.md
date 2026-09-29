---
id: ADR-0011
title: Widgets rendered as HTML to PNG with headless Chrome
status: accepted
date: 2026-09-29
deciders: [marce]
supersedes:
superseded_by:
tags: [widgets, rendering]
---

# ADR-0011: Widgets rendered as HTML to PNG with headless Chrome

## Context

The README promises shareable widgets: a recent-activity widget, a participation-stats image and a
year-in-review "wrapped" for social media. These are visual and design-heavy.

## Decision

Render HTML/CSS templates to PNG with headless Chrome via `chromedp`, reusing the admin console's
design system. Widgets are exposed through dedicated, cached endpoints behind a renderer interface.

## Consequences

**Positive**

- Rich, on-brand visuals with the same CSS skills as the console.
- The renderer interface keeps the dependency isolated and mockable.

**Negative / trade-offs**

- A headless browser is a heavy dependency to ship and run.
- Rendering is slower than native drawing, so caching is required.

## Alternatives considered

- **Native Go drawing (`fogleman/gg`)** — lightweight but limited typography and layout.
- **SVG + resvg** — lighter, but less flexible than CSS.
- **Defer widgets** — weakens the showcase the README advertises.
