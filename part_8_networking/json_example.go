package part8

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

// Requests define requests in json
type Requests struct {
	Requests []Request `json:"requests"`
}

// Request define single request in json
type Request struct {
	User   string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `amount:"amount"`
}

// ReadRequestsFromFile read json to struct from file
func ReadRequestsFromFile() Requests {
	//open json file
	jsonFile, e := os.Open("bank.json")
	if e != nil {
		return Requests{}
	}
	defer jsonFile.Close()

	bankStream := bufio.NewReader(jsonFile)
	return decodeToRequest(bankStream)
}

// ReadRequestsFromString read json to struct from string
func ReadRequestsFromString() Requests {
	var requestStr = `
	"requests": [
		{"user" : "alex", "type": "worker", "amount": 44.2},
		{"user" : "tom", "type": "manager", "amount": 32.2},
		{"user" : "veronica", "type": "designer", "amount": 45.2}
	]
	`

	bankStream := strings.NewReader(requestStr)
	return decodeToRequest(bankStream)
}

func decodeToRequest(r io.Reader) Requests {
	bankDecoder := json.NewDecoder(r)

	var requests Requests
	if e := bankDecoder.Decode(&requests); e != nil {
		log.Fatalf("Error: %e", e)
	}

	return requests
}
