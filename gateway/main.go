package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

func NewReverseProxy(target string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(targetURL), nil
}

func main() {
	r := mux.NewRouter()

	user_service_url := os.Getenv("USER_SERVICE_URL")
	notes_service_url := os.Getenv("NOTES_SERVICE_URL")
	gateway_port := os.Getenv("GATEWAY_PORT")

	user_proxy, err := NewReverseProxy(user_service_url)
	if err != nil {
		log.Fatal(err)
	}

	notes_proxy, err := NewReverseProxy(notes_service_url)
	if err != nil {
		log.Fatal(err)
	}

	r.HandleFunc("/user/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
		user_proxy.ServeHTTP(w, r)
	})

	r.HandleFunc("/notes/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
		notes_proxy.ServeHTTP(w, r)
	})

	log.Println("API Gateway started on :" + gateway_port)
	log.Fatal(http.ListenAndServe(":"+gateway_port, r))
}
