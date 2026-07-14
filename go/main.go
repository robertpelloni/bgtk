package main

import (
	"fmt"
	"github.com/bobtk/bobtk/bobtk/core"
	"github.com/bobtk/bobtk/bobtk/media"
	"github.com/bobtk/bobtk/bobtk/network"
	"github.com/bobtk/bobtk/bobtk/system"
	"github.com/bobtk/bobtk/bobtk/tools"
	"github.com/bobtk/bobtk/bobtk/visual"
)

func main() {
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
}
