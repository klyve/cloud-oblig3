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

// FBWebHook struct for data from Facebook
type FBWebHook struct {
	Object string    `json:"object"`
	Entry  []Entries `json:"entry"`
}

// Entries struct for entry data from Facebook
type Entries struct {
	ID        string      `json:"id"`
	Time      int         `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

// Messaging struct forn data from facebook
type Messaging struct {
	Sender    map[string]string `json:"sender"`
	Recipient map[string]string `json:"recipient"`
	Timestamp int64             `json:"timestamp"`
	Message   Message           `json:"message"`
}

// Message struct
type Message struct {
	Mid  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}
