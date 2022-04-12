module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	knative.dev/client v0.30.2-0.20220411194606-b702b42859a0
	knative.dev/eventing-kafka v0.30.1-0.20220411133507-bdbcc864cf94
	knative.dev/hack v0.0.0-20220411131823-6ffd8417de7c
	knative.dev/pkg v0.0.0-20220411234407-00c122e376d0
)
