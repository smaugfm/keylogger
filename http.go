package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var url = "http://localhost:8000"
var keysChannel = make(chan KeyPress, 1000)

type KeyPress struct {
	key     string
	keyCode int
	caps    bool
	shift   bool
	option  bool
	cmd     bool
	control bool
}

func keyPressesLoop() {
	fmt.Printf("\n\nListening for keys...")
	for keyPress := range keysChannel {
		j, err := json.Marshal(keyPress)
		if err != nil {
			fmt.Printf("Unable to marshal keypress to json: %v\n", err)
			continue
		}
		_, err = http.Post(url, "application/json", bytes.NewReader(j))
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
