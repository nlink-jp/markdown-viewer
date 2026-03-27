# markdown-viewer

A secure, single-binary Markdown viewer and file browser for local directories. Renders GFM, Mermaid diagrams, and syntax-highlighted code blocks.

## Features

- **2-pane UI**: Persistent, expandable file tree on the left; content viewer on the right
- **Single binary**: All UI assets (CSS, JS, HTML templates) are embedded — no external dependencies
- **Markdown rendering**: GitHub-like styling, syntax highlighting, and bundled Mermaid diagram support
- **Secure by default**: Directory traversal protection and HTML sanitization against XSS
- **Graceful shutdown**: Via UI button or `Ctrl+C`

## Installation

Download the latest `mdv` binary for your platform from the [releases page](https://github.com/nlink-jp/markdown-viewer/releases).

Extract and place the binary in your `$PATH`:

```sh
unzip mdv-<version>-<os>-<arch>.zip
mv mdv /usr/local/bin/
```

## Configuration

`mdv` can be configured via a `config.json` file or command-line flags. Precedence order (later overrides earlier):

1. Default values
2. `$HOME/.config/mdv/config.json`
3. `config.json` in the current working directory
4. Environment variables (prefixed `MDV_`, e.g. `MDV_PORT=9000`)
5. Command-line flags

See `config.json.example` for a reference configuration.

## Usage

Run `mdv` from any directory containing Markdown files:

```sh
mdv [flags]
```

Then open `http://127.0.0.1:8080` in your browser.

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-p, --port` | `8080` | Port to listen on |
| `-o, --open` | `false` | Open browser automatically on startup |
| `-d, --dir` | `.` | Root directory to serve |
| `--version` | | Print version and exit |

## Building

Requires Go 1.24 or later.

```sh
git clone https://github.com/nlink-jp/markdown-viewer.git
cd markdown-viewer
make build        # Build for the current platform → dist/mdv
make build-all    # Cross-compile for all platforms → dist/mdv-<os>-<arch>
make package      # Build and create .zip archives → dist/mdv-<version>-<os>-<arch>.zip
make test         # Run the test suite
make clean        # Remove dist/
```

Target platforms: `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`.

## See Also

- [CHANGELOG.md](CHANGELOG.md)
- [NOTICE.md](NOTICE.md) — third-party license attributions
