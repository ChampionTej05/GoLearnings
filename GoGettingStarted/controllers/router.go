package controllers

import "net/http"

func RegisterController() {
	uc := newUserController()
	sc := newSampleController()
	// uc implements Handle interface because we have serveHTTP defined for UC
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	http.Handle("/sample", *sc)
	http.Handle("/sample/", *sc)
}
