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

// getUsers returns list of users and errors if any
func getUsers() ([]*User, error) {
	return users, nil
}

// addUser adds the given user to the list of users
// returns newly added user and error if any
func addUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user) // need pointer of values
	return user, nil
}
