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

//   "object": "page",
//   "entry": [
//     {
//       "id": "146560019406297",
//       "time": 1510878676820,
//       "messaging": [
//         {
//           "sender": {
//             "id": "1536356476448699"
//           },
//           "recipient": {
//             "id": "146560019406297"
//           },
//           "timestamp": 1510878675989,
//           "message": {
//             "mid": "mid.$cAADfuGqZvYpl-KhMFVfx2MDyvN3F",
//             "seq": 1424515,
//             "text": "hello"
//           }
//         }
//       ]
//     }
//   ]
// }
