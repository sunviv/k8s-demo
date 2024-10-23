.PHONY: docker
docker:
	@rm k8s-demo || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o k8s-demo .
	@docker rmi -f sunviv/k8s-demo:v0.0.1
	@docker build -t sunviv/k8s-demo:v0.0.1 .