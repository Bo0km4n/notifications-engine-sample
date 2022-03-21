module github.com/Bo0km4n/notifications-engine-sample

go 1.16

require (
	github.com/argoproj/notifications-engine v0.3.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/spf13/cobra v1.4.0
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime v0.8.3
)

replace github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
