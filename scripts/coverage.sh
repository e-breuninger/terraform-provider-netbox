#!/bin/bash

# Script to run tests with coverage analysis
# Usage: ./scripts/coverage.sh [html|func|upload]

set -e

echo "Running tests with coverage..."

# Run tests with coverage
make test

# Generate coverage reports
echo "Generating coverage reports..."
go tool cover -func=coverage.out > coverage_func.txt
go tool cover -html=coverage.out -o coverage.html

echo "Coverage summary:"
cat coverage_func.txt

# Check coverage threshold
COVERAGE=$(grep "total:" coverage_func.txt | awk '{print substr($3, 1, length($3)-1)}')
echo "Current coverage: $COVERAGE%"

if (( $(echo "$COVERAGE < 70.0" | bc -l 2>/dev/null || echo "1") )); then
    echo "⚠️  Coverage is below 70%: $COVERAGE%"
    exit 1
else
    echo "✅ Coverage is above 70%: $COVERAGE%"
fi

# Handle different output formats
case "$1" in
    "html")
        echo "Opening HTML coverage report..."
        if command -v xdg-open > /dev/null; then
            xdg-open coverage.html
        elif command -v open > /dev/null; then
            open coverage.html
        else
            echo "HTML report generated: coverage.html"
        fi
        ;;
    "func")
        echo "Function coverage report:"
        cat coverage_func.txt
        ;;
    "upload")
        echo "Uploading to Codecov..."
        if command -v codecov > /dev/null; then
            codecov -f coverage.out
        else
            echo "Codecov CLI not found. Install with: pip install codecov"
        fi
        ;;
    *)
        echo "Coverage reports generated:"
        echo "  - coverage.out (raw coverage data)"
        echo "  - coverage.html (HTML report)"
        echo "  - coverage_func.txt (function summary)"
        ;;
esac
