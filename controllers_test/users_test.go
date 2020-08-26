package controllers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	//CleanupTestDatabase("users")

	//sql := `TRUNCATE TABLE core_test.users;`
	//t.Cleanup()

	bod, expected := map[string]interface{}{
		"username":"lampin_larry",
		"password":"serenitynow",
	}, `{"id":1,"username":"lampin_larry"}`

	endpoint := ControllerPrefix + "users"
	got := FetchPostResponseObject(t, endpoint, bod)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindUser(t *testing.T) {
	expected := `{"id":1,"username":"lampin_larry"}`
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

	//CleanupTestDatabase("users")
}
