package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to HTTP Requests in Go using net/http")

	// Replace this with a valid endpoint for actual testing (like https://httpbin.org/)
	apiURL := "https://httpbin.org"

	// ----------------------------
	// Example 1: GET Request
	fmt.Println("\nExample 1: GET Request")
	resp, err := http.Get(apiURL + "/get")
	if err != nil {
		fmt.Println("GET request error:", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("GET response:", string(body))
	}

	// ----------------------------
	// Example 2: POST Request with JSON
	fmt.Println("\nExample 2: POST Request")
	postData := map[string]string{"name": "Debarshee", "language": "Go"}
	postJSON, _ := json.Marshal(postData)

	resp, err = http.Post(apiURL+"/post", "application/json", bytes.NewBuffer(postJSON))
	if err != nil {
		fmt.Println("POST request error:", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("POST response:", string(body))
	}

	// ----------------------------
	// Examples 3-5: PUT, PATCH, DELETE using http.NewRequest()
	client := &http.Client{}

	// ----------------------------
	// Example 3: PUT Request
	fmt.Println("\nExample 3: PUT Request")
	putData := map[string]string{"update": "true"}
	putJSON, _ := json.Marshal(putData)

	putReq, _ := http.NewRequest(http.MethodPut, apiURL+"/put", bytes.NewBuffer(putJSON))
	putReq.Header.Set("Content-Type", "application/json")

	putResp, err := client.Do(putReq)
	if err != nil {
		fmt.Println("PUT request error:", err)
	} else {
		defer putResp.Body.Close()
		body, _ := io.ReadAll(putResp.Body)
		fmt.Println("PUT response:", string(body))
	}

	// ----------------------------
	// Example 4: PATCH Request
	fmt.Println("\nExample 4: PATCH Request")
	patchData := map[string]string{"field": "patched"}
	patchJSON, _ := json.Marshal(patchData)

	patchReq, _ := http.NewRequest(http.MethodPatch, apiURL+"/patch", bytes.NewBuffer(patchJSON))
	patchReq.Header.Set("Content-Type", "application/json")

	patchResp, err := client.Do(patchReq)
	if err != nil {
		fmt.Println("PATCH request error:", err)
	} else {
		defer patchResp.Body.Close()
		body, _ := io.ReadAll(patchResp.Body)
		fmt.Println("PATCH response:", string(body))
	}

	// ----------------------------
	// Example 5: DELETE Request
	fmt.Println("\nExample 5: DELETE Request")
	deleteReq, _ := http.NewRequest(http.MethodDelete, apiURL+"/delete", nil)

	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		fmt.Println("DELETE request error:", err)
	} else {
		defer deleteResp.Body.Close()
		body, _ := io.ReadAll(deleteResp.Body)
		fmt.Println("DELETE response:", string(body))
	}
}
