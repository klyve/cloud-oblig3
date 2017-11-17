package bot

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

// ReturnStruct struct
type ReturnStruct struct {
	MessagingType string        `json:"messaging_type"`
	Recipient     Recipient     `json:"recipient"`
	Message       ReturnMessage `json:"message"`
}

// ReturnMessage struct
type ReturnMessage struct {
	Text string `json:text`
}
