#!/bin/bash

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "🔄 Generating documentation..."
make docs

echo "🔍 Checking for documentation changes..."

# Get status of all changes (modified, added, deleted, untracked)
changes=$(git status --porcelain)

if [[ -n $changes ]]; then
    echo -e "${RED}❌ Documentation changes detected${NC}"
    echo "::error title=Documentation out of date::Please run 'make docs' and commit the changes"

    # Categorize changes for better error reporting
    modified=$(echo "$changes" | grep '^ M' | sed 's/^ M //' || true)
    added=$(echo "$changes" | grep '^??' | sed 's/^?? //' || true)
    deleted=$(echo "$changes" | grep '^ D' | sed 's/^ D //' || true)

    if [[ -n $modified ]]; then
        echo "Modified files:"
        echo "$modified"
    fi

    if [[ -n $added ]]; then
        echo "New files:"
        echo "$added"
    fi

    if [[ -n $deleted ]]; then
        echo "Deleted files:"
        echo "$deleted"
    fi

    echo ""
    echo "Full changes:"
    echo "$changes"
    exit 1
fi

echo -e "${GREEN}✅ Documentation is up-to-date${NC}"
exit 0
