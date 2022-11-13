package part8

import (
	"bufio"
	"encoding/json"
	"fmt"
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
func ReadRequestsFromFile(path string) Requests {
	//open json file
	jsonFile, e := os.Open(path)
	if e != nil {
		return Requests{}
	}
	defer jsonFile.Close()

	bankStream := bufio.NewReader(jsonFile)
	return decodeToRequest(bankStream)
}

// ReadRequestsFromString read json to struct from string
func ReadRequestsFromString(requestsAsJSONString string) Requests {
	bankStream := strings.NewReader(requestsAsJSONString)
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

// WriteRequestsToFile write requests to json file
func WriteRequestsToFile(path string, requests Requests) {
	if e := os.Remove(path); e != nil {
		fmt.Print("File removed")
	}

	requestsFile, e := os.Create(path)
	if e != nil {
		log.Fatalf("Can not create file. Error %e", e)
	}
	defer requestsFile.Close()

	requestsEncoder := json.NewEncoder(requestsFile)

	if e := requestsEncoder.Encode(requests); e != nil {
		log.Fatalf("Can not encode struct. Error %e", e)
	}
}
