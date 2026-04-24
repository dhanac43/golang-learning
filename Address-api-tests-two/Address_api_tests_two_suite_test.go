package Address_api_tests_two_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAddressApiTestsTwo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AddressApiTestsTwo Suite")
}
