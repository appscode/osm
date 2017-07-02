#!/usr/bin/env bash

pushd ~/go/src/github.com/appscode/osm
rm -rf dist
./hack/make.py build; env APPSCODE_ENV=prod ./hack/make.py push; ./hack/make.py push
./hack/docker/setup.sh; env APPSCODE_ENV=prod  ./hack/docker/setup.sh release
popd
