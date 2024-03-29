package main

import (
	"errors"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"strconv"
)

type controller struct {
	jwtSecret string
	users     Users
}

type Users interface {
	getUser(userId int64) (*User, error)
	postUser(user *User) (int64, error)
	queryUsers(email string) ([]*User, error)
}

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	SpexYears int    `json:"spex_years,omitempty"`
}

func (c *controller) getRouter() http.Handler {
	userRoute := mux.NewRouter().StrictSlash(true)
	userRoute.NewRoute().
		Path("/users/{id}").
		Methods("GET").
		HandlerFunc(c.getUser)
	userRoute.NewRoute().
		Path("/users").
		Methods("POST").
		HandlerFunc(c.postUser)
	userRoute.NewRoute().
		Path("/users").
		Methods("GET").
		HandlerFunc(c.getUserByQuery)

	secureHandler := negroni.New()
	secureHandler.Use(c.getJwtMiddleWare())
	secureHandler.UseHandler(userRoute)

	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/healthz", c.healthz)
	mainRouter.HandleFunc("/id", c.getId)
	mainRouter.PathPrefix("/users").Handler(secureHandler)

	n := negroni.Classic()
	n.UseHandler(mainRouter)
	return n
}

func (c *controller) getJwtMiddleWare() negroni.HandlerFunc {
	jwtFunc := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(c.jwtSecret), nil
		},
	}).HandlerWithNext
	return negroni.HandlerFunc(jwtFunc)
}

func (c *controller) getUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := c.users.getUser(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	codecs.JSON.Encode(writer, user)
}

func (c *controller) getId(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	users, err := c.users.queryUsers(email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		http.Error(writer, "could not find user", http.StatusNotFound)
		return
	}
	codecs.JSON.Encode(writer, users[0].Id)
}

func (c *controller) getUserByQuery(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	users, err := c.users.queryUsers(email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	codecs.JSON.Encode(writer, users)
}

func (c *controller) postUser(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	codecs.JSON.Decode(r.Body, user)
	if err := validateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, err := c.users.postUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user.Id = id
	codecs.JSON.Encode(w, user)
}

func validateUser(user *User) error {
	switch {
	case user.Id != 0:
		return errors.New("do not POST user id")
	case user.Name == "":
		return errors.New("user should contain name")
	case user.Email == "":
		return errors.New("user should contain email")
	case user.SpexYears == 0:
		return errors.New("user should contain spex_years")
	default:
		return nil
	}
}

func (c *controller) healthz(writer http.ResponseWriter, request *http.Request) {
	codecs.JSON.Encode(writer, "Everything's fine!")
}

func main() {
	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	dbConfig := dbConfig{
		instanceConnnectionName: os.Getenv("DB_INSTANCE_CONNECTION_NAME"),
		databaseName:            os.Getenv("DB_NAME"),
		user:                    os.Getenv("DB_USER"),
		password:                os.Getenv("DB_PASSWORD"),
	}
	controller := factory(jwtSecret, dbConfig)
	router := controller.getRouter()
	http.ListenAndServe(":"+port, router)
}

func factory(jwtToken string, config dbConfig) *controller {
	users, err := createUserService(config)
	if err != nil {
		panic(err)
	}
	return &controller{
		jwtSecret: jwtToken,
		users:     users,
	}
}
