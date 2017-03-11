[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Forum](https://discuss.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# osm
Object Store Manipulator - `curl` for cloud storage services. 🙌

## Usage
```bash
osm [command] [flags]
osm [command]

Available Commands:
  config      OSM configuration
  help        Help about any command
  ls          List container
  mc          Make container
  pull        Pull item from container
  push        Push item from container
  rc          Remove container
  rm          Remove item from container
  stat        Stat item from container
  version     Prints binary version number.

Flags:
      --alsologtostderr                  log to standard error as well as files
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging

Use "osm [command] --help" for more information about a command.

```

## Build

    ./hack/make.py

## Set Context
    osm config set-context gcs -p gce -c /var/credential/gce
    osm config use-context gcs

## View Context

    osm config view

Config YAML:

    contexts:
    - context:
        credential_dir: /var/credential/gce
        provider: gce
      name: gcs
    current-context: gcs

## Create Bucket
    osm create -b db-box

This will create `db-box` bucket in gcs

You can also use `context` while creating

    osm create -b db-box --context=gcs
