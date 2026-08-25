# AGENTS.md — markdown-viewer

## Project summary

Single-binary Markdown viewer with 2-pane UI (file tree + content).
Renders GFM, Mermaid, syntax highlighting. All assets embedded.
Part of util-series.

## Build commands

```bash
make build          # Build → dist/mdv
make test           # Run all tests
make build-all      # Cross-compile for 4 platforms (darwin arm64 only; no Intel)
make package        # Build + create .zip archives
make verify-release  # gate: .notarized marker + freshness (run before upload)
make clean          # Remove dist/
```

## Module path

`github.com/nlink-jp/markdown-viewer`

## Key structure

```
markdown-viewer/
├── main.go                        ← entry point
├── cmd/root.go                    ← cobra CLI, server lifecycle
├── internal/
│   ├── assets/                    ← embedded HTML/CSS/JS (go:embed)
│   │   └── embed_assets/
│   │       ├── static/            ← main.css, main.js, treeview.css, treeview.js
│   │       └── templates/         ← index.html, treeview.html, markdown.html, etc.
│   ├── browser/                   ← open URL in default browser
│   ├── config/                    ← viper-based config (JSON + env + flags)
│   ├── filebrowser/               ← directory listing (markdown files only)
│   ├── markdown/                  ← goldmark SafeLinkRenderer
│   └── server/                    ← HTTP server, handlers, routing
├── Makefile
└── config.json.example
```

## Configuration

Priority: defaults → `~/.config/mdv/config.json` → `./config.json` → env (`MDV_*`) → CLI flags.

- `-p, --port` (default: 8080)
- `-o, --open` (default: false)
- `-d, --dir` (default: .)

## Security

- Directory traversal: blocked by `containsDotDot()` on raw URI
- XSS: HTML sanitized via bluemonday UGCPolicy
- Links: only relative `.md`/`.markdown` files rendered as `<a>` tags
- Protocol schemes (`http://`, `ftp://`, `javascript:`, `data:`) blocked in links
- Server binds to 127.0.0.1 only (no network exposure)

## Gotchas

- All UI assets are embedded via `go:embed` — changes to HTML/CSS/JS require rebuild
- `isSafeLink` blocks all protocol schemes, not just http/https
- Config uses viper: flag binding happens in PersistentPreRunE, not init()
- Server implements `http.Handler` directly (no mux)
