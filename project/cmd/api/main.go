package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
)

func main() {
	apiUrl := os.Getenv("API_URL")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(apiUrl)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer resp.Body.Close()

		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		fmt.Fprint(w, fmt.Sprintf("Response from API: %s", string(html)))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("http server listening at %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
