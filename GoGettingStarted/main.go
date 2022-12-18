package main

import (
	"net/http"

	"github.com/ChampionTej05/GoLearnings/GoGettingStarted/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
