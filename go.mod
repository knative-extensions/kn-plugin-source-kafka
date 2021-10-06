module knative.dev/kn-plugin-source-kafka

go 1.16

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/client v0.26.1-0.20211005102953-1cb893e48144
	knative.dev/eventing-kafka v0.26.1-0.20211005035853-9fe562d91403
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/pkg v0.0.0-20211005203644-a00ba487121e
)
