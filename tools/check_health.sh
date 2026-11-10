#!/usr/bin/env bash
# tools/check_health.sh - Unified diagnostic tool verifying toolchains, submodules, and environment
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "================================================================================"
echo "          BGTK PROJECT HEALTH & ENVIRONMENT DIAGNOSTIC MONITOR"
echo "================================================================================"
echo "Repository Root: ${REPO_ROOT}"
echo "Timestamp:       $(date)"

echo ""
echo "[1. TOOLCHAIN & RUNTIME AVAILABILITY]"
if command -v go >/dev/null 2>&1; then
    echo "  ✔ Go Toolchain:       $(go version)"
else
    echo "  ✘ Go Toolchain:       NOT FOUND"
fi

if command -v gcc >/dev/null 2>&1; then
    echo "  ✔ C Compiler (gcc):   $(gcc --version | head -n 1)"
else
    echo "  ✘ C Compiler (gcc):   NOT FOUND"
fi

if command -v g++ >/dev/null 2>&1; then
    echo "  ✔ C++ Compiler (g++): $(g++ --version | head -n 1)"
else
    echo "  ✘ C++ Compiler (g++): NOT FOUND"
fi

if command -v meson >/dev/null 2>&1; then
    echo "  ✔ Meson Build System: $(meson --version)"
else
    echo "  ✘ Meson Build System: NOT FOUND"
fi

if command -v ninja >/dev/null 2>&1; then
    echo "  ✔ Ninja Generator:    $(ninja --version)"
else
    echo "  ✘ Ninja Generator:    NOT FOUND"
fi

echo ""
echo "[2. SUBMODULE SYNCHRONIZATION STATUS]"
for sub in "submodules/ultimatepp" "submodules/juce"; do
    if [ -d "${REPO_ROOT}/${sub}" ]; then
        SUB_COMMIT="$(git -C "${REPO_ROOT}/${sub}" rev-parse --short HEAD 2>/dev/null || echo 'uninitialized')"
        echo "  ✔ Submodule ${sub}: [Commit: ${SUB_COMMIT}]"
    else
        echo "  ℹ Submodule ${sub}: NOT INITIALIZED IN WORKSPACE (Configured in .gitmodules)"
    fi
done

echo ""
echo "[3. GO 6-PILLAR FRAMEWORK TEST SUITE]"
cd "${REPO_ROOT}/go"
TEST_OUTPUT="$(go test ./... 2>&1)"
echo "${TEST_OUTPUT}" | while IFS= read -r line; do
    echo "  ${line}"
done

echo ""
echo "================================================================================"
echo "Health check diagnostics completed successfully."
echo "================================================================================"
