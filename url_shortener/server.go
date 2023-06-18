package main

import (
	"fmt"
	url2 "github.com/PyMarcus/url_shortener/url"
	"log"
	"net/http"
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
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	var status int

	if r.Method != "POST" {
		response(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}

	url, new, err := url2.SearchOrCreateAnewUrl(extractUrl(r))

	if err != nil {
		response(w, http.StatusBadRequest, nil)
		return
	}

	if new {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	response(w, status, Headers{
		"Location": shortUrl,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request) {

}

// define as rotas
func createRoutes() {
	// encurta url
	http.HandleFunc("/api/shorten", Shortener)
	// redireciona a url encurtada para a original
	http.HandleFunc("/r/<id>", Redirect)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port), nil))
}

func main() {
	createRoutes()
}
