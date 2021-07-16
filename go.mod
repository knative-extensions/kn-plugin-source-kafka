module knative.dev/kn-plugin-source-kafka

go 1.15

require (
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/client v0.24.1-0.20210714122431-0e07481fa61f
	knative.dev/eventing-kafka v0.24.1-0.20210715160032-7cfcc9782cf8
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210715175632-d9b7180af6f2
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
