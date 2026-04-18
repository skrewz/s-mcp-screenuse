#!/bin/bash
set -e

# Change to script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

# Build if needed (idempotent via Makefile)
make build

# Execute the MCP server
./s-mcp-screenuse
