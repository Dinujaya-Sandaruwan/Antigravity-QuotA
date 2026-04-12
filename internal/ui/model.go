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
