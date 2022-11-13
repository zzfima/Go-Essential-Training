package part8

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
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

// ReadJSON read json to struct
func ReadJSON() Requests {
	//open json file
	jsonFile, e := os.Open("bank.json")
	if e != nil {
		return Requests{}
	}
	defer jsonFile.Close()

	bankStream := bufio.NewReader(jsonFile)
	bankDecoder := json.NewDecoder(bankStream)

	var requests Requests
	if e := bankDecoder.Decode(&requests); e != nil {
		log.Fatalf("Error: %e", e)
	}

	return requests
}
