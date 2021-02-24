module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.0
	knative.dev/eventing-kafka v0.21.0
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
	knative.dev/pkg v0.0.0-20210216013737-584933f8280b
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3

// added temporary until the other PR is merged
replace github.com/maximilien/kn-source-pkg => github.com/rhuss/kn-source-pkg v0.0.0-20210224093942-712ccae658de
