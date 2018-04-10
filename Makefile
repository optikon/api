.PHONY: build container

EXECUTABLE ?= optikon-api
IMAGE ?= bin/$(EXECUTABLE)
REPO = dockerhub.cisco.com/intelligent-edge-dev-docker-local/optikon-api
TAG = 0.1

all: build

build:
	CGO_ENABLED=0 go build --ldflags '${EXTLDFLAGS}' -o ${IMAGE} wwwin-github.cisco.com/edge/optikon-api/api/v0/cmd

container:
	docker run -t -w /go/src/wwwin-github.cisco.com/edge/optikon -v `pwd`:/go/src/wwwin-github.cisco.com/edge/optikon golang:1.10.1 make
	docker build -t $(REPO):$(TAG) .
