lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v