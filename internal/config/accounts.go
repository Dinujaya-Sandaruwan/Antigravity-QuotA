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

	// Respect OPENCODE_CONFIG_DIR override (same logic as the TS plugin)
	if override := os.Getenv("OPENCODE_CONFIG_DIR"); override != "" {
		return []string{filepath.Join(override, "antigravity-accounts.json")}
	}

	xdgData := os.Getenv("XDG_DATA_HOME")
	if xdgData == "" {
		xdgData = filepath.Join(home, ".local", "share")
	}

	seen := map[string]struct{}{}
	var paths []string
	for _, p := range []string{
		filepath.Join(home, ".config", "opencode", "antigravity-accounts.json"),
		filepath.Join(xdgData, "opencode", "antigravity-accounts.json"),
	} {
		if _, ok := seen[p]; !ok {
			seen[p] = struct{}{}
			paths = append(paths, p)
		}
	}
	return paths
}

// Load reads and parses the accounts config file from the first path that exists.
func Load() (*AccountsConfig, error) {
	paths := configPaths()
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, fmt.Errorf("reading %s: %w", p, err)
		}
		var cfg AccountsConfig
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", p, err)
		}
		// Assign fallback email labels (same as TS plugin)
		for i := range cfg.Accounts {
			if cfg.Accounts[i].Email == "" {
				cfg.Accounts[i].Email = fmt.Sprintf("account-%d", i+1)
			}
		}
		return &cfg, nil
	}
	return nil, fmt.Errorf("configuration file not found; checked:\n%s\n\nEnsure opencode-antigravity-auth is installed and configured", formatPaths(paths))
}

func formatPaths(paths []string) string {
	out := ""
	for _, p := range paths {
		out += "  - " + p + "\n"
	}
	return out
}
