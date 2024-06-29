package endpoints_test

import (
	"github.com/dblencowe/go-cocktails/src/pkg/cocktail"
	"github.com/dblencowe/go-cocktails/src/pkg/http/endpoints"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRandomCocktail(t *testing.T) {
	recipe := `{"name":"Daiquiri","instructions":"Combine rum, lime juice, and simple syrup in a shaker with ice. Shake well and strain into a chilled coupe glass.","ingredients":[{"name":"Rum","quantity":"2 oz"},{"name":"Lime juice","quantity":"1 oz"},{"name":"Simple syrup","quantity":"0.5 oz"}]}`
	recipeData := []byte("[" + recipe + "]")
	err := os.WriteFile("./test_recipes.json", recipeData, os.ModePerm)
	defer os.Remove("./test_recipes.json")
	assert.NoError(t, err)

	err = cocktail.LoadFromFile("./test_recipes.json")
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/random", nil)
	w := httptest.NewRecorder()
	endpoints.HandleRandomCocktail(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, recipe+"\n", string(data))
}
