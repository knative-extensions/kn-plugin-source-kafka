module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/client v0.26.1-0.20211019150534-534d91319f7d
	knative.dev/eventing-kafka v0.26.1-0.20211019171034-8b3f08b7e87f
	knative.dev/hack v0.0.0-20211019034732-ced8ce706528
	knative.dev/pkg v0.0.0-20211019132235-ba2b2b1bf268
)
