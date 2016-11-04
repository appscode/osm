package cmd

import (
	"os"

	_ "github.com/appscode/osm/pkg/cmd/provider/aws"
	"github.com/appscode/osm/pkg/cmd/provider/extpoints"
	_ "github.com/appscode/osm/pkg/cmd/provider/gce"
	osm_context "github.com/appscode/osm/pkg/context"
	"github.com/appscode/osm/pkg/printer"
	"github.com/appscode/osm/pkg/util"
	"github.com/spf13/cobra"
)

type bucketCreateRequest struct {
	context string
	bucket  string
}

func newCmdCreate() *cobra.Command {
	req := &bucketCreateRequest{}
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create Bucket",
		Example: "osm create -b test-bucket",
		Run: func(cmd *cobra.Command, args []string) {
			util.EnsureRequiredFlags(cmd, "bucket")
			createBucket(req)
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "The name of the osmconfig context to use")
	cmd.Flags().StringVarP(&req.bucket, "bucket", "b", "", "The name of the bucket")
	return cmd
}

func createBucket(req *bucketCreateRequest) {
	config, err := osm_context.GetConfigData()
	if err != nil {
		printer.Error(err)
		return
	}
	var currentContext string = config.CurrentContext
	if req.context != "" {
		currentContext = req.context
	}
	var contextData *osm_context.Context
	for _, osmCtx := range config.Contexts {
		if osmCtx.Name == currentContext {
			contextData = osmCtx
		}
	}

	if contextData == nil {
		printer.Error("Invalid context")
		os.Exit(1)
	}

	provider := extpoints.CloudProviders.Lookup(contextData.ContextData.Provider)
	if err := provider.CreateBucket(req.bucket, contextData); err != nil {
		printer.Error(err)
		os.Exit(1)
	}
}
