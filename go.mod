module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/client v0.28.1-0.20220114130644-a96ecf55f6de
	knative.dev/eventing-kafka v0.28.1-0.20220114075342-c97cbe0e8448
	knative.dev/hack v0.0.0-20220111151514-59b0cf17578e
	knative.dev/pkg v0.0.0-20220113045912-c0e1594c2fb1
)
