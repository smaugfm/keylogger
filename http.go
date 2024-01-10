package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var baseUrl = "http://localhost:8080"
var keyPressesChannel = make(chan KeyPress, 1000)

type KeyPress struct {
	key     string
	keyCode int
	caps    bool
	shift   bool
	option  bool
	cmd     bool
	control bool
}

type Event struct {
	Origin    string   `json:"origin"`
	Timestamp string   `json:"timestamp"`
	Type      string   `json:"type"`
	Data      KeyPress `json:"data"`
}

func keyPressesLoop() {
	fmt.Printf("\n\nListening for keys...")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for keyPress := range keyPressesChannel {
		j, err := json.Marshal(
			Event{
				Timestamp: time.Now().Format(time.RFC3339),
				Type:      "desktop-keyboard-keypress",
				Origin:    hostname,
				Data:      keyPress,
			},
		)
		if err != nil {
			fmt.Printf("Unable to marshal keypress to json: %v\n", err)
			continue
		}
		_, err = http.Post(baseUrl+"/events", "application/json", bytes.NewReader(j))
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
