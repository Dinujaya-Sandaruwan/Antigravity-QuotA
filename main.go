package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"antigravity-quota-tui/internal/config"
	"antigravity-quota-tui/internal/ui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
