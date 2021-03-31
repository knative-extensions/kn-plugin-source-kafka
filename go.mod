module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.1-0.20210330082320-885f2a9bc996
	knative.dev/eventing-kafka v0.21.1-0.20210330170621-ec5dc671a6a8
	knative.dev/hack v0.0.0-20210325223819-b6ab329907d3
	knative.dev/pkg v0.0.0-20210330162221-808d62257db6
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
