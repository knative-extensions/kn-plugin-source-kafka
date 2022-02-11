module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/client v0.29.1-0.20220204171521-6690a20e8f56
	knative.dev/eventing-kafka v0.28.1-0.20220210052605-59a3a1c55de3
	knative.dev/hack v0.0.0-20220209225905-7331bb16ba00
	knative.dev/pkg v0.0.0-20220210201907-fc93ac76d0b6
)
