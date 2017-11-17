package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

var recipeList map[string]Recipe

func LoadRecipeFile(path string) Recipe {

	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}
	var jsontype Recipe
	jsonError := json.Unmarshal(file, &jsontype)
	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
	}
	return jsontype
}
func LoadRecipes() {
	recipeList = make(map[string]Recipe)
	files, err := filepath.Glob("./recipes/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		recipe := LoadRecipeFile(file)
		recipeList[recipe.Name] = recipe
	}

}

func FindRecipe(trigger string) Recipe {
	for key := range recipeList {
		triggers := recipeList[key].Triggers
		for _, t := range triggers {
			if t == trigger {
				return recipeList[key]
			}
		}
	}
	return Recipe{}
	// return nil, ErrorInterface{Message: "Not found"}
}
