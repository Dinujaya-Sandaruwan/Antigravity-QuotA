#!/usr/bin/env node
/**
 * install.js - Postinstall script
 * Downloads the correct pre-built binary from GitHub Releases based on
 * the user's OS and architecture, and places it inside bin/.
 */
const fs = require("fs");
const path = require("path");
const https = require("https");
const os = require("os");
const { execSync } = require("child_process");

const REPO = "Dinujaya-Sandaruwan/Antigravity-QuotA";
const pkg = require("./package.json");
const VERSION = pkg.version;

// Map Node platform/arch to GoReleaser artifact names
function detectPlatform() {
  const platform = process.platform; // darwin, linux, win32
  const arch = process.arch;         // x64, arm64

  const platformMap = { darwin: "darwin", linux: "linux", win32: "windows" };
  const archMap = { x64: "amd64", arm64: "arm64" };

  const goOs = platformMap[platform];
  const goArch = archMap[arch];

  if (!goOs || !goArch) {
    console.error(
      `\nUnsupported platform: ${platform}/${arch}\n` +
      `Supported: darwin/x64, darwin/arm64, linux/x64, linux/arm64, windows/x64, windows/arm64\n`
    );
    process.exit(1);
  }

  const ext = goOs === "windows" ? "zip" : "tar.gz";
  const filename = `antigravity-quota_${goOs}_${goArch}.${ext}`;
  const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${filename}`;
  return { goOs, goArch, ext, filename, url };
}

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);

    const doGet = (link) => {
      https
        .get(link, (res) => {
          if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
            doGet(res.headers.location);
            return;
          }
          if (res.statusCode !== 200) {
            reject(new Error(`Download failed: HTTP ${res.statusCode} for ${link}`));
            return;
          }
          res.pipe(file);
          file.on("finish", () => file.close(resolve));
        })
        .on("error", reject);
    };

    doGet(url);
  });
}

async function main() {
  const { goOs, ext, filename, url } = detectPlatform();
  const binDir = path.join(__dirname, "bin");
  fs.mkdirSync(binDir, { recursive: true });

  const archivePath = path.join(binDir, filename);

  console.log(`\nDownloading antigravity-quota v${VERSION} for ${goOs}...`);
  console.log(`   ${url}\n`);

  try {
    await download(url, archivePath);
  } catch (err) {
    console.error("Failed to download binary:", err.message);
    console.error("You can download it manually from:");
    console.error(`   https://github.com/${REPO}/releases/tag/v${VERSION}`);
    process.exit(1);
  }

  // Extract archive
  try {
    if (ext === "zip") {
      // Try PowerShell on Windows, unzip elsewhere
      if (process.platform === "win32") {
        execSync(
          `powershell -Command "Expand-Archive -Force -Path '${archivePath}' -DestinationPath '${binDir}'"`,
          { stdio: "inherit" }
        );
      } else {
        execSync(`unzip -o "${archivePath}" -d "${binDir}"`, { stdio: "inherit" });
      }
    } else {
      execSync(`tar -xzf "${archivePath}" -C "${binDir}"`, { stdio: "inherit" });
    }
  } catch (err) {
    console.error("Failed to extract archive:", err.message);
    process.exit(1);
  }

  // Clean up archive
  fs.unlinkSync(archivePath);

  // Make binary executable on Unix
  const binaryName = process.platform === "win32" ? "antigravity-quota.exe" : "antigravity-quota";
  const binaryPath = path.join(binDir, binaryName);
  if (fs.existsSync(binaryPath) && process.platform !== "win32") {
    fs.chmodSync(binaryPath, 0o755);
  }

  console.log(`\nInstalled antigravity-quota v${VERSION}`);
  console.log(`   Run it with: ant-quota\n`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
