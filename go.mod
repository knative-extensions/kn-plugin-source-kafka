module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/client v0.28.1-0.20220113144412-0f30f0583333
	knative.dev/eventing-kafka v0.28.1-0.20220113122112-9376ea462462
	knative.dev/hack v0.0.0-20220111151514-59b0cf17578e
	knative.dev/pkg v0.0.0-20220113045912-c0e1594c2fb1
)
