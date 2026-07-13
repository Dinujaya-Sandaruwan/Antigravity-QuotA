#!/usr/bin/env node
/**
 * ant-quota launcher.
 * Spawns the platform-specific antigravity-quota binary downloaded during postinstall.
 */
const path = require("path");
const { spawnSync } = require("child_process");
const fs = require("fs");

const binaryName = process.platform === "win32" ? "antigravity-quota.exe" : "antigravity-quota";
const binaryPath = path.join(__dirname, "..", "bin", binaryName);

if (!fs.existsSync(binaryPath)) {
  console.error("antigravity-quota binary not found.");
  console.error("Try reinstalling: npm i -g antigravity-quota-tui --force");
  process.exit(1);
}

const result = spawnSync(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
});

process.exit(result.status ?? 0);
