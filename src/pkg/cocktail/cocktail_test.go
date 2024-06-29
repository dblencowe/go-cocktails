package cocktail

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCocktailFunctionality(t *testing.T) {
	t.Run("can load cocktails", func(t *testing.T) {
		recipeData := []byte(`[{"name":"Daiquiri","instructions":"Combine rum, lime juice, and simple syrup in a shaker with ice. Shake well and strain into a chilled coupe glass.","ingredients":[{"name":"Rum","quantity":"2 oz"},{"name":"Lime juice","quantity":"1 oz"},{"name":"Simple syrup","quantity":"0.5 oz"}]}]`)
		err := os.WriteFile("./test_recipes.json", recipeData, os.ModePerm)
		defer os.Remove("./test_recipes.json")
		assert.NoError(t, err)

		err = LoadFromFile("./test_recipes.json")
		assert.NoError(t, err)
		assert.Len(t, cocktails, 1)
	})

	t.Run("can produce a random recipe", func(t *testing.T) {
		cocktails = []Cocktail{
			{
				Name:         "Daiquiri",
				Instructions: "Muddle mint and lime juice in a highball glass. Add rum, top with soda water, and stir. Garnish with a mint sprig.",
				Ingredients: []Ingredient{
					{
						Name:     "Rum",
						Quantity: "2 oz",
					},
					{
						Name:     "Lime Juice",
						Quantity: "1 oz",
					},
					{
						Name:     "Simple Syrup",
						Quantity: "0.5 oz",
					},
				},
			},
		}

		assert.Len(t, cocktails, 1)

		recipe := Random()
		assert.Equal(t, "Daiquiri", recipe.Name)
		assert.Len(t, recipe.Ingredients, 3)
	})
}
