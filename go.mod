module knative.dev/kn-plugin-source-kafka

go 1.14

require (
	github.com/maximilien/kn-source-pkg v0.5.0
	github.com/spf13/cobra v1.0.1-0.20200715031239-b95db644ed1c
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/client v0.18.4
	knative.dev/eventing-contrib v0.18.0
	knative.dev/hack v0.0.0-20201120192952-353db687ec5b
	knative.dev/pkg v0.0.0-20200922164940-4bf40ad82aab
	knative.dev/test-infra v0.0.0-20200921012245-37f1a12adbd3
)

// Temporary pinning certain libraries. Please check periodically, whether these are still needed
// ----------------------------------------------------------------------------------------------
replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
