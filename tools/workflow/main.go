package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	reTicketFile = regexp.MustCompile(`^ticket-(\d{3})\.md$`)
	reADRFile    = regexp.MustCompile(`^(\d{4})-[a-z0-9-]+\.md$`)
	reDiaryFile  = regexp.MustCompile(`^entry-(\d{4}-\d{2}-\d{2})-[a-z0-9-]+\.md$`)
	reDate       = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	reADRID      = regexp.MustCompile(`^ADR-\d{4}$`)
	reTicketID   = regexp.MustCompile(`^ticket-\d{3}$`)
	reKey        = regexp.MustCompile(`^([A-Za-z0-9_]+):(.*)$`)
	reRegular    = regexp.MustCompile(`^(feat|fix|refactor|docs|chore|test|spike|perf|build|ci|style)(\([^)]+\))?!?: .+`)
	reTicketRef  = regexp.MustCompile(`\(ticket-\d{3}\)`)

	ticketStatuses = set("todo", "in_progress", "review", "blocked", "done", "cancelled")
	ticketTypes    = set("feat", "fix", "refactor", "docs", "chore", "test", "spike")
	priorities     = set("low", "medium", "high", "critical")
	estimates      = set("XS", "S", "M", "L", "XL")
	planKinds      = set("product", "meta")
	planStatuses   = set("draft", "active", "done")
	adrStatuses    = set("proposed", "accepted", "superseded", "deprecated")
	diaryStatuses  = set("open", "closed")

	placeholderTokens = []string{
		"<Ticket title>", "<Plan title>", "<short decision title>", "<session title>",
		"<session-slug>", "<plan-name>", "<findings-id>", "<alternative>",
		"ticket-000", "ADR-0000", "entry-YYYY-MM-DD-slug",
	}

	ticketSections = []string{"## Context", "## Scope", "## Out of scope", "## Definition of Done", "## Quality Gates"}
)

func set(vals ...string) map[string]bool {
	m := make(map[string]bool, len(vals))
	for _, v := range vals {
		m[v] = true
	}
	return m
}

// ---------- frontmatter ----------

type frontmatter struct {
	scalars map[string]string
	lists   map[string][]string
}

func (f frontmatter) get(k string) string { return f.scalars[k] }

func (f frontmatter) list(k string) []string {
	if v, ok := f.lists[k]; ok {
		return v
	}
	if s := f.scalars[k]; s != "" {
		return []string{s}
	}
	return nil
}

func parseFrontmatter(lines []string) frontmatter {
	fm := frontmatter{scalars: map[string]string{}, lists: map[string][]string{}}
	for i := 0; i < len(lines); i++ {
		line := strings.TrimRight(lines[i], "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		m := reKey.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		key, rest := m[1], strings.TrimSpace(m[2])

		if rest == "" {
			// Inline array on the next line?
			if i+1 < len(lines) && strings.HasPrefix(strings.TrimSpace(lines[i+1]), "[") {
				buf := strings.TrimSpace(lines[i+1])
				j := i + 1
				for !strings.Contains(buf, "]") && j+1 < len(lines) {
					j++
					buf += " " + strings.TrimSpace(lines[j])
				}
				fm.lists[key] = splitList(innerArray(buf))
				i = j
				continue
			}
			// Block list?
			var items []string
			j := i + 1
			for j < len(lines) {
				t := strings.TrimSpace(lines[j])
				if strings.HasPrefix(t, "- ") {
					items = append(items, unquote(strings.TrimSpace(strings.TrimPrefix(t, "- "))))
					j++
					continue
				}
				if t == "" {
					j++
					continue
				}
				break
			}
			if len(items) > 0 {
				fm.lists[key] = items
				i = j - 1
				continue
			}
			fm.scalars[key] = ""
			continue
		}

		if strings.HasPrefix(rest, "[") {
			buf := rest
			j := i
			for !strings.Contains(buf, "]") && j+1 < len(lines) {
				j++
				buf += " " + strings.TrimSpace(lines[j])
			}
			fm.lists[key] = splitList(innerArray(buf))
			i = j
			continue
		}

		fm.scalars[key] = unquote(stripComment(rest))
	}
	return fm
}

func innerArray(s string) string {
	open := strings.Index(s, "[")
	closeIdx := strings.LastIndex(s, "]")
	if open < 0 || closeIdx <= open {
		return ""
	}
	return s[open+1 : closeIdx]
}

func splitList(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var out []string
	for _, part := range strings.Split(s, ",") {
		v := unquote(strings.TrimSpace(part))
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func stripComment(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, `"`) {
		return s
	}
	if i := strings.Index(s, " #"); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return s
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// readDoc returns the parsed frontmatter, the full lines and the body (lines after frontmatter).
func readDoc(path string) (frontmatter, []string, []string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return frontmatter{}, nil, nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return frontmatter{}, lines, lines, nil
	}
	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end < 0 {
		return frontmatter{}, lines, lines, nil
	}
	return parseFrontmatter(lines[1:end]), lines, lines[end+1:], nil
}

// ---------- model ----------

type ticket struct {
	path string
	fm   frontmatter
}

func (t ticket) id() string         { return t.fm.get("id") }
func (t ticket) title() string      { return t.fm.get("title") }
func (t ticket) status() string     { return t.fm.get("status") }
func (t ticket) type_() string      { return t.fm.get("type") }
func (t ticket) plan() string       { return t.fm.get("plan") }
func (t ticket) deps() []string     { return t.fm.list("depends_on") }
func (t ticket) relatedADR() []string { return t.fm.list("related_adr") }
func (t ticket) finding() string    { return t.fm.get("finding") }

type plan struct {
	path    string
	dir     string
	fm      frontmatter
	tickets []ticket
}

func (p plan) id() string     { return p.fm.get("id") }
func (p plan) title() string  { return p.fm.get("title") }
func (p plan) kind() string   { return p.fm.get("kind") }
func (p plan) status() string { return p.fm.get("status") }

type adr struct {
	path string
	num  int
	fm   frontmatter
}

func (a adr) id() string            { return a.fm.get("id") }
func (a adr) title() string         { return a.fm.get("title") }
func (a adr) status() string        { return a.fm.get("status") }
func (a adr) date() string          { return a.fm.get("date") }
func (a adr) supersededBy() string  { return a.fm.get("superseded_by") }
func (a adr) supersedes() string    { return a.fm.get("supersedes") }

// ---------- load ----------

func loadTickets(planDir string) ([]ticket, error) {
	matches, err := filepath.Glob(filepath.Join(planDir, "ticket-*.md"))
	if err != nil {
		return nil, err
	}
	var out []ticket
	for _, m := range matches {
		fm, _, _, err := readDoc(m)
		if err != nil {
			return nil, err
		}
		out = append(out, ticket{path: m, fm: fm})
	}
	sort.Slice(out, func(i, j int) bool { return ticketNum(out[i].id()) < ticketNum(out[j].id()) })
	return out, nil
}

func loadPlans(repo string) ([]plan, error) {
	matches, err := filepath.Glob(filepath.Join(repo, "docs", "plans", "*", "README.MD"))
	if err != nil {
		return nil, err
	}
	var out []plan
	for _, m := range matches {
		fm, _, _, err := readDoc(m)
		if err != nil {
			return nil, err
		}
		dir := filepath.Dir(m)
		tk, err := loadTickets(dir)
		if err != nil {
			return nil, err
		}
		out = append(out, plan{path: m, dir: dir, fm: fm, tickets: tk})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].id() < out[j].id() })
	return out, nil
}

func loadADRs(repo string) ([]adr, error) {
	matches, err := filepath.Glob(filepath.Join(repo, "docs", "adr", "[0-9][0-9][0-9][0-9]-*.md"))
	if err != nil {
		return nil, err
	}
	var out []adr
	for _, m := range matches {
		fm, _, _, err := readDoc(m)
		if err != nil {
			return nil, err
		}
		out = append(out, adr{path: m, num: ticketNum(strings.TrimPrefix(fm.get("id"), "ADR-")), fm: fm})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].num < out[j].num })
	return out, nil
}

func ticketNum(id string) int {
	digits := regexp.MustCompile(`\d+`).FindString(id)
	n, _ := strconv.Atoi(digits)
	return n
}

// ---------- checker ----------

type checker struct{ errs []string }

func (c *checker) errf(format string, args ...any) {
	c.errs = append(c.errs, fmt.Sprintf(format, args...))
}

func (c *checker) ok() bool { return len(c.errs) == 0 }

func (c *checker) report() {
	if c.ok() {
		fmt.Println("workflow: OK")
		return
	}
	for _, e := range c.errs {
		fmt.Println("workflow: " + e)
	}
}

func rel(repo, p string) string {
	r, err := filepath.Rel(repo, p)
	if err != nil {
		return p
	}
	return filepath.ToSlash(r)
}

// ---------- validation ----------

func validateTicket(c *checker, repo string, p plan, t ticket, byID map[string]ticket) {
	loc := rel(repo, t.path)
	base := filepath.Base(t.path)
	m := reTicketFile.FindStringSubmatch(base)
	if m == nil {
		c.errf("%s: file name must match ticket-NNN.md", loc)
	}
	require := func(field, v string) {
		if strings.TrimSpace(v) == "" {
			c.errf("%s: missing required field %q", loc, field)
		}
	}
	require("id", t.id())
	require("title", t.title())
	require("plan", t.plan())
	require("status", t.status())
	require("type", t.type_())
	require("priority", t.fm.get("priority"))
	require("estimate", t.fm.get("estimate"))
	require("created", t.fm.get("created"))
	require("updated", t.fm.get("updated"))

	if t.id() != "" && base != t.id()+".md" {
		c.errf("%s: id %q does not match file name", loc, t.id())
	}
	if !reTicketID.MatchString(t.id()) {
		c.errf("%s: id %q must look like ticket-NNN", loc, t.id())
	}
	if t.plan() != "" && t.plan() != p.id() {
		c.errf("%s: plan %q does not match enclosing plan %q", loc, t.plan(), p.id())
	}
	if t.status() != "" && !ticketStatuses[t.status()] {
		c.errf("%s: invalid status %q", loc, t.status())
	}
	if t.type_() != "" && !ticketTypes[t.type_()] {
		c.errf("%s: invalid type %q", loc, t.type_())
	}
	if v := t.fm.get("priority"); v != "" && !priorities[v] {
		c.errf("%s: invalid priority %q", loc, v)
	}
	if v := t.fm.get("estimate"); v != "" && !estimates[v] {
		c.errf("%s: invalid estimate %q", loc, v)
	}
	for _, f := range []string{"created", "updated"} {
		if v := t.fm.get(f); v != "" && !reDate.MatchString(v) {
			c.errf("%s: %s %q is not YYYY-MM-DD", loc, f, v)
		}
	}
	for _, dep := range t.deps() {
		if !reTicketID.MatchString(dep) {
			c.errf("%s: depends_on %q is not a ticket id", loc, dep)
			continue
		}
		if dep == t.id() {
			c.errf("%s: depends_on cannot reference itself", loc)
		}
		if _, ok := byID[dep]; !ok {
			c.errf("%s: depends_on references unknown ticket %q", loc, dep)
		}
	}
	for _, a := range t.relatedADR() {
		if !reADRID.MatchString(a) {
			c.errf("%s: related_adr %q is not an ADR id", loc, a)
		}
	}
	if t.type_() == "spike" && t.status() == "done" {
		finding := t.finding()
		if len(t.relatedADR()) == 0 && finding == "" {
			c.errf("%s: a done spike must link an ADR (related_adr) or a finding", loc)
		}
		if finding != "" {
			if _, err := os.Stat(filepath.Join(repo, finding)); err != nil {
				c.errf("%s: finding %q does not exist", loc, finding)
			}
		}
	}

	raw, _ := os.ReadFile(t.path)
	for _, tok := range placeholderTokens {
		if strings.Contains(string(raw), tok) {
			c.errf("%s: leftover template placeholder %q", loc, tok)
		}
	}
	body := string(raw)
	for _, section := range ticketSections {
		if !strings.Contains(body, section) {
			c.errf("%s: missing section %q", loc, section)
		}
	}
}

func validatePlan(c *checker, repo string, p plan) {
	loc := rel(repo, p.path)
	require := func(field, v string) {
		if strings.TrimSpace(v) == "" {
			c.errf("%s: missing required field %q", loc, field)
		}
	}
	require("id", p.id())
	require("title", p.title())
	require("kind", p.kind())
	require("status", p.status())
	require("owner", p.fm.get("owner"))
	require("created", p.fm.get("created"))
	require("updated", p.fm.get("updated"))

	if p.id() != "" && p.id() != filepath.Base(p.dir) {
		c.errf("%s: id %q does not match directory name", loc, p.id())
	}
	if p.kind() != "" && !planKinds[p.kind()] {
		c.errf("%s: invalid kind %q", loc, p.kind())
	}
	if p.status() != "" && !planStatuses[p.status()] {
		c.errf("%s: invalid status %q", loc, p.status())
	}

	expected := renderPlanTickets(p.tickets)
	if got, ok := sectionText(p.path, "## Tickets"); ok {
		if normalize(got) != normalize(expected) {
			c.errf("%s: ticket index is out of sync; run `task plan:sync`", loc)
		}
	} else {
		c.errf("%s: missing `## Tickets` section", loc)
	}
}

func validateADR(c *checker, repo string, a adr, all map[string]adr) {
	loc := rel(repo, a.path)
	base := filepath.Base(a.path)
	m := reADRFile.FindStringSubmatch(base)
	if m == nil {
		c.errf("%s: file name must match NNNN-slug.md", loc)
	}
	require := func(field, v string) {
		if strings.TrimSpace(v) == "" {
			c.errf("%s: missing required field %q", loc, field)
		}
	}
	require("id", a.id())
	require("title", a.title())
	require("status", a.status())
	require("date", a.date())

	if !reADRID.MatchString(a.id()) {
		c.errf("%s: id %q must look like ADR-NNNN", loc, a.id())
	} else if m != nil && !strings.HasPrefix(a.id(), "ADR-"+m[1]) {
		c.errf("%s: id %q does not match file number", loc, a.id())
	}
	if a.status() != "" && !adrStatuses[a.status()] {
		c.errf("%s: invalid status %q", loc, a.status())
	}
	if a.date() != "" && !reDate.MatchString(a.date()) {
		c.errf("%s: date %q is not YYYY-MM-DD", loc, a.date())
	}
	if sb := a.supersededBy(); sb != "" {
		if _, ok := all[sb]; !ok {
			c.errf("%s: superseded_by references unknown %q", loc, sb)
		}
		if a.status() != "superseded" {
			c.errf("%s: superseded_by is set but status is %q", loc, a.status())
		}
	}
	if a.status() == "superseded" && a.supersededBy() == "" {
		c.errf("%s: status is superseded but superseded_by is empty", loc)
	}
	if sp := a.supersedes(); sp != "" {
		if _, ok := all[sp]; !ok {
			c.errf("%s: supersedes references unknown %q", loc, sp)
		}
	}
}

func validateADRs(repo string, c *checker) {
	adrs, err := loadADRs(repo)
	if err != nil {
		c.errf("loading ADRs: %v", err)
		return
	}
	all := map[string]adr{}
	for _, a := range adrs {
		all[a.id()] = a
	}
	for _, a := range adrs {
		validateADR(c, repo, a, all)
	}
	indexPath := filepath.Join(repo, "docs", "adr", "README.md")
	raw, err := os.ReadFile(indexPath)
	if err != nil {
		c.errf("docs/adr/README.md missing; run `task adr:index`")
		return
	}
	if normalize(string(raw)) != normalize(renderADRIndex(adrs)) {
		c.errf("docs/adr/README.md is out of sync; run `task adr:index`")
	}
}

func validateDiaries(repo string, c *checker, plans map[string]plan) {
	matches, _ := filepath.Glob(filepath.Join(repo, "docs", "diary", "entry-*.md"))
	for _, m := range matches {
		loc := rel(repo, m)
		fm, _, _, err := readDoc(m)
		if err != nil {
			c.errf("%s: %v", loc, err)
			continue
		}
		for _, f := range []string{"id", "date", "session", "plan", "status"} {
			if strings.TrimSpace(fm.get(f)) == "" {
				c.errf("%s: missing required field %q", loc, f)
			}
		}
		if base := filepath.Base(m); !reDiaryFile.MatchString(base) {
			c.errf("%s: file name must match entry-YYYY-MM-DD-slug.md", loc)
		} else if fm.get("date") != "" {
			if mm := reDiaryFile.FindStringSubmatch(base); mm != nil && mm[1] != fm.get("date") {
				c.errf("%s: date %q does not match file name", loc, fm.get("date"))
			}
		}
		if !diaryStatuses[fm.get("status")] && fm.get("status") != "" {
			c.errf("%s: invalid status %q", loc, fm.get("status"))
		}
		if pl := fm.get("plan"); pl != "" {
			if _, ok := plans[pl]; !ok {
				c.errf("%s: plan %q does not exist", loc, pl)
			}
		}
	}
}

func detectCycles(c *checker, repo string, plans []plan) {
	for _, p := range plans {
		byID := map[string]ticket{}
		for _, t := range p.tickets {
			byID[t.id()] = t
		}
		state := map[string]int{} // 0 unvisited, 1 visiting, 2 done
		var visit func(id string, stack []string)
		visit = func(id string, stack []string) {
			if state[id] == 2 {
				return
			}
			if state[id] == 1 {
				c.errf("%s: dependency cycle: %s -> %s", rel(repo, p.path), strings.Join(stack, " -> "), id)
				return
			}
			state[id] = 1
			t, ok := byID[id]
			if !ok {
				state[id] = 2
				return
			}
			for _, d := range t.deps() {
				visit(d, append(stack, id))
			}
			state[id] = 2
		}
		for _, t := range p.tickets {
			visit(t.id(), nil)
		}
	}
}

func runVerify(repo string) int {
	c := &checker{}
	plans, err := loadPlans(repo)
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	planByID := map[string]plan{}
	for _, p := range plans {
		planByID[p.id()] = p
	}
	for _, p := range plans {
		validatePlan(c, repo, p)
		byID := map[string]ticket{}
		for _, t := range p.tickets {
			byID[t.id()] = t
		}
		for _, t := range p.tickets {
			validateTicket(c, repo, p, t, byID)
		}
	}
	detectCycles(c, repo, plans)
	validateADRs(repo, c)
	validateDiaries(repo, c, planByID)
	c.report()
	if c.ok() {
		return 0
	}
	return 1
}

// ---------- index generation ----------

func renderPlanTickets(tickets []ticket) string {
	var b strings.Builder
	b.WriteString("## Tickets\n\n")
	b.WriteString("| ID | Title | Status | Depends on |\n")
	b.WriteString("|----|-------|--------|------------|\n")
	for _, t := range tickets {
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n", t.id(), t.title(), t.status(), strings.Join(t.deps(), ", "))
	}
	return b.String()
}

func renderADRIndex(adrs []adr) string {
	var b strings.Builder
	b.WriteString("# Architecture Decision Records\n\n")
	b.WriteString("| ID | Title | Status | Date | Superseded by |\n")
	b.WriteString("|----|-------|--------|------|---------------|\n")
	for _, a := range adrs {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n", a.id(), a.title(), a.status(), a.date(), a.supersededBy())
	}
	return b.String()
}

// sectionText returns the text of a "## Heading" section up to the next "## ".
func sectionText(path, heading string) (string, bool) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	doc := strings.ReplaceAll(string(raw), "\r\n", "\n")
	marker := "\n" + heading + "\n"
	i := strings.Index(doc, marker)
	var start int
	if i < 0 {
		if strings.HasPrefix(doc, heading+"\n") {
			start = 0
		} else {
			return "", false
		}
	} else {
		start = i + 1
	}
	rest := doc[start+len(heading+"\n"):]
	end := len(doc)
	if j := strings.Index(rest, "\n## "); j >= 0 {
		end = start + len(heading+"\n") + j
	}
	return doc[start:end], true
}

func replaceSection(path, heading, section string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	doc := strings.ReplaceAll(string(raw), "\r\n", "\n")
	marker := "\n" + heading + "\n"
	i := strings.Index(doc, marker)
	start := 0
	if i < 0 {
		if strings.HasPrefix(doc, heading+"\n") {
			start = 0
			marker = heading + "\n"
		} else {
			return fmt.Errorf("%s: section %q not found", path, heading)
		}
	} else {
		start = i + 1
	}
	rest := doc[start+len(heading+"\n"):]
	var end int
	if j := strings.Index(rest, "\n## "); j >= 0 {
		end = start + len(heading+"\n") + j
	} else {
		end = len(doc)
	}
	out := doc[:start] + strings.TrimRight(section, "\n") + "\n" + doc[end:]
	return os.WriteFile(path, []byte(out), 0o644)
}

func normalize(s string) string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		out = append(out, strings.TrimRight(line, " \t"))
	}
	return strings.TrimRight(strings.Join(out, "\n"), "\n")
}

func runPlanSync(repo string, only string) int {
	plans, err := loadPlans(repo)
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	for _, p := range plans {
		if only != "" && p.id() != only {
			continue
		}
		if err := replaceSection(p.path, "## Tickets", renderPlanTickets(p.tickets)); err != nil {
			fmt.Println("workflow: " + err.Error())
			return 1
		}
		fmt.Println("synced " + rel(repo, p.path))
	}
	return 0
}

func runADRIndex(repo string) int {
	adrs, err := loadADRs(repo)
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	path := filepath.Join(repo, "docs", "adr", "README.md")
	if err := os.WriteFile(path, []byte(renderADRIndex(adrs)), 0o644); err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	fmt.Println("wrote " + rel(repo, path))
	return 0
}

// ---------- lifecycle ----------

func findTicket(repo, id, planName string) (ticket, error) {
	plans, err := loadPlans(repo)
	if err != nil {
		return ticket{}, err
	}
	var found []ticket
	for _, p := range plans {
		if planName != "" && p.id() != planName {
			continue
		}
		for _, t := range p.tickets {
			if t.id() == id {
				found = append(found, t)
			}
		}
	}
	if len(found) == 0 {
		return ticket{}, fmt.Errorf("ticket %q not found", id)
	}
	if len(found) > 1 {
		return ticket{}, fmt.Errorf("ticket %q is ambiguous; pass a plan", id)
	}
	return found[0], nil
}

func setField(path, key, value string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n")
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			break
		}
		if m := reKey.FindStringSubmatch(lines[i]); m != nil && m[1] == key {
			lines[i] = key + ": " + value
			return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644)
		}
	}
	return fmt.Errorf("%s: field %q not found in frontmatter", path, key)
}

func runTicketStart(repo, id, planName string) int {
	t, err := findTicket(repo, id, planName)
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	if s := t.status(); s != "todo" && s != "blocked" {
		fmt.Printf("workflow: %s is %q; only todo/blocked can start\n", id, s)
		return 1
	}
	plans, _ := loadPlans(repo)
	byID := map[string]ticket{}
	for _, p := range plans {
		if p.id() != t.plan() {
			continue
		}
		for _, x := range p.tickets {
			byID[x.id()] = x
		}
	}
	for _, d := range t.deps() {
		if dep, ok := byID[d]; !ok || dep.status() != "done" {
			fmt.Printf("workflow: cannot start %s: dependency %s is not done\n", id, d)
			return 1
		}
	}
	today := time.Now().Format("2006-01-02")
	if err := setField(t.path, "status", "in_progress"); err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	_ = setField(t.path, "updated", today)
	if code := runPlanSync(repo, t.plan()); code != 0 {
		return code
	}
	fmt.Printf("started %s\n", rel(repo, t.path))
	return 0
}

func runTicketClose(repo, id, planName string) int {
	t, err := findTicket(repo, id, planName)
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	today := time.Now().Format("2006-01-02")
	if err := setField(t.path, "status", "done"); err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	_ = setField(t.path, "updated", today)
	if code := runPlanSync(repo, t.plan()); code != 0 {
		return code
	}
	fmt.Printf("closed %s\n\nCommit with:\n  git add -A && git commit -m \"type(scope): summary (ticket-NNN)\"\n", rel(repo, t.path))
	return 0
}

// ---------- commit message ----------

func validateCommitMessage(msg string) []string {
	lines := strings.Split(strings.ReplaceAll(msg, "\r\n", "\n"), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) == "" {
		return []string{"commit message is empty"}
	}
	subject := strings.TrimSpace(lines[0])
	lower := strings.ToLower(subject)
	if strings.HasPrefix(lower, "merge ") || strings.HasPrefix(lower, "revert ") ||
		strings.HasPrefix(subject, "fixup!") || strings.HasPrefix(subject, "squash!") {
		return nil
	}
	var errs []string
	if !reRegular.MatchString(subject) {
		errs = append(errs, "subject must be Conventional Commits: type(scope): summary")
	}
	if !reTicketRef.MatchString(subject) && !strings.Contains(msg, "[no-ticket]") {
		errs = append(errs, "subject must reference a ticket, e.g. (ticket-010), or include [no-ticket]")
	}
	return errs
}

func runCommitMsg(path string) int {
	var raw []byte
	var err error
	if path == "-" {
		raw, err = io.ReadAll(os.Stdin)
	} else {
		raw, err = os.ReadFile(path)
	}
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		return 1
	}
	if errs := validateCommitMessage(string(raw)); len(errs) > 0 {
		for _, e := range errs {
			fmt.Println("workflow: " + e)
		}
		return 1
	}
	return 0
}

func runVerifyCommits(repo, rng string) int {
	if rng == "" {
		rng = "-1"
	}
	out, err := exec.Command("git", "-C", repo, "log", "--format=%H%x09%s", rng).Output()
	if err != nil {
		fmt.Println("workflow: git log failed: " + err.Error())
		return 1
	}
	fail := 0
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "\t", 2)
		hash, subj := parts[0], ""
		if len(parts) > 1 {
			subj = parts[1]
		}
		if errs := validateCommitMessage(subj); len(errs) > 0 {
			fail++
			for _, e := range errs {
				fmt.Printf("workflow: %.8s: %s\n", hash, e)
			}
		}
	}
	if fail > 0 {
		return 1
	}
	fmt.Println("workflow: commits OK")
	return 0
}

// ---------- entry ----------

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "docs", "adr")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("repository root not found")
		}
		dir = parent
	}
}

func usage() {
	fmt.Println(`workflow - repository workflow tool

Usage:
  workflow verify
  workflow plan-sync [plan]
  workflow adr-index
  workflow ticket-start <id> [plan]
  workflow ticket-close <id> [plan]
  workflow commit-msg <path|-> 
  workflow verify-commits [range]`)
}

func main() {
	repo, err := findRepoRoot()
	if err != nil {
		fmt.Println("workflow: " + err.Error())
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	code := 0
	switch os.Args[1] {
	case "verify":
		code = runVerify(repo)
	case "plan-sync":
		only := ""
		if len(os.Args) > 2 {
			only = os.Args[2]
		}
		code = runPlanSync(repo, only)
	case "adr-index":
		code = runADRIndex(repo)
	case "ticket-start":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		planName := ""
		if len(os.Args) > 3 {
			planName = os.Args[3]
		}
		code = runTicketStart(repo, os.Args[2], planName)
	case "ticket-close":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		planName := ""
		if len(os.Args) > 3 {
			planName = os.Args[3]
		}
		code = runTicketClose(repo, os.Args[2], planName)
	case "commit-msg":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		code = runCommitMsg(os.Args[2])
	case "verify-commits":
		rng := ""
		if len(os.Args) > 2 {
			rng = os.Args[2]
		}
		code = runVerifyCommits(repo, rng)
	default:
		usage()
		os.Exit(1)
	}
	os.Exit(code)
}
