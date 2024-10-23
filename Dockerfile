FROM ubuntu:latest
LABEL authors="sunviv"
COPY k8s-demo /app/k8s-demo
WORKDIR /app
CMD ["/app/k8s-demo"]