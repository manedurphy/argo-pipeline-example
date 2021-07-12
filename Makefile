VERSION=1.2.1

compile:
	go build -o _output/ping .

docker-push:
	docker buildx build -t manedurphy/ping-pong:${VERSION} --platform linux/arm64,linux/amd64 . --push

docker-run:
	docker run -d --rm -p 8080:8080 --name ping-pong manedurphy/ping-pong

k8s-run:
	kubectl run ping-pong --image=manedurphy/ping-pong
