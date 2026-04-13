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
