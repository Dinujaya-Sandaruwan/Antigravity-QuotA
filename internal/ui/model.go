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

	for ci, cat := range cats {
		// Collect all model labels for the modal
		var labels []string
		for _, g := range cat.groups {
			labels = append(labels, g.Labels...)
		}

		// Compute overall average percentage across ALL accounts in this category
		var sum float64
		var count int
		for _, g := range cat.groups {
			for _, acc := range g.Accounts {
				sum += acc.Percentage
				count++
			}
		}
		var avgPct float64
		if count > 0 {
			avgPct = sum / float64(count)
		}

		// ── Category title ────────────────────────────────────────────────
		style := catStyle(cat.category)
		title := style.Render("  ◆  " + cat.category)
		discl := sDiscl.Render("▸ models")
		gap := w - lipgloss.Width(title) - lipgloss.Width(discl) - 2
		if gap < 1 {
			gap = 1
		}
		titleLine := len(r.lines) // index of the title row
		emit(title + strings.Repeat(" ", gap) + discl)
		nl() // blank line after title

		// ── Overall summary bar ───────────────────────────────────────────
		var pctStyle lipgloss.Style
		switch {
		case avgPct >= 70:
			pctStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cGood))
		case avgPct >= 30:
			pctStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cWarn))
		default:
			pctStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(cBad))
		}
		overallPctStr := pctStyle.Render(fmt.Sprintf("%5.1f%%", avgPct))
		overallBar := renderOverallBar(avgPct)
		overallLabel := sMuted.Render("  avg across all accounts")
		
		emit("     " + overallBar + "  " + overallPctStr + overallLabel)
		nl() // space after overall bar

		// Record click zones for exactly the title row
		r.clickZones = append(r.clickZones, clickZone{
			category:  cat.category,
			titleLine: titleLine,
			models:    labels,
		})

		// Scale ruler
		scaleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
		emit(scaleStyle.Render("      0%   20%  40%  60%  80% 100%"))
		nl() // space after scale

		// ── Column header ─────────────────────────────────────────────────
		emit(sColHdr.Render("     ACCOUNT              QUOTA                           PCT    RESET"))
		nl() // space after column header

		// ── Per-account rows ──────────────────────────────────────────────
		for _, g := range cat.groups {
			for _, acc := range g.Accounts {
				bar, pctStr := renderBar(acc.Percentage)
				reset := sReset.Render(padRight(acc.ResetIn, 6))
				email := sEmail.Render(padRight(api.ShortEmail(acc.Email), 20))
				emit(fmt.Sprintf("     %s  %s  %s  %s", email, bar, pctStr, reset))
				nl() // blank line between every account row
			}
		}

		// Category separator (dashed line between categories)
		if ci < len(cats)-1 {
			emit("     " + sDivLo.Render(strings.Repeat("╌", min(w-7, 52))))
			nl()
			nl()
		}
	}

	return r
}

// ═══════════════════════════════════════════════════════════════════════════════
// RIGHT PANEL — Local Cache
// ═══════════════════════════════════════════════════════════════════════════════

func buildRightPanel(accounts []config.Account, w int) []string {
	var lines []string
	nl := func() { lines = append(lines, "") }
	emit := func(s string) { lines = append(lines, s) }
	now := time.Now()
	nowMs := float64(now.UnixMilli())

	nl()
	emit(" " + sSection.Render(" Local Cache "))
	nl()

	// Gather all model keys
	allModels := map[string]struct{}{}
	for _, acc := range accounts {
		for k := range acc.RateLimitResetTimes {
			allModels[k] = struct{}{}
		}
	}
	if len(allModels) == 0 {
		emit("  " + sMuted.Render("No cache entries."))
		return lines
	}

	// Categorize
	buckets := map[string][]string{
		"Antigravity": {},
		"Gemini CLI":  {},
	}
	for k := range allModels {
		switch {
		case strings.HasPrefix(k, "gemini-antigravity:") || strings.Contains(k, "claude"):
			buckets["Antigravity"] = append(buckets["Antigravity"], k)
		case strings.HasPrefix(k, "gemini-cli:"):
			buckets["Gemini CLI"] = append(buckets["Gemini CLI"], k)
		default:
			buckets["Antigravity"] = append(buckets["Antigravity"], k)
		}
	}

	bucketOrder := []string{"Antigravity", "Gemini CLI"}
	for _, bkt := range bucketOrder {
		models := buckets[bkt]
		if len(models) == 0 {
			continue
		}
		sort.Strings(models)

		emit(" " + sCacheCat.Render(bkt))
		nl()

		for _, model := range models {
			cleanName := model
			if idx := strings.LastIndex(model, ":"); idx >= 0 {
				cleanName = model[idx+1:]
			}
			emit("   " + sCacheMod.Render(cleanName))
			nl()

			type row struct {
				email     string
				remaining float64
				ready     bool
				resetStr  string
				lastUsed  string
			}
			var rows []row

			for _, acc := range accounts {
				resetMs := acc.RateLimitResetTimes[model]
				remaining := resetMs - nowMs
				available := resetMs == 0 || remaining <= 0

				var rs, lu string
				switch {
				case resetMs == 0:
					rs = "—"
					lu = "never"
				case available:
					rs = "ready"
					d := time.Duration(math.Abs(remaining)) * time.Millisecond
					lu = api.FormatDuration(d) + " ago"
				default:
					d := time.Duration(remaining) * time.Millisecond
					rs = api.FormatDuration(d)
					lu = "—"
				}

				rows = append(rows, row{
					email:     api.ShortEmail(acc.Email),
					remaining: remaining,
					ready:     available,
					resetStr:  rs,
					lastUsed:  lu,
				})
			}

			sort.Slice(rows, func(i, j int) bool {
				return rows[i].remaining < rows[j].remaining
			})

			for _, r := range rows {
				var badge string
				if r.ready {
					badge = sBadgeOk.Render("OK")
				} else {
					badge = sBadgeWt.Render("WT")
				}
				info := sMuted.Render(fmt.Sprintf("%-7s %-9s", r.resetStr, r.lastUsed))
				name := sEmail.Render(r.email)
				emit(fmt.Sprintf("    %s %s %s", badge, info, name))
				nl()
			}
			nl()
		}
	}

	return lines
}

// ═══════════════════════════════════════════════════════════════════════════════
// CLICK DETECTION
// ═══════════════════════════════════════════════════════════════════════════════

func (m Model) findClickZone(x, y int) (clickZone, bool) {
	// The sticky header occupies rows 0 and 1.
	const hdrH = 2
	ftrH := m.footerLineCount()
	bodyH := m.height - hdrH - ftrH
	if bodyH < 1 || y < hdrH || y >= hdrH+bodyH {
		return clickZone{}, false
	}

	// We add +1 because the user's terminal rendering has an off-by-one
	// shift. This shifts the detection zone exactly one line UP on the screen
	// to perfectly align with the text they are clicking.
	bodyLine := (y - hdrH) + m.scrollOffset + 1

	for _, cz := range m.clickZones {
		// Only exact match on the title text line
		if bodyLine == cz.titleLine {
			return cz, true
		}
	}
	return clickZone{}, false
}

func (m Model) leftW() int {
	w := m.width
	if w < 100 {
		return w * 60 / 100
	}
	return w * 55 / 100
}

// ═══════════════════════════════════════════════════════════════════════════════
// CATEGORY CLASSIFICATION
// ═══════════════════════════════════════════════════════════════════════════════

func classifyGroup(g api.ModelGroup) string {
	for _, l := range g.Labels {
		low := strings.ToLower(l)
		if strings.Contains(low, "claude") || strings.Contains(low, "gpt-oss") || strings.Contains(low, "gpt oss") {
			return catClaude
		}
	}
	for _, l := range g.Labels {
		low := strings.ToLower(l)
		if strings.Contains(low, "tab_flash") || strings.Contains(low, "tab_jump") {
			return catFlashLite
		}
	}
	return catGemini
}

func categorizeGroups(groups []api.ModelGroup) []categorizedGroup {
	m := map[string][]api.ModelGroup{}
	for _, g := range groups {
		cat := classifyGroup(g)
