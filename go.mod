module knative.dev/kn-plugin-source-kafka

go 1.14

require (
	github.com/maximilien/kn-source-pkg v0.5.0
	github.com/spf13/cobra v1.0.1-0.20200715031239-b95db644ed1c
	github.com/spf13/pflag v1.0.5
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/client v0.19.1
	knative.dev/eventing-kafka v0.19.2
	knative.dev/hack v0.0.0-20201103151104-3d5abc3a0075
	knative.dev/pkg v0.0.0-20201103163404-5514ab0c1fdf
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------
replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
