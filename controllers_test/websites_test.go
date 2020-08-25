package controllers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWebsite(t *testing.T) {
	bod, expected := map[string]interface{}{
		"title":   "Real Python",
		"tech":    "Python",
		"company": "non-affiliated",
		"author":  "Dan",
		"url":     "realpython.com",
	}, `{"id":1,"title":"Real Python","tech":"Python","company":"non-affiliated","author":"Dan","url":"realpython.com"}`

	endpoint := ControllerPrefix + "websites/"

	got := FetchPostResponseObject(t, endpoint, bod)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindWebsite(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "websites/" + fmt.Sprint(entryID)

	expected := `{"id":1,"title":"Real Python","tech":"Python","company":"non-affiliated","author":"Dan","url":"realpython.com"}`

	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindWebsites(t *testing.T) {
	expected := `[{"id":1,"title":"Real Python","tech":"Python","company":"non-affiliated","author":"Dan","url":"realpython.com"}]`

	endpoint := ControllerPrefix + "websites/"
	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestUpdateWebsite(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "websites/" + fmt.Sprint(entryID)

	expected := `{"id":1,"title":"Real Python","tech":"Python","company":"non-affiliated","author":"Dan","url":"realpython.com"}`

	//initialRecord := FetchGetResponseObject(t, endpoint)
	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)

	bod := map[string]interface{}{"url": "https://www.realpython.com"}

	newExpected := `{"id":1,"title":"Real Python","tech":"Python","company":"non-affiliated","author":"Dan","url":"https://www.realpython.com"}`

	var nowGot string
	if ConfirmEntryUpdate(t, endpoint, bod) {
		nowGot = FetchGetResponseObject(t, endpoint)
	} else {
		t.Logf("Bad time with confirming the PATCH call as okay")
	}

	assertion.Equal(newExpected, nowGot)
}

func TestDeleteWebsite(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "websites/" + fmt.Sprint(entryID)

	isDeleted := ConfirmEntryDeletion(t, endpoint)

	assertion := assert.New(t)
	assertion.True(isDeleted)
}
