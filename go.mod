module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/client v0.24.1-0.20210709134519-4bb960927335
	knative.dev/eventing-kafka v0.24.1-0.20210708072323-21cafa72e334
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210708145023-4a3e56dc13b2
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
