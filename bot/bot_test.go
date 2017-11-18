package bot

import (
	"fmt"
	"testing"
)

func loadfiles() {
	LoadRecipes()
}

func TestRecipeLoadFile(t *testing.T) {
	recipe := LoadRecipeFile("recipes/test_basic.json")
	if recipe.Name != "test_basic" {
		fmt.Println("Can't read files")
		t.Fatal()
	}
}

func TestRecipeLoad(t *testing.T) {
	LoadRecipes()
	if len(recipeList) != 3 {
		fmt.Println("Can't load recipe files")
		t.Fatal()
	}
}

func TestFindRecipe(t *testing.T) {
	loadfiles()
	recipe := FindRecipe("test_basic")
	if recipe.Name != "test_basic" {
		fmt.Println("Can't find the recipes")
		t.Fatal()
	}

	recipe2 := FindRecipe("recipe_failing")
	if recipe2.Name != "" {
		fmt.Println("Found a recipe when it shouldn't")
		t.Fatal()
	}
}

func TestRecipeRouteBasic(t *testing.T) {
	loadfiles()
	CreateRoutes()
	recipe := FindRecipe("test_basic")
	data := RouterData{}
	retdata := Route(recipe, data)
	if retdata.Message != "Hello World" {
		fmt.Println("Could not run route")
		t.Fatal()
	}
}

func TestRecipeRouteAdvance(t *testing.T) {
	loadfiles()
	CreateRoutes()
	recipe := FindRecipe("test_advance")
	data := RouterData{
		Data: map[string]string{
			"username": "Test",
		},
	}
	retdata := Route(recipe, data)
	if retdata.Message != "Give me something to work with Test" {
		fmt.Println("Could not run route")
		t.Fatal()
	}
}

func TestRecipeRouteMulti(t *testing.T) {
	loadfiles()
	CreateRoutes()

	recipe := FindRecipe("test_multi")
	data := RouterData{
		Data: map[string]string{
			"username":      "Test",
			"amount":        "200",
			"currency-from": "NOK",
			"currency-to":   "EUR",
			"rate":          "200",
		},
	}
	retdata := Route(recipe, data)
	if retdata.Message != "200 NOK to EUR is 200" {
		fmt.Println("Could not run route")
		t.Fatal()
	}
}
