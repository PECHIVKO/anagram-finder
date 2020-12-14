package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PECHIVKO/anagram-finder/models"
)

type Router struct {
	mux  map[string]Handle
	dict *models.Dictionary
}

func NewRouter(d *models.Dictionary) *Router {
	var router Router = Router{
		mux:  make(map[string]Handle),
		dict: d,
	}
	router.Add("/", router.homePage)
	router.Add("/get", router.returnAnagrams)
	router.Add("/load", router.loadJSON)
	return &router
}

func (r *Router) Add(path string, handle Handle) {
	r.mux[path] = handle
}

func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := GetHeader(r.URL.Path)
	h, ok := rt.mux[head]
	if ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}
