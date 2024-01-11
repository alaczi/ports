package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"port_domain_service/services"
	"testing"
)

func TestPort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config suite")
}

var _ = Describe("Tests for configuration service", func() {
	BeforeEach(func() {
		os.Unsetenv("DOMAIN_PORT")
	})
	AfterEach(func() {
		os.Unsetenv("DOMAIN_PORT")
	})
	Context("When the environment is not set", func() {
		It("uses the default values", func() {
			config, err := services.NewConfig()
			Expect(config.Port).To(Equal(50051))
			Expect(err).To(BeNil())
		})
	})

	Context("When the environment is set", func() {
		It("uses the values from the environment", func() {
			os.Setenv("DOMAIN_PORT", "4041")
			config, err := services.NewConfig()
			Expect(config.Port).To(Equal(4041))
			Expect(err).To(BeNil())
		})
	})

	Context("When the environment is set", func() {
		It("uses the values from the environment", func() {
			os.Setenv("DOMAIN_PORT", "not-number")
			config, err := services.NewConfig()
			Expect(config).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})