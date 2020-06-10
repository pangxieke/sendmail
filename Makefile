IMAGE=sendmail
TAG1=$(shell git describe --always --tags)
TAG=${TAG1}_v1_20200609.2

REGISTRY=registry.cn-shenzhen.aliyuncs.com/pangxieke

default: all

fmt:
	go fmt ./...

all: fmt
	docker build -t ${REGISTRY}/${IMAGE}:${TAG} .

push:
	docker push ${REGISTRY}/${IMAGE}:${TAG}

publish: all push

builder:
	docker build -t go-builder:latest . -f builder.dockerfile
	docker push go-builder:latest

test:
	go test -v ./...

lint:
	ls -l | grep '^d' | awk '{print $$NF}' | grep -v vender | xargs golint

clear:
	rm -f ./*.log

docker-clean:
	docker images
	docker image prune --force

.PHONY: test pb stages

