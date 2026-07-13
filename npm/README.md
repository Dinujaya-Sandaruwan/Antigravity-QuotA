<div align="center">
  
# Antigravity QuotA (TUI)

**A sleek terminal dashboard to monitor your AI model quotas and rate limits across multiple accounts.**

[![npm version](https://img.shields.io/npm/v/antigravity-quota-tui.svg?style=flat-square)](https://www.npmjs.com/package/antigravity-quota-tui)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](#)

<img src="https://raw.githubusercontent.com/Dinujaya-Sandaruwan/Antigravity-QuotA/main/assets/dashboard.png" alt="Antigravity Quota Dashboard" width="800"/>

</div>

---

## ✨ Overview

`antigravity-quota-tui` is a standalone CLI tool that provides a beautiful, real-time dashboard for monitoring your Google Cloud Code / AI Companion API quotas. 

Whether you are managing multiple accounts for **Gemini**, **Claude**, or **Flash Lite** models, this tool aggregates all your remaining quotas, reset times, and local rate-limit caches into a single, interactive terminal interface.

*Note: This npm package is a lightweight JavaScript wrapper that automatically downloads the highly-optimized native Go binary for your specific operating system (Windows, macOS, or Linux).*

---

## 🚀 Installation

Install the package globally using npm to make the `ant-quota` command available everywhere on your system:

```bash
npm install -g antigravity-quota-tui
```

---

## ⚡️ Usage

Launch the dashboard from any terminal by typing:

```bash
ant-quota
```

### Controls
- **`R`**: Refresh quota data from the API
- **`Up` / `Down`** (or `K` / `J`): Scroll the view
- **`Click`**: Click on a category title (e.g., *Claude Models*) to view the full list of grouped models
- **`Q`** or **`Ctrl+C`**: Quit

---

## ⚙️ Configuration (Important!)

The tool requires your account credentials to fetch the data. It expects to find a configuration file named `antigravity-accounts.json`.

> 🔥 **CRITICAL: The `OPENCODE_CONFIG_DIR` Variable**
>
> To ensure the tool can reliably find your configuration file across different operating systems, **it is highly recommended to set the `OPENCODE_CONFIG_DIR` environment variable.**

Place your `antigravity-accounts.json` file in a secure folder, and point the tool to that folder:

**macOS / Linux:**
```bash
# Add this to your ~/.bashrc or ~/.zshrc profile:
export OPENCODE_CONFIG_DIR="/path/to/my/config/folder"
```

**Windows (PowerShell):**
```powershell
# Set it in your environment variables:
$env:OPENCODE_CONFIG_DIR="C:\Path\To\My\Config\Folder"
```

*(Fallback: If not set, it attempts to read from `~/.config/opencode/` on macOS/Linux, or `C:\Users\<User>\.config\opencode\` on Windows).*

### Sample `antigravity-accounts.json`

```json
{
  "accounts": [
    {
      "email": "your-email@example.com",
      "refreshToken": "1//0eabc...",
      "rateLimitResetTimes": {}
    }
  ],
  "activeIndex": 0
}
```

---

## 🔗 Links

- **GitHub Repository**: [Dinujaya-Sandaruwan/Antigravity-QuotA](https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA)
- **Report an Issue**: [GitHub Issues](https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA/issues)

## 📝 License
MIT License. See the [GitHub repository](https://github.com/Dinujaya-Sandaruwan/Antigravity-QuotA) for full details.
