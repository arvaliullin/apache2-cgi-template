package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	resp := Response{Message: "Hello, World!"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		return
	}
	
	fmt.Println("Content-Type: application/json")
	fmt.Println()
	fmt.Println(string(jsonResp))
}
