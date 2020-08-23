package controllers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateVizNug(t *testing.T) {
	endpoint := controllerPrefix + "viznugs/"

	bod, expected := map[string]interface{}{
		"title":   "Writing API unit tests with Go",
		"tech":    "golang, api, tdd",
		"company": "hunt",
		"author":  "Hunter Hartline",
		"gcsc":    false,
		"url":     "fuction.pot",
	}, `{"id":1,"tech":"golang, api, tdd","title":"Writing API unit tests with Go","author":"Hunter Hartline","gcsc":false,"url":"fuction.pot"}`

	got := FetchPostResponseObject(t, endpoint, bod)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindVizNug(t *testing.T) {
	entryID := 1
	endpoint := controllerPrefix + "viznugs/" + fmt.Sprint(entryID)

	expected := `{"id":1,"tech":"golang, api, tdd","title":"Writing API unit tests with Go","author":"Hunter Hartline","gcsc":false,"url":"fuction.pot"}`

	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestFindVizNugs(t *testing.T) {
	endpoint := controllerPrefix + "viznugs/"

	expected := `[{"id":1,"tech":"golang, api, tdd","title":"Writing API unit tests with Go","author":"Hunter Hartline","gcsc":false,"url":"fuction.pot"}]`
	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)
}

func TestUpdateVizNug(t *testing.T) {
	entryID := 1
	endpoint := controllerPrefix + "viznugs/" + fmt.Sprint(entryID)

	expected := `{"id":1,"tech":"golang, api, tdd","title":"Writing API unit tests with Go","author":"Hunter Hartline","gcsc":false,"url":"fuction.pot"}`

	got := FetchGetResponseObject(t, endpoint)

	assertion := assert.New(t)
	assertion.Equal(expected, got)

	bod := map[string]interface{}{"tech": "golang, api, tdd, gorm, gin, docker, swagger"}

	newExpected := `{"id":1,"tech":"golang, api, tdd, gorm, gin, docker, swagger","title":"Writing API unit tests with Go","author":"Hunter Hartline","gcsc":false,"url":"fuction.pot"}`

	var nowGot string
	if ConfirmEntryUpdate(t, endpoint, bod) {
		nowGot = FetchGetResponseObject(t, endpoint)
	} else {
		t.Logf("Bad time with confirming okay status on PATCH call")
	}

	assertion.Equal(newExpected, nowGot)
}

func TestDeleteVizNug(t *testing.T) {
	entryID := 1
	endpoint := controllerPrefix + "viznugs/" + fmt.Sprint(entryID)

	isDeleted := ConfirmEntryDeletion(t, endpoint)

	assertion := assert.New(t)
	assertion.True(isDeleted)
}
