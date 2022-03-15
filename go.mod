module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/client v0.30.2-0.20220314132918-d75836c4f42a
	knative.dev/eventing-kafka v0.30.1-0.20220314063517-bc254c796e2f
	knative.dev/hack v0.0.0-20220314052818-c9c3ea17a2e9
	knative.dev/pkg v0.0.0-20220315095603-616f1ab878c5
)
