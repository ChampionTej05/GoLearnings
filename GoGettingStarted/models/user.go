package models

import (
	"errors"
	"fmt"
)

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
func GetUsers() ([]*User, error) {
	return users, nil
}

// addUser adds the given user to the list of users
// returns newly added user and error if any with empty struct
func AddUser(user User) (User, error) {
	// default value of int = 0
	if user.ID != 0 {
		return User{}, errors.New("User already has the ID defined, not allowed")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user) // need pointer of values
	return user, nil
}

// getUserByID returns the user with the given ID
// returns error if any with empty struct
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' is not found in Database", id)
}

// updateUserByID updates the user with the given ID
// returns error if any with empty struct
func UpdateUserByID(u User) (User, error) {
	for idx, user := range users {
		if user.ID == u.ID {
			users[idx] = &u
			return u, nil

		}
	}
	return User{}, fmt.Errorf("User with ID '%v' is not found in Database", u.ID)
}

//removeUserByID removes the user with the given ID
// returns error if any
func RemoveUserByID(id int) error {
	for idx, user := range users {
		if user.ID == id {
			// standard syantax for removing an element in slice
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' is not found in Database", id)
}
