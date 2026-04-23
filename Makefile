hello:
	cd hello && echo "Hello"

arrays-code:
	cd arrays && go run main.go

maps-code:
	cd maps && go run main.go

functional-tests:
	ginkgo -v ./address-api-tests ./ginkgo-tests

deps:
	go mod download