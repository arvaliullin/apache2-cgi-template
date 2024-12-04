package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// @title My API
// @version 1.0.0
// @description Description of my API
// @host example.com
// @basePath /api/v1

// Response defines the structure of the response object.
type Response struct {
	Message string `json:"message"`
}

// @Summary Get Bye Message
// @Description Returns a JSON object containing a bye message
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /cgi-bin/bye.cgi [get]
func byeHandler() {
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

func main() {
	helloHandler()
}
