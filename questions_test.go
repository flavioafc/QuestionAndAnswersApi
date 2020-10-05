package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize("localhost", "db")
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func TestGetQuestionByInvalidObjectID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/question/invalidobjectid", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Invalid input to ObjectIdHex" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Invalid input to ObjectIdHex'. Got '%s'", m["error"])
	}
}
func TestQuestionNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/question/5f4d736a59ed57cdff62bec3", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Question not found'. Got '%s'", m["error"])
	}
}
func TestCreateQuestion(t *testing.T) {
	var jsonStr = []byte(`{"text":"Question test ?", "image": "http://ec2/img/1237998478_question_pointer.jpg", "likes" : 100}`)

	req, _ := http.NewRequest("POST", "/api/v1/question", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["text"] != "Question test ?" {
		t.Errorf("Expected Question text to be 'Question test ?'. Got '%v'", m["text"])
	}

	if m["image"] != "http://ec2/img/1237998478_question_pointer.jpg" {
		t.Errorf("Expected image question to be 'http://ec2/img/1237998478_question_pointer.jpg'. Got '%v'", m["image"])
	}

	// JSON unmarshaling converts numbers to floats, when using map[string]interface{}
	if m["likes"] != 100.0 {
		t.Errorf("Expected likes to be '100'. Got '%v'", m["likes"])
	}
}
func TestGetQuestionByID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/question/5f4a68c17ee3884ed8168759", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}
