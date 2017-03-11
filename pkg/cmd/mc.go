package cmd

import (
	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type bucketCreateRequest struct {
	context string
	bucket  string
}

func NewCmdCreate() *cobra.Command {
	req := &bucketCreateRequest{}
	cmd := &cobra.Command{
		Use:     "mc bucket-name",
		Short:   "Create Bucket",
		Example: "osm mc mybucket",
		Run: func(cmd *cobra.Command, args []string) {
			req.bucket = args[0]
			createBucket(req)
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "The name of the osmconfig context to use")
	return cmd
}

func createBucket(req *bucketCreateRequest) {
	cfg, err := otx.LoadConfig()
	if err != nil {
		term.Fatalln(err)
	}
	if req.context != "" {
		req.context = cfg.CurrentContext
	}
	var ctx *otx.Context
	for _, osmCtx := range cfg.Contexts {
		if osmCtx.Name == req.context {
			ctx = osmCtx
			break
		}
	}
	if ctx == nil {
		term.Fatalln("Missing context")
	}

	//provider := extpoints.CloudProviders.Lookup(contextData.ContextData.Provider)
	//if err := provider.CreateBucket(req.bucket, contextData); err != nil {
	//	printer.Error(err)
	//	os.Exit(1)
	//}
}
