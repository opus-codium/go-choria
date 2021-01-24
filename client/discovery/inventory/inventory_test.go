package inventory

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/choria-io/go-choria/config"
	"github.com/choria-io/go-choria/protocol"
)

func TestExternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client/Discovery/Inventory")
}

var _ = Describe("Inventory", func() {
	var (
		mockctl *gomock.Controller
		fw      *MockChoriaFramework
		inv     *Inventory
	)

	BeforeEach(func() {
		logger := logrus.New()
		logger.SetOutput(GinkgoWriter)

		mockctl = gomock.NewController(GinkgoT())
		fw = NewMockChoriaFramework(mockctl)
		fw.EXPECT().Logger(gomock.Any()).Return(logrus.NewEntry(logger)).AnyTimes()
		fw.EXPECT().Configuration().Return(config.NewConfigForTests()).AnyTimes()
		inv = New(fw)
		fw.Configuration().Choria.InventoryDiscoverySource = "testdata/good-inventory.yaml"
	})

	AfterEach(func() {
		mockctl.Finish()
	})

	Describe("Discover", func() {
		It("Should resolve nodes", func() {
			filter := protocol.NewFilter()
			nodes, err := inv.Discover(context.Background(), Collective("mt_collective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net"}))

			nodes, err = inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net", "dev2.example.net"}))

			filter.AddFactFilter("country", "==", "mt")
			nodes, err = inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net"}))
		})

		It("Should resolve groups", func() {
			filter := protocol.NewFilter()
			filter.AddIdentityFilter("group:malta")

			nodes, err := inv.Discover(context.Background(), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net"}))

			filter = protocol.NewFilter()
			filter.AddIdentityFilter("group:all")
			nodes, err = inv.Discover(context.Background(), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net", "dev2.example.net"}))

		})

		It("Should resolve multiple groups", func() {
			filter := protocol.NewFilter()
			filter.AddIdentityFilter("group:malta")
			filter.AddIdentityFilter("group:germany")

			nodes, err := inv.Discover(context.Background(), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net", "dev2.example.net"}))
		})

		It("Should match facts", func() {
			filter := protocol.NewFilter()
			filter.AddFactFilter("customer", "==", "acme")
			nodes, err := inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net"}))
		})

		It("Should match classes", func() {
			filter := protocol.NewFilter()
			filter.AddClassFilter("one")
			nodes, err := inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net", "dev2.example.net"}))

			filter.AddClassFilter("three")
			nodes, err = inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev2.example.net"}))
		})

		It("Should match agents", func() {
			filter := protocol.NewFilter()
			filter.AddAgentFilter("rpcutil")
			nodes, err := inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev1.example.net", "dev2.example.net"}))

			filter = protocol.NewFilter()
			filter.AddAgentFilter("rpcutil")
			filter.AddAgentFilter("other")
			nodes, err = inv.Discover(context.Background(), Collective("mcollective"), Filter(filter))
			Expect(err).To(Not(HaveOccurred()))
			Expect(nodes).To(Equal([]string{"dev2.example.net"}))
		})
	})
})
