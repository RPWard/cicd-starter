#!/bin/bash

# Print current directory for debugging
echo "Building from directory: $(pwd)"

# Ensure we're building with static file embedding
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely

# Verify the binary was created
if [ -f notely ]; then
    echo "Binary 'notely' built successfully"
    # Optional: you can verify the binary contains embedded files
    ls -lh notely
else
    echo "Failed to build binary"
    exit 1
fi
