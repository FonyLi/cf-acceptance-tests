package apps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"

	. "github.com/pivotal-cf-experimental/cf-test-helpers/cf"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/generator"
)

var _ = Describe("An application being staged", func() {
	var AppName string

	BeforeEach(func() {
		AppName = RandomName()
	})

	AfterEach(func() {
		Expect(Cf("delete", AppName, "-f")).To(Say("OK"))
	})

	It("has its staging log streamed during a push", func() {
		push := Cf("push", AppName, "-p", TestAssets.Dora)

		// Expect(push).To(Say("Installing dependencies"))
		Expect(push).To(Say("Uploading droplet"))
		Expect(push).To(Say("App started"))
	})
})
