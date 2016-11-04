package extpoints

import osm_context "github.com/appscode/osm/pkg/context"

type CloudProvider interface {
	CreateBucket(string, *osm_context.Context) error
}
