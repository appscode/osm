## osm config set-context

Set context

### Synopsis


Set context

```
osm config set-context <name> [flags]
```

### Examples

```
osm config set-context <name>
```

### Options

```
      --azure.account string           Azure config account
      --azure.key string               Azure config key
      --google.json_key_path string    GCS config json key path
      --google.project_id string       GCS config project id
      --google.scopes string           GCS config scopes
  -h, --help                           help for set-context
      --local.path string              Local config key path
      --provider string                Cloud storage provider
      --s3.access_key_id string        S3 config access key id
      --s3.auth_type string            S3 config auth type (accesskey, iam) (default "accesskey")
      --s3.disable_ssl                 S3 config disable SSL
      --s3.endpoint string             S3 config endpoint
      --s3.region string               S3 config region
      --s3.secret_key string           S3 config secret key
      --swift.auth_token string        Swift AuthToken
      --swift.domain string            Swift domain
      --swift.key string               Swift config key
      --swift.region string            Swift region
      --swift.storage_url string       Swift StorageURL
      --swift.tenant_auth_url string   Swift teanant auth url
      --swift.tenant_domain string     Swift TenantDomain
      --swift.tenant_id string         Swift TenantId
      --swift.tenant_name string       Swift tenant name
      --swift.trust_id string          Swift TrustId
      --swift.username string          Swift username
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --osmconfig string                 Path to osm config (default "/home/tamal/.osm/config")
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [osm config](osm_config.md)	 - OSM configuration

