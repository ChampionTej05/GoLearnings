package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp // to help route
}

//ServeHTTP writes the data to w and returns the response
// method is tied to the controller
// ServeHTTP is from Handle Interface so by default UC will implement that interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Rakshit"))
}

// go doesn't have constructor to initialiaze the userIDPattern
// instead we use constructor func to do the work for us
// constructor function name start with "new" followed by object
func newUserController() *userController {
	return &userController{userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`)}
}

type sampleController struct {
	sampleIDPattern *regexp.Regexp // to help route
}

func (sc sampleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Sample"))
}

func newSampleController() *sampleController {
	return &sampleController{sampleIDPattern: regexp.MustCompile(`^/sample/(\d+)/?`)}
}
