package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/DiTo04/spexflix/common/codecs"
	"os"
	"github.com/urfave/negroni"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type controller struct {
	jwtSecret string
	users Users
}

type Users interface {
	getUser(userId int64) User
}

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	SpexYears int `json:"spex_years"`
}

func (c *controller) getRouter() http.Handler {
	userRoute := mux.NewRouter()
	userRoute.NewRoute().
		Path("/users/{id}").
		Methods("GET").
		HandlerFunc(c.makeGetUserHandler)

	secureHandler := negroni.New()
	secureHandler.Use(c.getJwtMiddleWare())
	secureHandler.UseHandler(userRoute)

	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/healthz", c.healthz)
	mainRouter.PathPrefix("/users").Handler(secureHandler)

	n := negroni.Classic()
	n.UseHandler(mainRouter)
	return n
}

func (c *controller) getJwtMiddleWare() (negroni.HandlerFunc) {
	jwtFunc := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(c.jwtSecret), nil
		},
	}).HandlerWithNext
	return negroni.HandlerFunc(jwtFunc)
}


func (c *controller) makeGetUserHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
	codecs.JSON.Encode(writer, c.users.getUser(id))
}


func (c *controller) healthz(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	codecs.JSON.Encode(writer, "Everything's fine!")
}

func main() {
	port      := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	controller := factory(jwtSecret)
	router := controller.getRouter()
	http.ListenAndServe(":" + port, router)
}

func factory(jwtToken string) *controller {
	return &controller{
		jwtSecret:jwtToken,
		users:&userService{},
	}
}


//claim := request.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
//		id :=int64(claim["id"].(float64))
