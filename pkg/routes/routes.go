package routes

import (
	_ "embed"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/servusdei2018/shawty/pkg/db"
)

//go:embed templates/index
var index string
var indexTempl *template.Template

func init() {
	indexTempl = template.Must(template.New("index").Parse(index))
}

// Router contains router data.
type Router struct {
	DB  *db.DB
	Mux *http.ServeMux
	key string
}

// New creates a new Router with provided db and key.
func New(db *db.DB, key string) *Router {
	var r Router
	mux := http.NewServeMux()

	mux.HandleFunc("/delete", r.deleteHandler)
	mux.HandleFunc("/shorten", r.shortenHandler)
	mux.HandleFunc("/", r.catchAllHandler)

	r.DB = db
	r.key = key
	r.Mux = mux

	return &r
}

// deleteHandler handles the /delete endpoint.
//
// curl -H "Auth: YourAuthToken" -X POST -d "shortID" http://localhost/delete
func (rtr *Router) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	if r.Header.Get("Auth") != rtr.key {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	rtr.DB.Delete(string(body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(http.StatusText(http.StatusNoContent)))
	}
}

// shortenHandler handles the /shorten endpoint.
//
// curl -H "Auth: YourAuthToken" -X POST -d "https://random.site/a-very-long-url" http://localhost/shorten
func (rtr *Router) shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	if r.Header.Get("Auth") != rtr.key {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	url, err := rtr.DB.Store(string(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(url))
	}
}

// catchAllHandler handles everything else.
func (rtr *Router) catchAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	shortID := r.URL.Path
	shortID = strings.TrimPrefix(shortID, "/")

	// Serve index.
	if shortID == "" {
		w.WriteHeader(http.StatusOK)
		indexTempl.Execute(w, rtr.DB.Stats())
		return
	}

	// Serve shortened URL.
	url, err := rtr.DB.Get(shortID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
