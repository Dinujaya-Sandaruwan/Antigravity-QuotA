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
