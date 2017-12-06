all: docker

bin/server: server/*
	GOOS=linux go build -a -o bin/server --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo ./server

docker: bin/server Dockerfile html-templates/*
	docker build  -t "tennander/spexflix-gateway" . && rm -rf bin/

