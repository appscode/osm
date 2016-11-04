package util

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Checks if a flag value in a command has been provided by the user
// Or not. The ordering of the flags can be set for nested flags.
func EnsureRequiredFlags(cmd *cobra.Command, name ...string) {
	for _, n := range name {
		flag := cmd.Flag(n)
		if flag == nil {
			continue
		}
		if !flag.Changed {
			color.Set(color.FgRed, color.Bold)
			fmt.Printf("flag [--%v] is required but not provided.\n", flag.Name)
			os.Exit(1)
		}
	}
}
