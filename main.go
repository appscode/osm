//go:generate go-extpoints pkg//cmd/provider/extpoints

package main

import (
	"os"
	"runtime"
	"github.com/appscode/osm/pkg/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd := cmd.NewCmd()
	// execute commands.
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
