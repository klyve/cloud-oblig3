package bot

type Recipe struct {
	Name     string   `json:"name"`
	Messages []string `json:"messages"`
	Triggers []string `json:"triggers"`
	Route    []string `json:"route"`
}

type Recipes struct {
	RecipeList map[string]Recipe
}

type ErrorInterface struct {
	Message string
}

type RouterData struct {
	Message string
	Count   int
}
