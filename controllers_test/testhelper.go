package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// ControllerPrefix global exported variable referencing localhost url for controller calls
var ControllerPrefix = "http://localhost:8080/api/v1/"

// performRequest helper method to correctly handle executing gin requests necessary to test api
func performRequest(r http.Handler, method, path string, body []byte) (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, req
}

// FetchGetResponseObject fetch response returned from passed url (endpoint) to api
func FetchGetResponseObject(t *testing.T, endpoint string) string {
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := performRequest(router, "GET", endpoint, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	checkStatusCode(t, resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	got := trimResponse(body)
	return got
}

// FetchPostResponseObject fetch response returned from passed request (endpoint) using passed body (bod)
func FetchPostResponseObject(t *testing.T, endpoint string, bod map[string]interface{}) string {
	requestBody, err := json.Marshal(bod)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := performRequest(router, "POST", endpoint, requestBody)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	checkStatusCode(t, resp)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	got := trimResponse(body)

	return got
}

// ConfirmEntryUpdate confirm targeted record was updated based on id and body (bod)
func ConfirmEntryUpdate(t *testing.T, endpoint string, bod map[string]interface{}) bool {
	requestBody, err := json.Marshal(bod)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := performRequest(router, "PATCH", endpoint, requestBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	isStatusOkay := checkStatusCode(t, resp) == http.StatusOK

	return isStatusOkay
}

// ConfirmEntryDeletion confirm (return true or false) if a specific record has been deleted based on id
func ConfirmEntryDeletion(t *testing.T, endpoint string) bool {
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := performRequest(router, "DELETE", endpoint, nil)

	_, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//checkStatusCode(t, resp)

	_, fetchReq := performRequest(router, "GET", endpoint, nil)

	fetchResponse, err := client.Do(fetchReq)
	if err != nil {
		log.Fatalln(err)
	}

	isDeleted := checkStatusCode(t, fetchResponse) == http.StatusBadRequest

	return isDeleted
}

// trimResponse trim extraneous bits from returned response
func trimResponse(body []byte) string {
	rawBody := strings.TrimSuffix(string(body), `}`)
	got := strings.TrimPrefix(rawBody, `{"data":`)
	return got
}

// checkStatusCode check status code of returned response
func checkStatusCode(t *testing.T, resp *http.Response) int {
	if status := resp.StatusCode; status != http.StatusOK {
		t.Logf("handler returned wrong status code: got %v want %v",
			fmt.Sprint(status), http.StatusOK)
		//return resp.StatusCode
	} else {
		t.Logf("passed with Status Code %v", status)
	}
	return resp.StatusCode
}
