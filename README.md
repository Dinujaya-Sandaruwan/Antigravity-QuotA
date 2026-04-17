# Antigravity QuotA

A sleek terminal dashboard to monitor your AI model quotas and rate limits across multiple accounts.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Features

- Real-time quota monitoring for Gemini, Claude, and Flash Lite models
- Multi-account support with per-account breakdowns
- Visual progress bars showing remaining API quota
- Local rate-limit cache tracking
- Interactive modal popups for model discovery
- Mouse and keyboard navigation
- Cross-platform (Windows, macOS, Linux)

## Installation

### Download Pre-built Binaries

Head to the [Releases](https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA/releases) page and download the appropriate file for your operating system:

| OS      | File                                      |
| ------- | ----------------------------------------- |
| Windows | `antigravity-quota_windows_amd64.zip`     |
| macOS   | `antigravity-quota_darwin_arm64.tar.gz`   |
| Linux   | `antigravity-quota_linux_amd64.tar.gz`    |

Extract the archive and run the binary directly from your terminal.

### Build from Source

Requires [Go 1.21+](https://go.dev/dl/).

```bash
git clone https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA.git
cd Antigravity-QuotA
go build -o antigravity-quota .
./antigravity-quota
```

## Configuration

The application reads account credentials from a JSON file called `antigravity-accounts.json`. It searches the following locations in order:

| OS      | Path                                                              |
| ------- | ----------------------------------------------------------------- |
| macOS   | `~/.config/opencode/antigravity-accounts.json`                    |
| Linux   | `~/.config/opencode/antigravity-accounts.json`                    |
| Windows | `C:\Users\<You>\.config\opencode\antigravity-accounts.json`  |

You can override the config location by setting the `OPENCODE_CONFIG_DIR` environment variable:

```bash
# macOS / Linux
OPENCODE_CONFIG_DIR="/path/to/config" ./antigravity-quota

# Windows (PowerShell)
$env:OPENCODE_CONFIG_DIR="C:\path\to\config"
.\antigravity-quota.exe
```

### Sample Configuration

```json
{
  "accounts": [
    {
      "email": "your-email@gmail.com",
      "refreshToken": "your-refresh-token",
      "rateLimitResetTimes": {}
    }
  ],
  "activeIndex": 0
}
```

## Usage

### Keyboard Shortcuts

| Key        | Action              |
| ---------- | ------------------- |
| `R`        | Refresh quota data  |
| `Up` / `K` | Scroll up           |
| `Down` / `J` | Scroll down       |
| `PgUp`     | Page up             |
| `PgDn`     | Page down           |
| `Q`        | Quit                |
| `Ctrl+C`   | Force quit          |

### Mouse

- **Scroll wheel** to navigate
- **Click** on a category header to view the full list of models in that category

## How It Works

1. Reads account credentials from your local configuration file
2. Refreshes OAuth2 tokens for each account
3. Queries the Google Cloud Code API for available models and their quota limits
4. Displays the results in a split-panel terminal dashboard:
   - **Left panel**: Live API quota with visual progress bars
   - **Right panel**: Local rate-limit cache status

## Built With

- [Go](https://go.dev/) - Programming language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
