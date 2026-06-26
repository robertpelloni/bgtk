package main

import (
	"fmt"

	"github.com/robertpelloni/bobtk/go/src/core"
	"github.com/robertpelloni/bobtk/go/src/visual"
	"github.com/robertpelloni/bobtk/go/src/media"
	"github.com/robertpelloni/bobtk/go/src/network"
	"github.com/robertpelloni/bobtk/go/src/system"
	"github.com/robertpelloni/bobtk/go/src/tools"
)

func main() {
	fmt.Println("Bobtk Go Port Initialization")
	core.Initialize()
	visual.Initialize()
	media.Initialize()
	network.Initialize()
	system.Initialize()
	tools.Initialize()
}
