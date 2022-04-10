build:
	go build .
run:
	go run .
generate-mocks:
	sh tools/generate-mocks.sh
test-unit:
	go test -count=1 ./...
test-coverage:
	go test -count=1 -coverprofile=coverage.xml ./...
#	go tool cover -html=coverage.out -o coverage.html
#	open coverage.html
lint:
	golangci-lint run -c ./.github/.golangci.yml
