lint:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v

docker-build:
	docker build -t gotest .

docker-run: docker-build docker-down
	docker run --name my-app --rm -d -p 8080:8181 gotest

docker-down:
	docker rm -f my-app