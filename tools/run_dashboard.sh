#!/usr/bin/env bash
# tools/run_dashboard.sh - Unified launcher for Bobtk Operations Dashboard and Server Supervisor
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "================================================================================"
echo "          BOBTK DASHBOARD RUNNER & PROCESS MONITOR"
echo "================================================================================"
echo "Repository Root: ${REPO_ROOT}"
echo "Current Time: $(date)"

cd "${REPO_ROOT}/go"

echo "Running Go Unit Tests..."
go test -v ./...

echo "Executing Unified Single-Page Dashboard & Server Supervisor..."
go run main.go

echo "Dashboard runner execution concluded cleanly."
