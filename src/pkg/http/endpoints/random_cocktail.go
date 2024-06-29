package endpoints

import (
	"encoding/json"
	"github.com/dblencowe/go-cocktails/src/pkg/cocktail"
	"net/http"
)

func HandleRandomCocktail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cocktail.Random())
}
