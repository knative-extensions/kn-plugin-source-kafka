module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/client v0.26.1-0.20211028082427-f027b38e200a
	knative.dev/eventing-kafka v0.26.1-0.20211028065026-ee67c5013735
	knative.dev/hack v0.0.0-20211027200727-f1228dd5e3e6
	knative.dev/pkg v0.0.0-20211027171921-f7b70f5ce303
)
