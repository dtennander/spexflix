all: docker

bin/server: server.go
	GOOS=linux go build -a -o bin/server --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo .

docker: bin/server Dockerfile
	docker build  -t "tennander/spexflix-content" . && rm -rf bin/

