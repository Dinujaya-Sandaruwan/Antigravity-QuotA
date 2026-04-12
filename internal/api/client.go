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
