package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Command represents a write operation
type Command interface {
	Execute()
}

// Query represents a read operation
type Query interface {
	Execute() string
}

// CreateUserCommand is an example of a write operation
type CreateUserCommand struct {
	UserID   int
	Username string
	Email    string
}

// Execute creates a new user
func (c *CreateUserCommand) Execute() {
	// Logic to create a user (write operation)
	fmt.Printf("User %s created with ID %d\n", c.Username, c.UserID)
}

// GetUserQuery is an example of a read operation
type GetUserQuery struct {
	UserID int
}

// Execute retrieves user details
func (q *GetUserQuery) Execute() string {
	// Logic to retrieve user details (read operation)
	return fmt.Sprintf("User details for ID %d", q.UserID)
}

// CommandHandler handles commands
type CommandHandler interface {
	Handle(Command)
}

// QueryHandler handles queries
type QueryHandler interface {
	Handle(Query) string
}

// UserCommandHandler handles CreateUserCommand
type UserCommandHandler struct{}

// Handle executes the CreateUserCommand
func (h *UserCommandHandler) Handle(cmd Command) {
	if createCmd, ok := cmd.(*CreateUserCommand); ok {
		createCmd.Execute()
	} else {
		fmt.Println("Invalid command type")
	}
}

// UserQueryHandler handles GetUserQuery
type UserQueryHandler struct{}

// Handle executes the GetUserQuery
func (h *UserQueryHandler) Handle(query Query) string {
	if getUserQuery, ok := query.(*GetUserQuery); ok {
		return getUserQuery.Execute()
	}
	return "Invalid query type"
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create CreateUserCommand
	// For simplicity, assume the parameters are passed in the URL query string
	userID := 1 // replace with actual parsing logic
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	createUserCommand := &CreateUserCommand{UserID: userID, Username: username, Email: email}

	// Handle the command
	commandHandler := &UserCommandHandler{}
	commandHandler.Handle(createUserCommand)

	// Send a response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s created with ID %d\n", username, userID)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request and create GetUserQuery
	// For simplicity, assume the user ID is passed in the URL path
	userID := 1 // replace with actual parsing logic
	getUserQuery := &GetUserQuery{UserID: userID}

	// Handle the query
	queryHandler := &UserQueryHandler{}
	result := queryHandler.Handle(getUserQuery)

	// Send a response
	fmt.Fprint(w, result)
}

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/user", createUserHandler).Methods("POST")
	router.HandleFunc("/user/{userID}", getUserHandler).Methods("GET")

	// Start the HTTP server
	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
