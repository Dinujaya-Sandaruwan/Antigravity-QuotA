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
