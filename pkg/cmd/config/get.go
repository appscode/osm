package config

import (
	"os"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newCmdGet() *cobra.Command {
	setCmd := &cobra.Command{
		Use:     "get-contexts",
		Short:   "Get contexts",
		Example: "osm config get-contexts",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}
			getContexts()
		},
	}
	return setCmd
}

func getContexts() {
	config, err := otx.LoadConfig()
	term.ExitOnError(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTRE)
	table.SetHeader([]string{"CURRENT", "NAME", "PROVIDER"})
	ctx := config.CurrentContext
	for _, osmCtx := range config.Contexts {
		if osmCtx.Name == ctx {
			table.Append([]string{"*", osmCtx.Name, osmCtx.Provider})
		} else {
			table.Append([]string{"", osmCtx.Name, osmCtx.Provider})
		}
	}
	table.Render()
}
