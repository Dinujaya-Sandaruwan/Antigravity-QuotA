package ui

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"antigravity-quota-tui/internal/api"
	"antigravity-quota-tui/internal/config"
)

// ═══════════════════════════════════════════════════════════════════════════════
// CATEGORIES
// ═══════════════════════════════════════════════════════════════════════════════

const (
	catClaude    = "Claude Models"
	catGemini    = "Gemini Models"
	catFlashLite = "Flash Lite Models"
)

var catOrder = []string{catClaude, catGemini, catFlashLite}

// ═══════════════════════════════════════════════════════════════════════════════
// COLOR PALETTE
// ═══════════════════════════════════════════════════════════════════════════════

const (
	cSurface  = "235"
	cBorder   = "238"
	cBorderHi = "63"
	cText     = "252"
	cMuted    = "242"
	cFaint    = "238"
	cGood     = "84"
	cWarn     = "215"
	cBad      = "203"
	cFetch    = "220"
	cError    = "210"
	cFooter   = "240"
	cScroll   = "63"
	cLogo     = "99"
	cBadgeOk  = "35"
	cBadgeWt  = "208"
	cClaude   = "204"
	cGemini   = "39"
	cFlash    = "220"
	cSep      = "240"
	cAccent   = "105"
	cAccent2  = "75"
	cDiscl    = "245"
)

// ═══════════════════════════════════════════════════════════════════════════════
// STYLES
// ═══════════════════════════════════════════════════════════════════════════════

var (
	sHeaderBg = lipgloss.NewStyle().Background(lipgloss.Color(cSurface)).Padding(0, 1)
	sLogo     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cLogo))
	sAppName  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cText))
	sVer      = lipgloss.NewStyle().Foreground(lipgloss.Color(cMuted))
	sFetch    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cFetch))
	sHintKey  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cAccent2))
	sHint     = lipgloss.NewStyle().Foreground(lipgloss.Color(cMuted))
	sDivHi    = lipgloss.NewStyle().Foreground(lipgloss.Color(cBorderHi))
	sDivLo    = lipgloss.NewStyle().Foreground(lipgloss.Color(cBorder))
	sSepStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(cSep))

	sSection = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cAccent)).
			PaddingLeft(1).BorderLeft(true).BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color(cBorderHi))

	sClaude = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cClaude))
	sGemini = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cGemini))
	sFlash  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cFlash))

	sDiscl   = lipgloss.NewStyle().Foreground(lipgloss.Color(cDiscl))
	sColHdr  = lipgloss.NewStyle().Foreground(lipgloss.Color(cFaint))
	sEmail   = lipgloss.NewStyle().Foreground(lipgloss.Color(cText))
	sReset   = lipgloss.NewStyle().Foreground(lipgloss.Color(cMuted))
	sPctGood = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cGood))
	sPctWarn = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cWarn))
	sPctBad  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cBad))
	sBadgeOk = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color(cBadgeOk)).Padding(0, 1)
	sBadgeWt = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color(cBadgeWt)).Padding(0, 1)
	sError    = lipgloss.NewStyle().Foreground(lipgloss.Color(cError))
	sFooter   = lipgloss.NewStyle().Foreground(lipgloss.Color(cFooter))
	sMuted    = lipgloss.NewStyle().Foreground(lipgloss.Color(cMuted))
	sCacheMod = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cAccent2))
	sCacheCat = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cAccent2))
	sModalBdr = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(cBorderHi)).Padding(1, 3)
	sModalHnt = lipgloss.NewStyle().Foreground(lipgloss.Color(cFaint)).Italic(true)
)

func catStyle(cat string) lipgloss.Style {
	switch cat {
	case catClaude:
		return sClaude
	case catGemini:
		return sGemini
	case catFlashLite:
		return sFlash
	default:
		return sGemini
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TYPES
// ═══════════════════════════════════════════════════════════════════════════════

type categorizedGroup struct {
	category string
	groups   []api.ModelGroup
}

type clickZone struct {
	category  string
	titleLine int
	models    []string
}

type leftPanelResult struct {
	lines      []string
	clickZones []clickZone
}

type fetchDoneMsg struct {
	groups []api.ModelGroup
	errors []string
	took   time.Duration
}

// ═══════════════════════════════════════════════════════════════════════════════
// MODEL
// ═══════════════════════════════════════════════════════════════════════════════

type Model struct {
	Accounts    []config.Account
	groups      []api.ModelGroup
	fetchErrors []string
	loading     bool
	lastFetch   time.Time
	fetchTook   time.Duration

	width        int
	height       int
	scrollOffset int

	categories []categorizedGroup
	clickZones []clickZone

	modalOpen     bool
	modalCategory string
	modalModels   []string
}

// ═══════════════════════════════════════════════════════════════════════════════
// INIT / UPDATE / VIEW
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) Init() tea.Cmd { return doFetch(m.Accounts) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.modalOpen {
			m.modalOpen = false
			return m, nil
		}
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "r":
			if !m.loading {
				m.loading = true
				return m, doFetch(m.Accounts)
			}
		case "up", "k":
			if m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down", "j":
			m.scrollOffset++
		case "pgup":
			m.scrollOffset -= m.height / 2
			if m.scrollOffset < 0 {
				m.scrollOffset = 0
			}
		case "pgdn":
			m.scrollOffset += m.height / 2
		}

	case tea.MouseMsg:
		if m.modalOpen {
			if msg.Action == tea.MouseActionPress {
				m.modalOpen = false
			}
			return m, nil
		}
		switch msg.Button {
		case tea.MouseButtonWheelUp:
			m.scrollOffset -= 3
			if m.scrollOffset < 0 {
				m.scrollOffset = 0
			}
		case tea.MouseButtonWheelDown:
			m.scrollOffset += 3
		case tea.MouseButtonLeft:
			if msg.Action == tea.MouseActionPress {
				if cz, ok := m.findClickZone(msg.X, msg.Y); ok {
					m.modalOpen = true
					m.modalCategory = cz.category
					m.modalModels = cz.models
				}
			}
		}

	case fetchDoneMsg:
		m.loading = false
		m.groups = msg.groups
		m.fetchErrors = msg.errors
		m.lastFetch = time.Now()
		m.fetchTook = msg.took
		m.scrollOffset = 0
		m.categories = categorizeGroups(msg.groups)
		// Pre-compute click zones (line numbers are width-independent)
		result := buildLeftPanel(m.categories, 80)
		m.clickZones = result.clickZones
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}
	if m.width < 50 {
		return sMuted.Render("Terminal too narrow. Resize to at least 50 columns.")
	}

	// Modal overlay (model list popup) takes the full screen
	if m.modalOpen {
		return m.viewModal()
	}

	// Loading overlay — shown on initial fetch AND on R refresh
	if m.loading {
		return m.viewLoading()
	}

	w := m.width

	// ── Header (2 rows, sticky) ──────────────────────────────────────────────
	header := m.viewHeader(w)
	hdrH := 2

	// ── Footer (sticky bottom) ───────────────────────────────────────────────
	footer := m.viewFooter(w)
	ftrH := m.footerLineCount()

	// ── Scrollable body ──────────────────────────────────────────────────────
	bodyH := m.height - hdrH - ftrH
	if bodyH < 1 {
		bodyH = 1
	}

	// No data yet?
	if m.categories == nil {
		var msg string
		if m.lastFetch.IsZero() {
			msg = sMuted.Render("  Fetching data for all accounts…")
		} else {
			msg = sMuted.Render("  No data. Press ") + sHintKey.Render("R") + sMuted.Render(" to refresh.")
		}
		bodyLines := make([]string, bodyH)
		bodyLines[0] = ""
		if bodyH > 1 {
			bodyLines[1] = msg
		}
		return header + strings.Join(bodyLines, "\n") + "\n" + footer
	}

	// Build left & right panels
	leftW := m.leftW()
	sepW := 3 // " │ "
	rightW := w - leftW - sepW
	if rightW < 10 {
		rightW = 10
	}

	lp := buildLeftPanel(m.categories, leftW)
	leftLines := lp.lines
	rightLines := buildRightPanel(m.Accounts, rightW)

	// Pad to same height
	maxH := max(len(leftLines), len(rightLines))
	for len(leftLines) < maxH {
		leftLines = append(leftLines, "")
	}
	for len(rightLines) < maxH {
		rightLines = append(rightLines, "")
	}

	// Join side-by-side with separator
	sep := sSepStyle.Render("│")
	bodyLines := make([]string, maxH)
	for i := 0; i < maxH; i++ {
		l := fitWidth(leftLines[i], leftW)
		r := fitWidth(rightLines[i], rightW)
		bodyLines[i] = l + " " + sep + " " + r
	}

	// Apply scroll
	maxScroll := len(bodyLines) - bodyH
	if maxScroll < 0 {
		maxScroll = 0
	}
	scroll := m.scrollOffset
	if scroll > maxScroll {
		scroll = maxScroll
	}
	if scroll < 0 {
		scroll = 0
	}
	end := scroll + bodyH
	if end > len(bodyLines) {
		end = len(bodyLines)
	}
	visible := strings.Join(bodyLines[scroll:end], "\n")

	return header + visible + "\n" + footer
}

// ═══════════════════════════════════════════════════════════════════════════════
// HEADER
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) viewHeader(w int) string {
	logo := sLogo.Render("◈")
	name := sAppName.Render(" Antigravity Quota ")
	ver := sVer.Render("v1.0")

	var status string
	if m.loading {
		status = "  " + sFetch.Render("⟳ fetching…")
	} else if !m.lastFetch.IsZero() {
		status = "  " + sMuted.Render(m.lastFetch.Format("15:04:05"))
	}

	left := logo + name + ver + status
	leftW := lipgloss.Width(left)

	hints := viewHints()
	hintsW := lipgloss.Width(hints)

	gap := w - leftW - hintsW - 2
	if gap < 1 {
		gap = 1
	}

	row1 := sHeaderBg.Width(w).Render(left + strings.Repeat(" ", gap) + hints)
	row2 := sDivHi.Render(strings.Repeat("━", w))
	return row1 + "\n" + row2 + "\n"
}

func viewHints() string {
	entries := []struct{ k, v string }{
		{"R", "refresh"}, {"↑↓", "scroll"}, {"Q", "quit"},
	}
	var parts []string
	for _, e := range entries {
		parts = append(parts, sHintKey.Render(e.k)+" "+sHint.Render(e.v))
	}
	return strings.Join(parts, sHint.Render(" · "))
}

// ═══════════════════════════════════════════════════════════════════════════════
// FOOTER
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) viewFooter(w int) string {
	var sb strings.Builder
	sb.WriteString(sDivHi.Render(strings.Repeat("━", w)) + "\n")

	if len(m.fetchErrors) > 0 {
		for _, e := range m.fetchErrors {
			sb.WriteString(sError.Render(" ✗  "+e) + "\n")
		}
	}

	if !m.lastFetch.IsZero() {
		sb.WriteString(sFooter.Render(fmt.Sprintf(
			" Updated %s · %s · %d accounts",
			m.lastFetch.Format("Mon 15:04:05"),
			m.fetchTook.Round(time.Millisecond),
			len(m.Accounts),
		)) + "\n")
	} else if m.loading {
		sb.WriteString(sFooter.Render(" Loading…") + "\n")
	}

	return sb.String()
}

func (m Model) footerLineCount() int {
	h := 1 // divider
	if len(m.fetchErrors) > 0 {
		h += len(m.fetchErrors)
	}
	if !m.lastFetch.IsZero() {
		h++
	} else if m.loading {
		h++
	}
	return h
}

// ═══════════════════════════════════════════════════════════════════════════════
// LOADING OVERLAY
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) viewLoading() string {
	// Spinner frames cycle on each render — we use a fixed char since we
	// have no tick, but the bold box is enough visual weight.
	spinnerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(cFetch))

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(cText))

	subStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cMuted))

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(cFetch)).
		Background(lipgloss.Color(cSurface)).
		Padding(2, 6).
		Align(lipgloss.Center)

	n := len(m.Accounts)
	accountWord := "account"
	if n != 1 {
		accountWord = "accounts"
	}

	content := spinnerStyle.Render("⟳  Fetching Data") + "\n\n" +
		titleStyle.Render("Querying Antigravity API") + "\n" +
		subStyle.Render(fmt.Sprintf("Refreshing quota for %d %s…", n, accountWord)) + "\n\n" +
		subStyle.Render("Please wait, this may take a few seconds.")

	box := boxStyle.Render(content)

	// Place the header + centered box + hint
	header := m.viewHeader(m.width)
	hdrH := 2
	bodyH := m.height - hdrH
	if bodyH < 1 {
		bodyH = 1
	}

	hint := subStyle.Render("Press Q to quit")
	hintLine := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, hint)

	// Center the box in remaining space (leave 1 row for hint at bottom)
	centered := lipgloss.Place(m.width, bodyH-1, lipgloss.Center, lipgloss.Center, box)

	return header + centered + "\n" + hintLine
}

// ═══════════════════════════════════════════════════════════════════════════════
// MODAL (model list popup)
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) viewModal() string {
	style := catStyle(m.modalCategory)

	var sb strings.Builder
	sb.WriteString(style.Render("◆  "+m.modalCategory) + "\n")

	maxLabelW := 0
	for _, model := range m.modalModels {
		if len(model) > maxLabelW {
			maxLabelW = len(model)
		}
	}
	divW := maxLabelW + 6
	if divW < 30 {
		divW = 30
	}
	sb.WriteString(sDivLo.Render(strings.Repeat("─", divW)) + "\n\n")

	for i, model := range m.modalModels {
		num := sHint.Render(fmt.Sprintf(" %2d. ", i+1))
		name := sEmail.Render(model)
		sb.WriteString(num + name + "\n")
	}

	sb.WriteString("\n")
	sb.WriteString(sModalHnt.Render("  Press any key or click to close"))

	box := sModalBdr.Render(sb.String())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

// ═══════════════════════════════════════════════════════════════════════════════
// LEFT PANEL — API Quota
// ═══════════════════════════════════════════════════════════════════════════════

func buildLeftPanel(cats []categorizedGroup, w int) leftPanelResult {
	var r leftPanelResult
	nl := func() { r.lines = append(r.lines, "") } // blank line
	emit := func(s string) { r.lines = append(r.lines, s) }

	nl()
	emit(" " + sSection.Render(" API Quota "))
	nl()

