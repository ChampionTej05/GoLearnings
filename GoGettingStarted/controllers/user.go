package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ChampionTej05/GoLearnings/GoGettingStarted/models"
)

type userController struct {
	userIDPattern *regexp.Regexp // to help route
}

//ServeHTTP writes the data to w and returns the response
// method is tied to the controller
// ServeHTTP is from Handle Interface so by default UC will implement that interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	fmt.Println("ServeHTTP : urlPath: ", urlPath)
	if urlPath == "/users" {
		//getAll and post
		switch r.Method {
		case http.MethodGet:
			uc.GetAllUsers(w, r)
		case http.MethodPost:
			uc.PostUser(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintf("Request URI doesn't exist")))
		}
	} else {
		if urlPath == "/users/" {
			uc.GetAllUsers(w, r)
			return
		}
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		fmt.Println("Matches ", matches)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.GetUser(w, id)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// go doesn't have constructor to initialiaze the userIDPattern
// instead we use constructor func to do the work for us
// constructor function name start with "new" followed by object
func newUserController() *userController {
	return &userController{userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`)}
}

func (uc *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request object : ", *r)
	allUsers, err := models.GetUsers()
	if err != nil {
		println("Error getting all users")
	}
	if allUsers == nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("No Users present currently")))
		return
	}
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJson(allUsers, w)
}

func (uc *userController) GetUser(w http.ResponseWriter, userId int) {
	user, err := models.GetUserByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error in Getting user : %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJson(user, w)
}

func ParseUserBody(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("Error in Parsing User Request Body: %v", err.Error())
	}
	return user, nil
}

func (uc *userController) PostUser(w http.ResponseWriter, r *http.Request) {
	user, err := ParseUserBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error in ParseUserBody : %v", err.Error())))
	}
	addedUser, err := models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error in Adding User : %v", err.Error())))
	}
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJson(addedUser, w)
}

// extra route controller for "/sample"
type sampleController struct {
	sampleIDPattern *regexp.Regexp // to help route
}

func (sc sampleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Sample"))
}

func newSampleController() *sampleController {
	return &sampleController{sampleIDPattern: regexp.MustCompile(`^/sample/(\d+)/?`)}
}
