package port_test

import (
	"context"
	repo "github.com/alaczi/ports/repository"
	"port_domain_service/domain/port"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Port Suite")
}

var _ = Describe("Tests for in memory repository", func() {
	var port1, port2 *repo.Port
	var testRepo *port.InMemoryPortRepository

	BeforeEach(func() {
		testRepo = port.NewInMemoryPortRepository()
		province := "Ajman"
		port1 = &repo.Port{
			Id:          "AEAJM",
			Code:        "52000",
			Name:        "Ajman",
			City:        "Ajman",
			Country:     "United Arab Emirates",
			Timezone:    "Asia/Dubai",
			Province:    &province,
			Coordinates: []float32{55.513645, 25.405216},
			Unlocs:      []string{"AEAJM"},
		}
		province2 := "Khulna Division"
		port2 = &repo.Port{
			Id:      "BDMGL",
			Name:    "Mongla",
			City:    "Mongla",
			Country: "Bangladesh",
			Alias:   []string{},
			Regions: []string{},
			Coordinates: []float32{
				89.6016171,
				22.4942196,
			},
			Province: &province2,
			Timezone: "Asia/Dhaka",
			Unlocs: []string{
				"BDMGL",
			},
		}
	})

	It("Return nil pointer when the given port does not exist", func() {
		testPort, err := testRepo.GetPort(context.Background(), port1.Id)
		Expect(testPort).To(BeNil())
		Expect(err).To(BeNil())
	})

	It("Successfully adds a port, and returns the given port when the port exists", func() {
		err := testRepo.UpsertPort(context.Background(), port2)
		testPort, err := testRepo.GetPort(context.Background(), port2.Id)
		Expect(testPort).To(Equal(port2))
		Expect(err).To(BeNil())
	})

	It("Successfully upserts a port", func() {
		err := testRepo.UpsertPort(context.Background(), port1)
		err = testRepo.UpsertPort(context.Background(), port2)
		testPort, err := testRepo.GetPort(context.Background(), port2.Id)
		Expect(testPort).To(Equal(port2))
		Expect(err).To(BeNil())
		updatedPort := &*port2
		updatedPort.City = "SomeOtherCity"
		err = testRepo.UpsertPort(context.Background(), updatedPort)
		Expect(err).To(BeNil())

		testPort, err = testRepo.GetPort(context.Background(), port2.Id)
		Expect(testPort).To(Equal(updatedPort))
		Expect(testPort.City).To(Equal("SomeOtherCity"))
		Expect(err).To(BeNil())
	})
})