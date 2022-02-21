package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User // slice of pointers to User object
	// this helps in maintaining/manipulating objects throughout packages

	nextID = 1 // kind of primary key to keep ID from DB
)
