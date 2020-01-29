PROJECT_NAME = latency
TAG_NAME = latency
install:
	@GO111MODULE=on; go mod tidy
build:
	GOOS=linux go build -o ${PROJECT_NAME}

test:
	@go test -v main_test.go

clean:
	@echo "cleaning"
	@go clean
	@rm -f main
	@rm -f ${PROJECT_NAME}
