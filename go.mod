module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/client v0.29.1-0.20220217100713-af052088caa5
	knative.dev/eventing-kafka v0.28.1-0.20220217082312-7e15b864e492
	knative.dev/hack v0.0.0-20220216040439-0456e8bf6547
	knative.dev/pkg v0.0.0-20220215153400-3c00bb0157b9
)
