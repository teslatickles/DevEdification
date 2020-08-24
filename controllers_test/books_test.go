package controllers_test

import (
	"fmt"
	. "github.com/DevEdification/v2/controllers_test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// need to handle these tests with a mock db????
// or some other way of handling ID not matching after running intial tests
var globalEntryID uint = 1

func TestCreateBook(t *testing.T) {
	bod, expected := map[string]interface{} {
		"title":   "The Wind-up Bird Chronicles",
		"release": "1995",
		"author":  "Haruki Murakami",
		"URL":     "https://google.com",
	}, fmt.Sprintf(`{"id":%v,"title":"The Wind-up Bird Chronicles","release":"1995","author":"Haruki Murakami","url":"https://google.com"}`, globalEntryID)

	endpoint := ControllerPrefix + "books"
	got := FetchPostResponseObject(t, endpoint, bod)

	//globalEntryID++

	assertions := assert.New(t)
	assertions.Equal(expected, got, "The returned response %v should match expected %v", got, expected)
}

func TestFindBooks(t *testing.T) {
	endpoint := ControllerPrefix + "books"

	got := FetchGetResponseObject(t, endpoint)
	expected := fmt.Sprintf(`[{"id":%v,"title":"The Wind-up Bird Chronicles","release":"1995","author":"Haruki Murakami","url":"https://google.com"}]`, globalEntryID)

	assertions := assert.New(t)
	assertions.Equal(expected, got, "The returned response %v should match expected %v", got, expected)
}

func TestFindBook(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "books/" + fmt.Sprint(entryID)

	expected := `{"id":1,"title":"The Wind-up Bird Chronicles","release":"1995","author":"Haruki Murakami","url":"https://google.com"}`
	got := FetchGetResponseObject(t, endpoint)

	assertions := assert.New(t)
	assertions.Equal(expected, got)
}

func TestUpdateBook(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "books/" + fmt.Sprint(entryID)

	expected := `{"id":1,"title":"The Wind-up Bird Chronicles","release":"1995","author":"Haruki Murakami","url":"https://google.com"}`
	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)

	bod := map[string]interface{}{"title":"Norwegian Wood", "release":"1987"}
	newExpected := `{"id":1,"title":"Norwegian Wood","release":"1987","author":"Haruki Murakami","url":"https://google.com"}`

	var nowGot string
	if ConfirmEntryUpdate(t, endpoint, bod) {
		nowGot = FetchGetResponseObject(t, endpoint)
	} else {
		t.Logf("Bad time with confirming the PATCH call as okay")
	}

	assertion.Equal(newExpected, nowGot)
}

func TestDeleteBook(t *testing.T) {
	entryID := 1
	endpoint := ControllerPrefix + "books/" + fmt.Sprint(entryID)

	isDeleted := ConfirmEntryDeletion(t, endpoint)

	globalEntryID += 1

	assertions := assert.New(t)
	assertions.True(isDeleted)
}
