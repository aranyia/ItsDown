package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type WebHookMessage struct {
	Text    string `json:"text"`
	Channel string `json:"channel,omitempty"`
}

type WebHook struct {
	URL string
}

func (WebHook WebHook) PostMessage(text string) {
	log.Println("sending message to web hook URL:", WebHook.URL)

	msg := WebHookMessage{Text: text}
	jsonContent, _ := json.Marshal(msg)
	resp, err := http.Post(WebHook.URL, http.DetectContentType(jsonContent), bytes.NewBuffer(jsonContent))

	log.Print("web hook status: ", resp.Status)
	if err != nil {
		log.Panic(err)
	}
}
