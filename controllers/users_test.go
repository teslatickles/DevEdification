package controllers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	bod, expected := map[string]interface{}{
		"username":"lampin_larry",
		"password":"serenitynow",
	}, `{"id":1,"username":"lampin_larry"}`

	endpoint := controllerPrefix + "users"
	got := FetchPostResponseObject(t, endpoint, bod)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindUser(t *testing.T) {
	expected := `{"id":1,"username":"lampin_larry"}`
	endpoint := controllerPrefix + "users/" + fmt.Sprint(1)

	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestDeleteUser(t *testing.T) {
	entryID := 1
	endpoint := controllerPrefix + "users/" + fmt.Sprint(entryID)
	isDeleted := ConfirmEntryDeletion(t, endpoint)
	assertion := assert.New(t)
	assertion.True(isDeleted)
}
