package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	log.Println("Received request to authenticate user")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Println("This is user email: ", requestPayload.Email)

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	log.Println("This is fetched user email", user.Email)

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	log.Println("User is valid")

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	log.Println("Sending response")

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Println("Response sent successfully")
}
