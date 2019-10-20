prepare:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/sys/unix
	go get -u github.com/jinzhu/configor
	go get -u go.uber.org/zap
	go get -u github.com/segmentio/kafka-go
	go get -u golang.org/x/net/context

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

init:
	dep init

deps:
	dep ensure -v

dkb:
	docker build -t superhero-register .

dkr:
	docker run --rm -p "3000:3000" superhero-register

launch: dkb dkr

api-log:
	docker logs api -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it api /bin/bash

db-ssh:
	docker exec -it db /bin/bash

PHONY: prepare build init deps dkb dkr launch api-log rmc rmi clear