package part8

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// MathBinaryOperations describe binary math operation: +, -, *, /
type MathBinaryOperations struct {
	Operation string  `json:"operation"`
	Left      float64 `json:"left"`
	Right     float64 `json:"right"`
}

// MathResponse response to mathematical operation
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := MathBinaryOperations{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)
	var result float64 = 0
	switch req.Operation {
	case "-":
		result = req.Left - req.Right
		break
	case "+":
		result = req.Left + req.Right
		break
	case "*":
		result = req.Left * req.Right
		break
	case "/":
		result = req.Left / req.Right
		break
	default:
		result = 0.0
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(MathResponse{"nil", result})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	s := "welcome to math server"
	w.Write([]byte(s))
}

// RunServer run http web server
func RunServer() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	if e := http.ListenAndServe(addr, nil); e != nil {
		log.Fatal(e)
	}
}

// GetHomePage get homepage string
func GetHomePage() string {
	r, e := http.Get("http://localhost:8080")
	if e != nil {
		log.Fatal(e)
	}
	defer r.Body.Close()

	body, e := io.ReadAll(r.Body)
	if e != nil {
		log.Fatal(e)
	}

	return string(body)
}
