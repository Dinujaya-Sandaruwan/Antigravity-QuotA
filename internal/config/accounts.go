package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Account mirrors the JSON structure in antigravity-accounts.json
type Account struct {
	Email              string             `json:"email"`
	RefreshToken       string             `json:"refreshToken"`
	ProjectID          string             `json:"projectId,omitempty"`
	ManagedProjectID   string             `json:"managedProjectId,omitempty"`
	RateLimitResetTimes map[string]float64 `json:"rateLimitResetTimes"`
}

// AccountsConfig is the top-level structure of antigravity-accounts.json
type AccountsConfig struct {
	Accounts    []Account `json:"accounts"`
	ActiveIndex int       `json:"activeIndex"`
}

// configPaths returns candidate paths for the accounts file, in priority order.
func configPaths() []string {
	home, _ := os.UserHomeDir()

