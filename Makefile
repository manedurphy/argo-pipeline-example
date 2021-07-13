VERSION=1.1.1
COMMIT=$(shell git log -n 1 --pretty=format:'%h')

compile:
	go build -o _output/ping .

docker-push:
	docker buildx build -t manedurphy/ping-pong:${VERSION} --platform linux/arm64,linux/amd64 . --push

docker-push-pr:
	docker buildx build -t manedurphy/ping-pong:development --platform linux/arm64,linux/amd64 . --push

docker-run:
	docker run -d --rm -p 8080:8080 --name ping-pong manedurphy/ping-pong

k8s-run:
	kubectl run ping-pong --image=manedurphy/ping-pong

pr:
	go run scripts/generate_pr_env.go --commit=${COMMIT}
