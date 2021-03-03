module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.1-0.20210301230649-270e32240152
	knative.dev/eventing-kafka v0.21.1-0.20210303075215-b7499c4d956d
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
	knative.dev/pkg v0.0.0-20210303111915-08fc6268bf96
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
