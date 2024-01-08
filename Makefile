PHONY: handler
handler:
	FILE_NAME=${FILE_NAME} go generate ./script/generator/handler

PHONY: usecase
usecase:
	FILE_NAME=${FILE_NAME} go generate ./script/generator/usecase

PHONY: start
start:
	go run cmd/main.go
