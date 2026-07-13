<div align="center">
  
# Antigravity QuotA

**A sleek terminal dashboard to monitor your AI model quotas and rate limits across multiple accounts.**

[![npm version](https://img.shields.io/npm/v/antigravity-quota-tui.svg?style=flat-square)](https://www.npmjs.com/package/antigravity-quota-tui)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)](#)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-blue?style=flat-square)](#)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](#)

<img src="./assets/dashboard.png" alt="Antigravity Quota Dashboard" width="800"/>

</div>

## ✨ Features

- **Real-Time Monitoring**: Instantly check API quotas for **Gemini**, **Claude**, and **Flash Lite** models.
- **Multi-Account Support**: Aggregate and view quotas for dozens of accounts side-by-side.
- **Visual Progress Bars**: Quickly identify depleted or available accounts with color-coded (Green/Yellow/Red) status bars.
- **Local Cache Tracking**: Monitor local rate-limits and cooldowns for specific extensions and CLI tools.
- **Cross-Platform**: Works flawlessly on Windows, macOS, and Linux terminal emulators.
- **Keyboard & Mouse**: Fully interactive TUI with scroll and click support.

---

## 🚀 Installation

### Option 1: Using NPM (Recommended)
If you have Node.js installed, you can install the CLI globally. This will automatically download the correct highly-optimized Go binary for your operating system.

```bash
npm install -g antigravity-quota-tui
```

### Option 2: Pre-built Binaries
Don't have Node.js? Head to the [Releases](https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA/releases) page and download the standalone executable for your OS (Windows `.zip`, macOS/Linux `.tar.gz`). Extract it and run it directly in your terminal.

### Option 3: Build from Source
```bash
git clone https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA.git
cd Antigravity-QuotA
go build -o antigravity-quota .
./antigravity-quota
```

---

## ⚡️ Usage

Simply run the following command in your terminal:

```bash
ant-quota
```
*(If you downloaded the standalone binary, run `./antigravity-quota` or `antigravity-quota.exe`)*

### Keyboard & Mouse Controls
| Input | Action |
| --- | --- |
| `R` | Refresh quota data from Google APIs |
| `Up` / `K` | Scroll up |
| `Down` / `J` | Scroll down |
| `PgUp` / `PgDn` | Scroll by page |
| `Mouse Click` | Click a category title to view all associated models |
| `Q` or `Ctrl+C` | Quit the application |

---

## ⚙️ Configuration (Important!)

The application needs your account credentials to fetch quota limits. It reads these from a JSON file named `antigravity-accounts.json`.

> 💡 **CRITICAL: The `OPENCODE_CONFIG_DIR` Environment Variable**
>
> By default, the application looks for your configuration file in standard OS locations (e.g., `~/.config/opencode/` on Mac/Linux or `C:\Users\<You>\.config\opencode\` on Windows).
> 
> **The best way to configure this tool** is to place your `antigravity-accounts.json` wherever you want, and set the `OPENCODE_CONFIG_DIR` environment variable to point to that folder.

#### Setting the Environment Variable:

**macOS / Linux:**
```bash
# Set it temporarily for one run:
OPENCODE_CONFIG_DIR="/path/to/my/custom/config/folder" ant-quota

# Or add it permanently to your ~/.bashrc or ~/.zshrc:
export OPENCODE_CONFIG_DIR="/path/to/my/custom/config/folder"
```

**Windows (PowerShell):**
```powershell
# Set it for your session:
$env:OPENCODE_CONFIG_DIR="C:\Path\To\My\Config\Folder"
ant-quota
```

### Sample `antigravity-accounts.json` Format
Make sure your configuration file looks like this:

```json
{
  "accounts": [
    {
      "email": "info.samantha@example.com",
      "refreshToken": "1//0e...",
      "rateLimitResetTimes": {}
    },
    {
      "email": "study.dinujaya@example.com",
      "refreshToken": "1//0c...",
      "rateLimitResetTimes": {
        "claude": 1730000000000
      }
    }
  ],
  "activeIndex": 0
}
```

---

## 🛠 Built With

* **[Go](https://go.dev/)** - The core programming language
* **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - The terminal UI framework
* **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for terminal layouts

## 📝 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
