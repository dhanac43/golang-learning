package godog_tests

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

func TestName(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	status := suite.Run()

	if status != 0 {
		t.Fatalf("expected 0 but got %d", status)
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^there are (\d+) godogs$`, functionOne)
	sc.When(`^I eat (\d+)$`, functionTwo)
	sc.Then(`^there should be (\d+) remaining$`, functionThree)
}

func functionOne(n int) error {
	fmt.Println("There are", n, "godogs")
	return nil
}

func functionTwo(n int) error {
	fmt.Println("I eat", n)
	return nil
}

func functionThree(n int) error {
	fmt.Println("Remaining:", n)
	return nil
}
