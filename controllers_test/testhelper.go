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

var ControllerPrefix = "http://localhost:8080/"

func PerformRequest(r http.Handler, method, path string, body []byte) (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, req
}

func FetchGetResponseObject(t *testing.T, endpoint string) string {
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := PerformRequest(router, "GET", endpoint, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	CheckStatusCode(t, resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	got := TrimResponse(body)
	return got
}

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
	_, req := PerformRequest(router, "POST", endpoint, requestBody)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	CheckStatusCode(t, resp)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	got := TrimResponse(body)

	return got
}

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
	_, req := PerformRequest(router, "PATCH", endpoint, requestBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	isStatusOkay := CheckStatusCode(t, resp) == http.StatusOK

	return isStatusOkay
}

func ConfirmEntryDeletion(t *testing.T, endpoint string) bool {
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	router := gin.Default()
	_, req := PerformRequest(router, "DELETE", endpoint, nil)

	_, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//CheckStatusCode(t, resp)

	_, fetchReq := PerformRequest(router, "GET", endpoint, nil)

	fetchResponse, err := client.Do(fetchReq)
	if err != nil {
		log.Fatalln(err)
	}

	isDeleted := CheckStatusCode(t, fetchResponse) == http.StatusBadRequest

	return isDeleted
}

func TrimResponse(body []byte) string {
	rawBody := strings.TrimSuffix(string(body), `}`)
	got := strings.TrimPrefix(rawBody, `{"data":`)
	return got
}

func CheckStatusCode(t *testing.T, resp *http.Response) int {
	if status := resp.StatusCode; status != http.StatusOK {
		t.Logf("handler returned wrong status code: got %v want %v",
			fmt.Sprint(status), http.StatusOK)
		//return resp.StatusCode
	} else {
		t.Logf("passed with Status Code %v", status)
	}
	return resp.StatusCode
}
