# Unified Bobtk build script — one entry point replacing scattered commands.
# Default: Go dashboard build + vet + test (fast, offline). -Native adds the
# Meson C++ build. Heavy work runs at BelowNormal priority on request so the
# laptop stays responsive.
# Usage:
#   powershell -File scripts/build.ps1
#   powershell -File scripts/build.ps1 -Native -Priority BelowNormal
#   powershell -File scripts/build.ps1 -Serve   (build, then run dashboard)
param(
    [switch]$Native,
    [switch]$Serve,
    [string]$Addr = ":8484",
    [ValidateSet("Normal", "BelowNormal", "Idle")][string]$Priority = "BelowNormal"
)
$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
$proc = [System.Diagnostics.Process]::GetCurrentProcess()
try { $proc.PriorityClass = $Priority } catch { Write-Host "priority unchanged: $($_.Exception.Message)" }

Push-Location (Join-Path $root "go")
try {
    Write-Host "== go vet ==";   go vet ./...
    Write-Host "== go test ==";  go test ./...
    Write-Host "== go build =="; go build -o bobtk.exe .
    Write-Host "Go build OK: go/bobtk.exe"
    if ($Serve) { Write-Host "Serving dashboard on $Addr (Ctrl+C to stop)"; & .\bobtk.exe -serve -addr $Addr }
} finally { Pop-Location }

if ($Native) {
    if (-not (Get-Command meson -ErrorAction SilentlyContinue)) { throw "meson not on PATH — see DEPLOY.md" }
    Push-Location $root
    try {
        if (-not (Test-Path "_build/build.ninja")) { meson setup _build }
        meson compile -C _build
        Write-Host "Native build OK"
    } finally { Pop-Location }
}
