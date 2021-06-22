module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/client v0.23.1-0.20210618090750-a54abbdfa1a2
	knative.dev/eventing-kafka v0.23.1-0.20210622065127-98b1b1aa5960
	knative.dev/hack v0.0.0-20210614141220-66ab1a098940
	knative.dev/pkg v0.0.0-20210618060751-f454995ff92b
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
