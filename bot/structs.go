package bot

// Recipe contains data from single recipies
type Recipe struct {
	Name     string   `json:"name"`
	Messages []string `json:"messages"`
	Triggers []string `json:"triggers"`
	Route    []string `json:"route"`
}

// Recipes contains data from the recipies
type Recipes struct {
	RecipeList map[string]Recipe
}

// ErrorInterface contains data about the errors
type ErrorInterface struct {
	Message string
}

// RouterData contains data about the routers
type RouterData struct {
	Message string
	Count   int
	Data    map[string]string
	Error   bool
	ErrorTo string
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
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

// Message struct
type Message struct {
	Mid  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}

// Sender struct
type Sender struct {
	ID string `json:"id"`
}

// Recipient struct
type Recipient struct {
	ID string `json:"id"`
}

// FBReturnStruct struct
type FBReturnStruct struct {
	MessagingType string          `json:"messaging_type"`
	Recipient     Sender          `json:"recipient"`
	Message       FBReturnMessage `json:"message"`
}

// FBReturnMessage struct
type FBReturnMessage struct {
	Text string `json:"text"`
}

// DialogFlowQuery struct
type DialogFlowQuery struct {
	Language  string `json:"lang"`
	Message   string `json:"query"`
	SessionID string `json:"sessionId"`
}

// DialogFlowResponse struct
type DialogFlowResponse struct {
	Result    DialogFlowResult `json:"result"`
	SessionID string           `json:"sessionId"`
}

// DialogFlowResult struct
type DialogFlowResult struct {
	Parameters map[string]string  `json:"parameters"`
	Metadata   DialogFlowMetadata `json:"metadata"`
	Score      float64            `json:"score"`
}

// DialogFlowMetadata struct contains data from DialogFlow's metadata
type DialogFlowMetadata struct {
	IntentName string `json:"intentName"`
}

// FBUser struct contains data about facebook user
type FBUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
