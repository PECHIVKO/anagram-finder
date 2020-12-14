package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PECHIVKO/anagram-finder/service"
)

type Handle func(http.ResponseWriter, *http.Request)
type Router struct {
	mux  map[string]Handle
	dict *service.Dictionary
}

// NewRouter creates new router
func NewRouter(d *service.Dictionary) *Router {
	var router Router = Router{
		mux:  make(map[string]Handle),
		dict: d, // handle adress of Dictionary from main.go
	}
	router.Add("/", router.homePage)
	router.Add("/get", router.returnAnagrams)
	router.Add("/load", router.loadJSON)
	return &router
}

// Add  adds new handler
func (r *Router) Add(path string, handle Handle) {
	r.mux[path] = handle
}

// GetHeader gets adress after "/"
func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

// ServeHTTP implements http.Handler method
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := GetHeader(r.URL.Path)
	h, ok := rt.mux[head]
	if ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}
