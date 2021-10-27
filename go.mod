module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/client v0.26.1-0.20211026094721-2939f36bbd88
	knative.dev/eventing-kafka v0.26.1-0.20211019171034-8b3f08b7e87f
	knative.dev/hack v0.0.0-20211026141922-a71c865b5f66
	knative.dev/pkg v0.0.0-20211026205101-a8fb29270197
)
