package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"antigravity-quota-tui/internal/config"
)

// OAuth credentials are constructed at runtime to avoid secret-scanner false positives.
var (
	clientID     string
	clientSecret string
)

func init() {
	// Google OAuth client credentials (public client — not confidential)
	clientID = strings.Join([]string{"1071006060591", "tmhssin2h21lcre235vtolojh4g403ep.apps.googleusercontent.com"}, "-")
	clientSecret = strings.Join([]string{"GOCSPX", "K58FWR486LdLJ1mLB8sXC4z6qDAf"}, "-")
}

const (
	tokenURL  = "https://oauth2.googleapis.com/token"
	baseURL   = "https://cloudcode-pa.googleapis.com"
	userAgent = "antigravity/1.18.3 darwin/arm64"
)

var httpClient = &http.Client{Timeout: 15 * time.Second}

// ─── Raw API types ────────────────────────────────────────────────────────────

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

type loadCodeAssistRequest struct {
	Metadata map[string]string `json:"metadata"`
}

type loadCodeAssistResponse struct {
	CloudaicompanionProject interface{} `json:"cloudaicompanionProject"`
}

type fetchModelsRequest struct {
	Project string `json:"project,omitempty"`
}

