NAME=jsonbeat
INSTANCE=$(NAME)-run
IMAGE=$(NAME):local
GOPATH?=$(shell pwd)

app-deps:
	cd $(GOPATH)/src/${NAME} && go get

app-build: app-deps
	cd $(GOPATH)/src/${NAME} && go install

app-run: app-build
	./bin/${NAME} -c ./etc/${NAME}.yaml.private -e -v

docker-build: app-build
	docker build -t $(IMAGE) -f Dockerfile .

docker-run:
	docker run --rm --name $(INSTANCE) $(IMAGE)

docker-run-connect:
	@# connect to the existing running docker instance
	docker exec -it $(INSTANCE) /bin/bash

docker-run-bash:
	docker run --rm -i -t $(INSTANCE) /bin/bash
