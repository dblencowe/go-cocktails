package http

import (
	"github.com/dblencowe/go-cocktails/src/pkg/http/endpoints"
	"log"
	"net/http"
)

func Serve(bindAddr string) error {
	http.HandleFunc("/random", endpoints.HandleRandomCocktail)
	log.Printf("server listening on %s", bindAddr)
	return http.ListenAndServe(bindAddr, nil)
}
