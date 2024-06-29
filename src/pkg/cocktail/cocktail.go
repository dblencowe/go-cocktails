package cocktail

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
)

var ErrNoRecipes = errors.New("no recipes loaded")

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Cocktail struct {
	Name         string       `json:"name"`
	Instructions string       `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
}

var cocktails []Cocktail

func Random() Cocktail {
	return cocktails[rand.Intn(len(cocktails))]
}

func LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &cocktails); err != nil {
		return err
	}

	if len(cocktails) == 0 {
		return ErrNoRecipes
	}

	return nil
}
