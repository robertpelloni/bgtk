// Command bobtk is the Go entry point for the 6-pillar framework.
//
// Default run (no flags) keeps the original behavior: initialize all six
// pillars, print progress, exit. Dashboard mode (-serve) keeps the process
// alive serving the consolidated single-page dashboard described in
// docs/manual/MANUAL.md.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bobtk/bobtk/bobtk/core"
	"github.com/bobtk/bobtk/bobtk/dashboard"
	"github.com/bobtk/bobtk/bobtk/media"
	"github.com/bobtk/bobtk/bobtk/network"
	"github.com/bobtk/bobtk/bobtk/system"
	"github.com/bobtk/bobtk/bobtk/tools"
	"github.com/bobtk/bobtk/bobtk/tray"
	"github.com/bobtk/bobtk/bobtk/visual"
)

func main() {
	serve := flag.Bool("serve", false, "serve the consolidated dashboard instead of exiting after init")
	addr := flag.String("addr", ":8484", "dashboard listen address (with -serve)")
	withTray := flag.Bool("tray", false, "show the native system-tray icon (with -serve: same process; alone: control the server at -addr)")
	showVersion := flag.Bool("version", false, "print version (from VERSION file) and exit")
	flag.Parse()

	version := dashboard.ReadVersion("1.1.0")
	if *showVersion {
		fmt.Println(version)
		return
	}
	dashURL := "http://" + *addr
	if strings.HasPrefix(*addr, ":") {
		dashURL = "http://localhost" + *addr
	}
	if *serve && *withTray {
		srv := dashboard.New(version)
		go func() {
			log.Printf("Bobtk %s dashboard starting on %s (tray mode)", version, *addr)
			if err := srv.Run(*addr); err != nil {
				log.Printf("dashboard stopped: %v", err)
			}
		}()
		tray.Run(dashURL, func() { os.Exit(0) })
		return
	}
	if *withTray {
		tray.Run(dashURL, nil)
		return
	}
	if *serve {
		srv := dashboard.New(version)
		log.Printf("Bobtk %s dashboard starting on %s", version, *addr)
		if err := srv.Run(*addr); err != nil {
			log.Fatalf("dashboard: %v", err)
		}
		return
	}

	fmt.Println("Welcome to the Bobtk Go Port")
	fmt.Println("Initializing 6-Pillar Framework...")

	corePillar := core.NewCorePillar()
	corePillar.Initialize()

	mediaPillar := media.NewMediaPillar()
	mediaPillar.Initialize()

	networkPillar := network.NewNetworkPillar()
	networkPillar.Initialize()

	systemPillar := system.NewSystemPillar()
	systemPillar.Initialize()

	toolsPillar := tools.NewToolsPillar()
	toolsPillar.Initialize()

	visualPillar := visual.NewVisualPillar()
	visualPillar.Initialize()

	fmt.Println("Bobtk Go Port Initialization Complete.")
	os.Exit(0)
}
