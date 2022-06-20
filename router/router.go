package router

import (
	"log"
	"mvp/repository"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	DB     *repository.DB
	Logger *log.Logger
}

func New(db *repository.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	h := Handlers{DB: db, Logger: log.Default()}

	// Create a route along /files that will serve contents from
	// the ./front/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "front"))
	FileServer(r, "/mvp", filesDir)

	r.Route("/product", func(r chi.Router) {
		r.Get("/list", h.List)
		r.Get("/list/{name}", h.Search)
		r.Get("/one/{productID}", h.One)
	})
	r.Route("/order", func(r chi.Router) {
		r.Post("/buy", h.Buy)
	})
	r.Route("/basket", func(r chi.Router) {
		r.Post("/add", h.AddBasket)
		r.Get("/{userID}", h.GetBasket)
	})
	r.Route("/auth", func(r chi.Router) {
		r.Get("/user/{userID}", h.getUser)
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://"+r.Host+"/mvp/", http.StatusMovedPermanently)
	})

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
