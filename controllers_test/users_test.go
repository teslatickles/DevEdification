package controllers_test

import (
	"fmt"
	. "github.com/DevEdification/v2/controllers_test"
	"github.com/DevEdification/v2/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	bod, expected := map[string]interface{}{
		"username": "lampin_larry",
		"password": "serenitynow",
		"email": 	"hohoho@thotbotbootyhole.word",
		"role":		"member",
	}, `{"id":1,"username":"lampin_larry","email":"hohoho@thotbotbootyhole.word","role":"member"}`

	endpoint := ControllerPrefix + "users"
	got := FetchPostResponseObject(t, endpoint, bod)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindUser(t *testing.T) {
	expected := `{"id":1,"username":"lampin_larry","email":"hohoho@thotbotbootyhole.word","role":"member"}`
	endpoint := ControllerPrefix + "users/" + fmt.Sprint(1)

	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestDeleteUser(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "users/" + fmt.Sprint(entryID)
	isDeleted := ConfirmEntryDeletion(t, endpoint)
	assertion := assert.New(t)
	assertion.True(isDeleted)
}

func TestLogin(t *testing.T) {
	//validUser := models.User{
	//	Username: "momonono",
	//	Password: "password",
	//	Email: 	  "meow@moo.neigh",
	//	Role:	  "member",
	//}
	//validBod := map[string]interface{}{
	//	"username": validUser.Username,
	//	"password": validUser.Password,
	//	"email":    validUser.Email,
	//	"role":     "member",
	//}

	invalidUser := models.User{
		Username: "papatofu",
		Password: "pwd",
		Email:    "loveisreal@maybe.barf",
		Role:     "guest",
	}
	invalidBod := map[string]interface{}{
		"username": invalidUser.Username,
		"password": invalidUser.Password,
		"email":    invalidUser.Email,
		"role":     "member",
	}

	// should return JSON containing jwt
	endpoint := ControllerPrefix + "login/"
	//got := FetchPostResponseObject(t, endpoint, validBod)
	unauthorizedGot := FetchPostResponseObject(t, endpoint, invalidBod)

	assertion := assert.New(t)
	//assertion.Equal(147, len(got))
	fmt.Println(unauthorizedGot)
	assertion.Equal(http.StatusUnauthorized, unauthorizedGot)
}
