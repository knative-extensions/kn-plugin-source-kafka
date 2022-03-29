module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	knative.dev/client v0.30.2-0.20220329090315-1bab8209ceab
	knative.dev/eventing-kafka v0.30.1-0.20220322160026-ad099b074a0f
	knative.dev/hack v0.0.0-20220328133751-f06773764ce3
	knative.dev/pkg v0.0.0-20220325200448-1f7514acd0c2
)
