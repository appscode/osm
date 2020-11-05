module github.com/appscode/osm

go 1.12

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/spf13/cobra v1.1.1
	gomodules.xyz/runtime v0.0.0-20201104200926-d838b09dda8b
	gomodules.xyz/stow v0.2.3
	gomodules.xyz/x v0.0.0-20201105065653-91c568df6331
	k8s.io/client-go v0.18.9
	kmodules.xyz/client-go v0.0.0-20201105071625-0b277310b9b8
	sigs.k8s.io/yaml v1.2.0
)

replace cloud.google.com/go => cloud.google.com/go v0.38.0

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2

replace k8s.io/api => github.com/kmodules/api v0.18.10-0.20200922195318-d60fe725dea0

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200922195535-0c9a1b86beec

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.18.10-0.20200922195747-1bd1cc8f00d1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.9

replace k8s.io/client-go => github.com/kmodules/k8s-client-go v0.18.10-0.20200922201634-73fedf3d677e

replace k8s.io/component-base => k8s.io/component-base v0.18.9

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200922200158-8b13196d8dc4

replace k8s.io/utils => k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
