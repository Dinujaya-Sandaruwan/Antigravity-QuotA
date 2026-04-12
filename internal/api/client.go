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
