# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.0] - 2026-07-12

### Removed

- **darwin/amd64 (Intel) pre-built binary.** macOS releases now ship
  **arm64 only**, per the org-wide policy (darwin is Apple-Silicon only; no
  universal binaries). Intel Mac users can build from source.

### Changed

- **Linux release archives are now `.tar.gz`** (darwin/windows remain `.zip`),
  per `nlink-jp/.github` CONVENTIONS.md §Release Archive Standard. Archives
  still bundle `LICENSE` + `README.md` alongside the canonical `mdv` binary.
- **darwin code-signature identifier** is now the canonical `mdv`
  (was `mdv-darwin-arm64`), set via `codesign -i` so it stays stable after
  the archived binary is renamed to its canonical name.
- **Dropped the `-s -w` linker strip flags**, aligning `LDFLAGS` with the
  org-standard form (`-X main.version=…`); also avoids a false-positive
  antivirus quarantine of the stripped Windows binary during cross-build.

No change to the binary's behaviour — a packaging / build-config release.

## [1.3.1] - 2026-05-23

### Changed

- **Darwin releases are now Developer ID signed and Apple-notarized.**
  `mdv-v1.3.1-darwin-{amd64,arm64}.zip` carry full Apple Developer
  ID Application signatures and notarization tickets from Apple.
  End users on macOS no longer need to bypass Gatekeeper with
  right-click → Open or `xattr -d com.apple.quarantine` on first
  launch; local users who place `mdv` under Dropbox-synced (or any
  other FileProvider-managed) paths are no longer killed by
  macOS's ad-hoc + provenance distrust policy. Pipeline:
  `scripts/codesign-darwin.sh` + `scripts/notarize-darwin.sh`,
  driven by `make package`. Adopts the org-wide convention in
  `nlink-jp/.github` CONVENTIONS.md §Code Signing.
- **Release zips now include LICENSE and README.md** alongside
  the binary, matching the sibling util-series tools. Previous
  v1.3.0 assets shipped the binary only.
- **Release zip filenames now effectively embed the version**
  (`mdv-vX.Y.Z-<os>-<arch>.zip`); the Makefile produced versioned
  names already but v1.3.0 was published under version-less
  filenames. v1.3.1 is the first release with the Makefile-driven
  versioned names matching the artifact list.

No behaviour change to the binary itself — feature-wise this is
identical to v1.3.0.

## [1.3.0] - 2026-04-15

### Added

- Draggable sidebar divider for resizing file tree pane (150px–60%)
- File names truncated with ellipsis when too long for sidebar width

### Fixed

- Sidebar file list no longer hidden behind shutdown button when list is long
- `isSafeLink` now blocks all protocol schemes (ftp://, javascript:, data:, etc.), not just http/https
- Content-Type header set before response body in ApiListHandler
- JavaScript null check uses `Array.isArray()` instead of `=== null`

### Added

- Unit tests for config, filebrowser, markdown renderer, browser, and server
- README.ja.md (Japanese documentation)
- AGENTS.md

### Changed

- Default port standardized to 8080 (was 8888 in config.go, 8080 in CLI)
- Sidebar footer uses flex layout instead of position:absolute
- Removed dead code (internal/server/templates.go)
- Removed stale TODO comment (browser auto-open already implemented)
- Removed invalid JSON comments from config.json.example
- Replaced `javascript:void(0)` with `#` + preventDefault in treeview.js

## [1.2.0] - 2026-03-28

### Changed
- Migrated to nlink-jp organisation — module path updated to `github.com/nlink-jp/markdown-viewer`
- Standardised Makefile: `dist/` output, `build` / `build-all` / `package` / `test` / `clean` targets; separate amd64/arm64 binaries (no universal binary)
- Updated README to follow organisation conventions

## [1.1.0] - 2025-08-10

### Added
- **Embedded UI Assets:** All CSS, JavaScript, and HTML templates are now embedded directly into the executable, making the application a truly single, self-contained binary. This eliminates the need for a separate `static/` or `templates/` directory alongside the binary at runtime.

### Changed
- **Refactored Asset Management:** The internal structure for managing UI assets has been refactored for better organization and maintainability.
- **Updated Documentation:** `DEVELOPMENT.md` and `README.md` have been updated to reflect the single-binary nature and embedded assets.


## [1.0.1] - 2025-08-10

### Security
- **Resolved GO-2025-3595 Vulnerability:** Updated `golang.org/x/net` from v0.26.0 to v0.38.0 to address a vulnerability related to incorrect neutralization of input during web page generation.

### Added
- **Local Bundling of Client-Side Assets:** `highlight.js` and `Mermaid.js` are now bundled directly with the application, removing external CDN dependencies for offline and closed-network environments.
- **Third-Party License Notice:** Included `NOTICE.md` in release packages for compliance with third-party software licenses.

### Changed
- **Updated Documentation:** `DEVELOPMENT.md` and `README.md` have been updated to reflect local asset bundling and license information.


## [1.0.0] - 2025-08-10

This is the first stable release after a major security and functionality overhaul.

### Security
- **Prevented Directory Traversal:** Implemented a custom router to validate request paths, blocking access to files outside the intended directory.
- **Prevented Cross-Site Scripting (XSS):** Integrated the `bluemonday` library to sanitize all HTML rendered from Markdown, mitigating XSS risks.
- **Hardened Browser Auto-Open:** Re-implemented the browser auto-open feature with strict URL validation and command path verification to prevent command injection vulnerabilities.
- **Disabled Unsafe Links:** The Markdown renderer now disables links to local non-Markdown files and requires user confirmation before opening external links in a new tab.
- **Updated Dependencies:** Updated all dependencies to their latest versions to patch known vulnerabilities.
- **Addressed `gosec` Findings:** Fixed all issues reported by the `gosec` static analysis tool, including potential Slowloris attacks and unhandled errors.

### Fixed
- **Markdown Rendering Engine:** Replaced the `blackfriday` parser with `goldmark` and its GFM extension. This fixed numerous rendering bugs, including incorrect handling of code blocks within lists, unwanted line breaks, and improved standards compliance.
- **Shutdown Mechanism:** Fixed a bug where the shutdown button was not working due to an uninitialized channel.

### Changed
- **Improved Layout:** Increased the maximum width of the Markdown rendering area to 1140px for better readability on wider screens.

[Unreleased]: https://github.com/nlink-jp/markdown-viewer/compare/v1.2.0...HEAD
[1.2.0]: https://github.com/nlink-jp/markdown-viewer/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/nlink-jp/markdown-viewer/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/nlink-jp/markdown-viewer/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/nlink-jp/markdown-viewer/releases/tag/v1.0.0
