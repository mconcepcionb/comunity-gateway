package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestParseFrontmatterScalarsListsAndMultiline(t *testing.T) {
	lines := []string{
		"id: ticket-001",
		"depends_on: [ticket-002, ticket-003]",
		"related_adrs:",
		"  [ADR-0001, ADR-0002,",
		"   ADR-0003]",
		`title: "Hello: world"`,
		"empty:",
		"tags:",
		"  - a",
		"  - b",
		"finding:",
	}
	fm := parseFrontmatter(lines)

	if got := fm.get("id"); got != "ticket-001" {
		t.Fatalf("id = %q", got)
	}
	if got := fm.get("title"); got != "Hello: world" {
		t.Fatalf("title = %q", got)
	}
	if got := fm.list("depends_on"); strings.Join(got, ",") != "ticket-002,ticket-003" {
		t.Fatalf("depends_on = %v", got)
	}
	if got := fm.list("related_adrs"); strings.Join(got, ",") != "ADR-0001,ADR-0002,ADR-0003" {
		t.Fatalf("related_adrs = %v", got)
	}
	if got := fm.list("tags"); strings.Join(got, ",") != "a,b" {
		t.Fatalf("tags = %v", got)
	}
	if got := fm.get("finding"); got != "" {
		t.Fatalf("finding = %q, want empty", got)
	}
}

func TestSplitList(t *testing.T) {
	got := splitList("ADR-0001, ADR-0002 ,ADR-0003")
	if strings.Join(got, "|") != "ADR-0001|ADR-0002|ADR-0003" {
		t.Fatalf("splitList = %v", got)
	}
	if len(splitList("  ")) != 0 {
		t.Fatal("expected empty list")
	}
}

func TestValidateCommitMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"valid with scope", "feat(games): add adapter (ticket-010)", false},
		{"valid no scope", "docs: update readme (ticket-019)", false},
		{"no ticket ref", "chore: bootstrap the repo", true},
		{"no-ticket escape hatch", "chore: bootstrap the repo [no-ticket]", false},
		{"not conventional", "updated stuff (ticket-001)", true},
		{"merge", "Merge branch 'main'", false},
		{"revert", "Revert \"feat: x (ticket-001)\"", false},
		{"empty", "   ", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validateCommitMessage(tt.msg)
			if (len(errs) > 0) != tt.wantErr {
				t.Fatalf("validateCommitMessage(%q) errs=%v wantErr=%v", tt.msg, errs, tt.wantErr)
			}
		})
	}
}

func TestRenderPlanTickets(t *testing.T) {
	tickets := []ticket{
		{fm: frontmatter{scalars: map[string]string{"id": "ticket-001", "title": "First", "status": "done"}, lists: map[string][]string{}}},
		{fm: frontmatter{scalars: map[string]string{"id": "ticket-002", "title": "Second", "status": "todo"}, lists: map[string][]string{"depends_on": {"ticket-001"}}}},
	}
	out := renderPlanTickets(tickets)
	for _, want := range []string{"| ticket-001 | First | done |  |", "| ticket-002 | Second | todo | ticket-001 |"} {
		if !strings.Contains(out, want) {
			t.Fatalf("renderPlanTickets missing %q in:\n%s", want, out)
		}
	}
}

func TestNormalizeTrimsTrailingWhitespaceAndNewlines(t *testing.T) {
	if normalize("a  \nb\t\n\n") != "a\nb" {
		t.Fatalf("normalize = %q", normalize("a  \nb\t\n\n"))
	}
}

func TestCoveragePercent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "coverage.out")
	content := "mode: atomic\n" +
		"example.com/a.go:1.1,1.10 2 1\n" +
		"example.com/a.go:2.1,2.10 2 0\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	pct, err := coveragePercent(path)
	if err != nil {
		t.Fatal(err)
	}
	if pct != 50 {
		t.Fatalf("coveragePercent = %v, want 50", pct)
	}
}

func TestNextCalVer(t *testing.T) {
	sep := time.Date(2026, 9, 15, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		last string
		want string
	}{
		{"", "v2026.09.0"},
		{"v2026.09.3", "v2026.09.4"},
		{"v2026.08.2", "v2026.09.0"},
		{"v2025.09.9", "v2026.09.0"},
	}
	for _, tt := range tests {
		if got := nextCalVer(tt.last, sep); got != tt.want {
			t.Fatalf("nextCalVer(%q) = %q, want %q", tt.last, got, tt.want)
		}
	}
}
