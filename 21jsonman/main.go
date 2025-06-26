package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct for encoding/decoding
type User struct {
	Name     string   `json:"name"` // Struct tag for JSON key
	Email    string   `json:"email"`
	Age      int      `json:"age"`
	Skills   []string `json:"skills,omitempty"` // omit if empty
	IsActive bool     `json:"is_active"`
}

func main() {
	fmt.Println("Welcome to JSON Manipulation in Go!")

	// ----------------------------
	// Example 1: Encoding (struct → JSON)
	fmt.Println("\nExample 1: Encode struct to JSON")
	user := User{
		Name:     "Debarshee Chakraborty",
		Email:    "debarshee@example.com",
		Age:      21,
		Skills:   []string{"Go", "Python", "ML"},
		IsActive: true,
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	} else {
		fmt.Println("JSON (compact):", string(jsonBytes))
	}

	// Pretty-print JSON
	jsonPretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("JSON (pretty):\n", string(jsonPretty))

	// ----------------------------
	// Example 2: Decoding (JSON → struct)
	fmt.Println("\nExample 2: Decode JSON to struct")
	jsonInput := `{"name":"Saptak Biswas","email":"saptak@example.com","age":22,"skills":["Cloud","DevOps"],"is_active":false}`

	var decodedUser User
	err = json.Unmarshal([]byte(jsonInput), &decodedUser)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	} else {
		fmt.Printf("Decoded User: %+v\n", decodedUser)
	}

	// ----------------------------
	// Example 3: Working with dynamic JSON (map[string]interface{})
	fmt.Println("\nExample 3: Decode to dynamic map")
	dynamicJSON := `{"project": "Build With AI", "participants": 120, "online": true}`

	var data map[string]interface{}
	err = json.Unmarshal([]byte(dynamicJSON), &data)
	if err != nil {
		fmt.Println("Error decoding dynamic JSON:", err)
	} else {
		fmt.Println("Decoded map:")
		for k, v := range data {
			fmt.Printf("Key: %s, Value: %v, Type: %T\n", k, v, v)
		}
	}

	// ----------------------------
	// Example 4: Encoding a map to JSON
	fmt.Println("\nExample 4: Encode map to JSON")
	config := map[string]interface{}{
		"theme":    "dark",
		"version":  2.1,
		"features": []string{"json", "file", "http"},
	}
	encodedMap, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println("Map as JSON:\n", string(encodedMap))
}
