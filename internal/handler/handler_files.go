package handler


// User represents a user structure.
type User struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
    // Add other fields as needed
}

// Users represents a list of users.
type Users struct {
    Data []User `json:"data"`
}

