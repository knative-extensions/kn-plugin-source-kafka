module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.2-0.20210219230238-c655d7ca141e
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.20.1-0.20210219184703-18ff59604195
	knative.dev/eventing-kafka v0.19.1-0.20210219230103-ae8020f1c58b
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
	knative.dev/pkg v0.0.0-20210219164703-8a9bf766d36d
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
