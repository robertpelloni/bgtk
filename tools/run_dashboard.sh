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

# Process command-line flags
CHECK_HEALTH=0
for arg in "$@"; do
    case "${arg}" in
        --check-health|-c)
            CHECK_HEALTH=1
            ;;
    esac
done

if [ "${CHECK_HEALTH}" -eq 1 ]; then
    echo ""
    echo "Running Pre-flight Environment & Health Diagnostics..."
    bash "${SCRIPT_DIR}/check_health.sh"
    echo ""
fi

cd "${REPO_ROOT}/go"

echo "Running Go Unit Tests..."
go test -v ./...

echo "Executing Unified Single-Page Dashboard & Server Supervisor..."
go run main.go

echo "Dashboard runner execution concluded cleanly."
