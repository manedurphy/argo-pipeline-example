compile:
	go build -o _output/ping .

buildx:
	docker buildx build -t manedurphy/ping-pong --platform linux/arm64,linux/amd64 . --push

docker-run:
	docker run -d --rm -p 8080:8080 --name ping-pong manedurphy/ping-pong

k8s-run:
	kubectl run ping-pong --image=manedurphy/ping-pong