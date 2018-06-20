package main

import (
	"github.com/DiTo04/spexflix/content/server"
	"net/http"
	"os"
)

var (
	jwtSecret   = os.Getenv("JWT_SECRET")
	serverPort  = os.Getenv("SERVER_PORT")
	bucketName  = os.Getenv("MEDIA_BUCKET_NAME")
	keyFilePath = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
)

func main() {
	controller, err :=
		server.CreateController(jwtSecret, bucketName, keyFilePath)
	if err != nil {
		panic(err)
	}
	routes := controller.CreateRoutes()
	http.ListenAndServe("0.0.0.0:"+serverPort, routes)
}
