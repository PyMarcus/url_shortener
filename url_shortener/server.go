package main

import (
	"fmt"
	"github.com/PyMarcus/url_shortener/url"
	"log"
	"net/http"
	"strings"
)

type Headers map[string]string

var (
	port    int
	urlBase string
)

// função inicializadora
func init() {
	port = 8000
	urlBase = fmt.Sprintf("http://localhost:%d", port)
}

// define o cabecalho da resposta e seu status
func response(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

// extrai url do corpo da requisicao
func extractUrl(r *http.Request) string {
	urll := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(urll)
	return string(urll)
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	var status int

	if r.Method != "POST" {
		response(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}

	urll, neww, err := url.SearchOrCreateAnewUrl(extractUrl(r))
	fmt.Println("Connection received from ", r.RemoteAddr)
	fmt.Println("Order to shorten ", urll.Dest)
	if err != nil {
		response(w, http.StatusBadRequest, nil)
		return
	}

	if neww {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", urlBase, urll.Id)
	fmt.Println("Response: ", shortUrl, " Status: ", status)
	response(w, status, Headers{
		"Location": shortUrl,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		http.Redirect(w, r, url.Dest, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}

// define as rotas
func createRoutes() {
	url.SettingRepository(url.NewMemoryRepository())
	// encurta url
	http.HandleFunc("/api/shorten", Shortener)
	// redireciona a url encurtada para a original
	http.HandleFunc("/r/", Redirect)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port), nil))
}

func main() {
	fmt.Println("Server is running in port: ", port)
	createRoutes()
}
