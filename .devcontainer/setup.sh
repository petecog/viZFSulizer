#!/bin/bash
set -e

# Initialize Go module if it doesn't exist
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module..."
    go mod init github.com/petecog/vizfsulizer
fi

# Install Go tools
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/go-delve/delve/cmd/dlv@latest

# Install project dependencies
go mod tidy

# Skip ZFS setup in container environment - not needed for development
echo "Skipping ZFS environment setup (not supported in container)"
echo "ZFS testing should be done on host system or dedicated test environment"

echo "Setup complete!"
