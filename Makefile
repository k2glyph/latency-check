PROJECT_NAME = latency
install:
	@GO111MODULE=on; go mod tidy
build:
	GOOS=linux go build -o main

test:
	@go test -v main_test.go

clean:
	@echo "cleaning"
	@go clean
	@rm -f main
	@rm -f ${PROJECT_NAME}