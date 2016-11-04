package gce

import (
	"github.com/appscode/osm/pkg/cmd/provider/extpoints"
	osm_context "github.com/appscode/osm/pkg/context"
)

type biblio struct{}

func init() {
	extpoints.CloudProviders.Register(new(biblio), "aws")
}

func (b *biblio) CreateBucket(bucket string, osmContext *osm_context.Context) error {
	return nil
}
