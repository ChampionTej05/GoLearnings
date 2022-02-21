package main

import (
	"fmt"

	"github.com/ChampionTej05/GoLearnings/GoGettingStarted/models"
)

func main() {
	user := models.User{
		ID:        1,
		FirstName: "Rakshit",
		LastName:  "Kathawate",
	}
	fmt.Println(user)
	models.User{}
}
