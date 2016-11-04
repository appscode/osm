package gce

import (
	"encoding/json"

	"cloud.google.com/go/storage"
	"github.com/appscode/osm/pkg/cmd/provider/extpoints"
	osm_context "github.com/appscode/osm/pkg/context"
	"github.com/appscode/osm/pkg/file"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	gcs "google.golang.org/api/storage/v1"
)

type biblio struct{}

func init() {
	extpoints.CloudProviders.Register(new(biblio), "gce")
}

func gcsClient(cred map[string]string, scope ...string) (*storage.Client, error) {
	credGCP, err := json.Marshal(cred)
	conf, err := google.JWTConfigFromJSON(credGCP, scope...)
	if err != nil {
		return nil, err
	}
	httpClient := conf.Client(oauth2.NoContext)

	ctx := context.Background()
	clientOption := option.WithHTTPClient(httpClient)

	client, err := storage.NewClient(ctx, clientOption)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (b *biblio) CreateBucket(bucket string, osmContext *osm_context.Context) error {
	credential, err := file.GetGCSCred(osmContext.ContextData.CredentialDir)
	if err != nil {
		return err
	}
	client, err := gcsClient(credential, gcs.DevstorageReadWriteScope)
	if err != nil {
		return err
	}

	if err := client.Bucket(bucket).Create(context.Background(), credential["project_id"], nil); err != nil {
		return err
	}
	return nil
}
