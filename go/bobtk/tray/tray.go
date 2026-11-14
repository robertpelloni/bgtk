// Package tray provides the native Windows system-tray icon for the Bobtk
// dashboard (getlantern/systray, pure Go — no cgo). It replaces the earlier
// PowerShell tray (scripts/tray.ps1, kept as fallback) with the same three
// sensible actions: Open Dashboard, live Status, Quit Server.
//
// Run modes (wired in go/main.go):
//   -tray            tray only; controls the server at -addr via its API
//   -serve -tray     server (goroutine) + tray (main thread) in one process;
//                    Quit Server stops both
package tray

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/getlantern/systray"
)

//go:embed icon.ico
var iconBytes []byte

type status struct {
	Version     string `json:"version"`
	Initialized int    `json:"initializedCount"`
	Total       int    `json:"totalCount"`
	UptimeSec   int64  `json:"uptimeSec"`
}

// Run blocks on the tray event loop (must be the main goroutine on Windows).
// baseURL is the dashboard root, e.g. http://localhost:8484. onQuit, when
// non-nil, runs after a successful Quit Server request (combined mode uses
// it to also stop the in-process server by exiting).
func Run(baseURL string, onQuit func()) {
	systray.Run(func() { onReady(baseURL, onQuit) }, func() {})
}

func fetchStatus(baseURL string) (status, error) {
	var st status
	cli := &http.Client{Timeout: 3 * time.Second}
	resp, err := cli.Get(baseURL + "/api/status")
	if err != nil {
		return st, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return st, fmt.Errorf("status %d", resp.StatusCode)
	}
	return st, json.NewDecoder(resp.Body).Decode(&st)
}

func onReady(baseURL string, onQuit func()) {
	systray.SetIcon(iconBytes)
	systray.SetTitle("Bobtk")
	systray.SetTooltip("Bobtk Dashboard — right-click for Open / Status / Quit Server")

	mOpen := systray.AddMenuItem("Open Dashboard", "Open the dashboard in your browser")
	mStatus := systray.AddMenuItem("Status: checking…", "Live pillar status (refreshes when menu opens)")
	mStatus.Disable()
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit Server", "Gracefully stop the dashboard server")
	mExit := systray.AddMenuItem("Exit tray icon (server keeps running)", "Close only the tray icon")

	refresh := func() {
		if st, err := fetchStatus(baseURL); err == nil {
			mStatus.SetTitle(fmt.Sprintf("Status: %d/%d pillars, v%s, up %ds", st.Initialized, st.Total, st.Version, st.UptimeSec))
		} else {
			mStatus.SetTitle("Status: server not reachable at " + baseURL)
		}
	}
	refresh()

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", baseURL).Start()
			case <-mQuit.ClickedCh:
				cli := &http.Client{Timeout: 3 * time.Second}
				if resp, err := cli.Post(baseURL+"/api/quit", "application/json", nil); err == nil {
					resp.Body.Close()
				}
				if onQuit != nil {
					onQuit()
				}
				systray.Quit()
				return
			case <-mExit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Refresh the status line periodically while the tray lives.
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()
		for range t.C {
			refresh()
		}
	}()
}
