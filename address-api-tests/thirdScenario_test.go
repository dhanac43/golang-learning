package address_api_tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ThirdScenario", func(a, b, c int) {

	result := a + b
	Expect(result).To(Equal(c))

},
	Entry("test with combitation 1", 2, 3, 4),
	Entry("Test with combination 2", 3, 3, 6),
)
