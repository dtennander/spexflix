all: docker

bin/server: main.go */*.go
	GOOS=linux go build -a -o bin/server --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo .

docker: bin/server Dockerfile
	docker build  -t "tennander/spexflix-authentication" . && rm -rf bin/

