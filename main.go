package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type Message struct {
	Title string                 `json:"title"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	data := map[string]interface{}{
		"email": "user1@gmail.com",
	}
	message := Message{
		Title: "UserDeleted",
		Type:  "event",
		Data:  data,
	}
	serializedData, _ := json.Marshal(message)
	err := nc.Publish("events", serializedData)
	if err != nil {
		return
	}
}
