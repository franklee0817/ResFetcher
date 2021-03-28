package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/TarsCloud/ResFetcher/TarsTestToolKit"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(fetcherImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("fetcherImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(TarsTestToolKit.Fetcher)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".fetcherObj")

	// Run application
	tars.Run()
}