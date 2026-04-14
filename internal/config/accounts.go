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
