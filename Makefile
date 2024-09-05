BINARY_NAME=golang-boilerplate

hello:
	echo "Hello"
build:
	go build -o ./build/${BINARY_NAME} .
reload-server:
	air
clean:
	go clean
test_coverage:
 	go test ./... -coverprofile=coverage.out
dep:
	go mod download
vet:
	go vet
install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
# lint:
#  	golangci-lint run
format:
	go fmt
run: clean format vet build test_coverage reload-server
