# Bobtk system-tray icon (Windows).
# Why a script: this repo has no native tray code (GTK4 removed GtkStatusIcon
# and no systray implementation exists), so the supported tray is this small
# WinForms NotifyIcon. It offers the two sensible actions plus status:
#   - Open Dashboard (also on double-click)
#   - Quit Server  (graceful POST /api/quit — stops only the dashboard process)
# Usage:  powershell -ExecutionPolicy Bypass -File scripts/tray.ps1 [-Url http://localhost:8484]
param([string]$Url = "http://localhost:8484")

Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

$notify = New-Object System.Windows.Forms.NotifyIcon
$notify.Icon = [System.Drawing.SystemIcons]::Application
$notify.Text = "Bobtk Dashboard"
$notify.Visible = $true

$menu = New-Object System.Windows.Forms.ContextMenuStrip
$open = $menu.Items.Add("Open Dashboard")
$status = $menu.Items.Add("Status: checking...")
$status.Enabled = $false
$quit = $menu.Items.Add("Quit Server")
$exit = $menu.Items.Add("Exit tray icon (server keeps running)")
$notify.ContextMenuStrip = $menu

function Get-BobtkStatus {
    try {
        $s = Invoke-RestMethod -Uri "$Url/api/status" -TimeoutSec 3
        return "Running: $($s.initializedCount)/$($s.totalCount) pillars, v$($s.version), up $($s.uptimeSec)s"
    } catch { return "Server not reachable at $Url" }
}
$status.Text = "Status: $(Get-BobtkStatus)"
$open.add_Click({ Start-Process $Url })
$notify.add_DoubleClick({ Start-Process $Url })
$quit.add_Click({
    try { Invoke-RestMethod -Uri "$Url/api/quit" -Method Post -TimeoutSec 3 | Out-Null } catch {}
    $notify.Visible = $false; $notify.Dispose(); [System.Windows.Forms.Application]::Exit()
})
$exit.add_Click({ $notify.Visible = $false; $notify.Dispose(); [System.Windows.Forms.Application]::Exit() })
$menu.add_Opening({ $status.Text = "Status: $(Get-BobtkStatus)" })

[System.Windows.Forms.Application]::Run()
