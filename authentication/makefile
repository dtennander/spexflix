all: docker

bin/server: server/server.go
	GOOS=linux go build -a -o bin/server --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo ./server

docker: bin/server Dockerfile
	docker build  -t "tennander/spexflix-authentication" . && rm -rf bin/

