# osm
Object Store Manipulator - `curl` for cloud storage services. 🙌

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
