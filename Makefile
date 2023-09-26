PHONY: handler
handler:
	FILE_NAME=${FILE_NAME} go generate ./script/generator/handler
