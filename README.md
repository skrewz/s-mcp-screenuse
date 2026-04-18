# Wayland Screenshot MCP Server

A Model Context Protocol (MCP) server that provides screenshot capture capabilities on Wayland systems using `grim`.

## Features

- Single tool: `take_screenshot` - captures the entire screen as a PNG image
- Returns base64-encoded image data via MCP's `ImageContent` type
- Initialization validation: verifies grim access during startup
- 5-second timeout for all grim operations

## Prerequisites

- Linux with Wayland session
- `grim` installed (`sudo pacman -S grim` on Arch, `sudo apt install grim` on Debian/Ubuntu)

## Usage

### Build

```bash
make build
```

### Run

```bash
make run
```

### Build and run in one step

```bash
make run
```

## Configuration

No configuration required. The server uses default grim behavior (captures primary display).

## Opencode MCP Configuration

Add this to your `~/.config/opencode/opencode.jsonc` under the `mcp` section:

```json
"wayland-screenshot": {
  "type": "local",
  "command": ["{env:HOME}/repos/s-mcp-screenuse/scripts/s-mcp-screenuse.sh"],
  "enabled": true
}
```

## License

MIT
