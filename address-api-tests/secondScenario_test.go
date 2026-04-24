package address_api_tests_test

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecondScenario", func() {

	BeforeEach(func() {
		log.Println(" we are inside")
	})

	Context("When", func() {

		It("addition shoud be as expected", func() {
			a := 2 + 3
			Expect(a).To(Equal(5))
		})

		It("multiplication shoud be as expected", func() {
			m := 3 * 3
			Expect(m).To(Equal(5))
		})
	})

	AfterEach(func() {
		log.Println(" We are outside")
	})
})
