package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Course and Author - course.go and author.go
// Course represents a training course with details and author information
type Course struct {
	ID       string  `json:"id"`       // Unique identifier for the course
	Name     string  `json:"name"`     // Course title/name
	Duration string  `json:"duration"` // Course duration (e.g., "3h", "5h")
	Price    float64 `json:"price"`    // Course price in dollars
	Author   *Author `json:"author"`   // Pointer to author information
}

// Author represents the course instructor/creator
type Author struct {
	ID       string `json:"id"`       // Unique identifier for the author
	Fullname string `json:"fullname"` // Author's full name
	Email    string `json:"email"`    // Author's email address
}

// Fake DB - In-memory slice to simulate database storage
var courses = []Course{
	{
		ID:       "1",
		Name:     "Go Basics",
		Duration: "3h",
		Price:    29.99,
		Author: &Author{
			ID:       "a1",
			Fullname: "John Doe",
			Email:    "john@example.com",
		},
	},
	{
		ID:       "2",
		Name:     "Advanced Go",
		Duration: "5h",
		Price:    49.99,
		Author: &Author{
			ID:       "a2",
			Fullname: "Jane Smith",
			Email:    "jane@example.com",
		},
	},
}

// Middleware for empty fields - middleware.go
// IsEmpty checks if a course has essential data (name is required)
func (c *Course) IsEmpty() bool {
	// Check if the course name is empty (primary validation)
	return c.Name == ""
}

// Controller for Course - course_controller.go

// serveHome handles the root endpoint and displays welcome message
// GET /
func serveHome(w http.ResponseWriter, r *http.Request) {
	// Send HTML welcome message to client
	w.Write([]byte("<h1>Welcome to Course API</h1>"))
}

// getAllCourse handles retrieving all courses from the database
// GET /courses
func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a GET All route")

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the entire courses slice to JSON and send response
	json.NewEncoder(w).Encode(courses)
}

// getOneCourse handles retrieving a single course by ID
// GET /courses/{id}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a GET One route")

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract URL parameters using Gorilla Mux
	params := mux.Vars(r)
	courseID := params["id"] // Get the course ID from URL path

	// Search through courses slice to find matching ID
	for _, course := range courses {
		if course.ID == courseID {
			// Course found - return it as JSON
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// No course found with the given ID
	json.NewEncoder(w).Encode("No course found with the given ID")
}

// createOneCourse handles creating a new course
// POST /courses
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a create one route")

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Validate request body exists
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
		return
	}

	// Parse JSON request body into Course struct
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// Validate that course has required data
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in the request body")
		return
	}

	// Generate unique ID for the new course
	// Note: In production, use UUID or database auto-increment
	course.ID = strconv.Itoa(rand.Intn(100))

	// Add the new course to our in-memory database
	courses = append(courses, course)

	// Return the created course with generated ID
	json.NewEncoder(w).Encode(course)
}

// updateOneCourse handles updating an existing course
// PUT /courses/{id}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is an update one route")

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract course ID from URL parameters
	params := mux.Vars(r)
	courseID := params["id"]

	// Validate request body exists
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
		return
	}

	// Parse JSON request body into Course struct
	var updatedCourse Course
	_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

	// Validate that updated course has required data
	if updatedCourse.IsEmpty() {
		json.NewEncoder(w).Encode("No valid data in the request body")
		return
	}

	// Find and update the course with matching ID
	for index, course := range courses {
		if course.ID == courseID {
			// Preserve the original ID and update other fields
			updatedCourse.ID = courseID
			courses[index] = updatedCourse

			// Return the updated course
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	// No course found with the given ID
	json.NewEncoder(w).Encode("No course found with the given ID")
}

// deleteOneCourse handles deleting a course by ID
// DELETE /courses/{id}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a delete one route")

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Extract course ID from URL parameters
	params := mux.Vars(r)
	courseID := params["id"]

	// Find and remove the course with matching ID
	for index, course := range courses {
		if course.ID == courseID {
			// Remove course from slice using slice operations
			// courses[:index] gets elements before the target
			// courses[index+1:]... gets elements after the target
			courses = append(courses[:index], courses[index+1:]...)

			// Confirm deletion with success message
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	// No course found with the given ID
	json.NewEncoder(w).Encode("No course found with the given ID")
}

func main() {
	fmt.Println("Course API Server Starting...")

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define API routes with their corresponding handlers

	// Home/Welcome route
	r.HandleFunc("/", serveHome).Methods("GET")

	// CRUD operations for courses
	r.HandleFunc("/courses", getAllCourse).Methods("GET")            // Read all courses
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")       // Read one course
	r.HandleFunc("/courses", createOneCourse).Methods("POST")        // Create new course
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")    // Update existing course
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE") // Delete course

	// Start the HTTP server on port 4000
	fmt.Println("Server is listening on port 4000...")
	http.ListenAndServe(":4000", r)
}
