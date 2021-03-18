module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.1-0.20210316000941-6a4f308e5e73
	knative.dev/eventing-kafka v0.21.1-0.20210316174342-cca17c800a6b
	knative.dev/hack v0.0.0-20210309141825-9b73a256fd9a
	knative.dev/pkg v0.0.0-20210315160101-6a33a1ab29ac
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
