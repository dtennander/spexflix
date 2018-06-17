package server

import (
	"net/http"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/DiTo04/spexflix/common/codecs"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"cloud.google.com/go/storage"
	"context"
)

type Year struct {
	Year string `json:"year"`
	Name string `json:"name"`
	Eller string `json:"eller"`
	Description string `json:"description"`
	PosterUri string `json:"poster_uri"`
	Uri  string `json:"uri"`
}

type Movie struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Uri string `json:"uri"`
}

type StorageService interface {
	GetYears(ctx context.Context) ([]Year, error)
	GetContent(ctx context.Context, year string) ([]Movie, error)
}

type controller struct {
	storageService StorageService
	jwtSecret string
}

func CreateController(jwtSecret string, bucketName string) (*controller, error) {
	storageService, err := createStorageService(bucketName)
	if err != nil {
		return nil, err
	}
	return &controller{
		storageService: storageService,
		jwtSecret:      jwtSecret,
	}, nil
}

func createStorageService(bucketName string) (*cloudStorageService, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(bucketName)
	storageService := &cloudStorageService{
		client: bucket,
	}
	return storageService, nil
}

func (c *controller) CreateRoutes() *negroni.Negroni {
	moviesRouter := mux.NewRouter()
	moviesRouter.HandleFunc("/movies", c.listYears)
	moviesRouter.HandleFunc("/movies/{year}", c.listContent)

	secureHandler := negroni.New()
	secureHandler.Use(c.getJwtMiddleWare())
	secureHandler.UseHandler(moviesRouter)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", c.healthz)
	r.PathPrefix("/movies").Handler(secureHandler)
	n := negroni.Classic()
	n.UseHandler(r)
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

func (c *controller) listYears(w http.ResponseWriter, r *http.Request) {
	years, err := c.storageService.GetYears(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	codecs.JSON.Encode(w, years)
}

func (c *controller) listContent(w http.ResponseWriter, r *http.Request) {
	year := mux.Vars(r)["year"]
	content, err := c.storageService.GetContent(r.Context(), year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	codecs.JSON.Encode(w, content)
}

func (c *controller) getContent(w http.ResponseWriter, r *http.Request) {


}

func (c *controller) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
