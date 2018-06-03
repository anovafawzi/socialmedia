package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	dbFactory "github.com/anovafawzi/socialmedia/db"
	models "github.com/anovafawzi/socialmedia/models"
	server "github.com/anovafawzi/socialmedia/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var repo = dbFactory.NewSQLiteRepository("./dbsocmed.db")

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPostFriendConnection(t *testing.T) {
	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"friends\":[\"andy@example.com\",\"john@example.com\"]}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendconnection", bodyPost)

	// check the request gives a 201
	assert.Equal(t, http.StatusCreated, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success": true,
	}

	// do assert test
	assert.Equal(t, body["success"], result.Success)

	// clean data
	db := repo.InitDb()
	defer db.Close()
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", "andy@example.com", "john@example.com", "john@example.com", "andy@example.com").Delete(models.Relations{})
}

func TestPostFriendList(t *testing.T) {
	// create initial data
	db := repo.InitDb()
	defer db.Close()

	email1 := "andy@example.com"
	email2 := "john@example.com"

	addfriends1 := models.Relations{Email1: email1, Email2: email2, Friend: true, Subscribe: false, Block: false}
	addfriends2 := models.Relations{Email1: email2, Email2: email1, Friend: true, Subscribe: false, Block: false}
	db.Create(&addfriends1)
	db.Create(&addfriends2)

	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"email\":\"andy@example.com\"}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendlist", bodyPost)

	// check the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success": true,
		"friends": []string{"john@example.com"},
		"count":   1,
	}

	// do assert test
	assert.Equal(t, body["success"], result.Success)
	assert.Equal(t, body["count"], result.Count)
	assert.Equal(t, body["friends"], result.Friends)

	// clean data
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email1, email2, email2, email1).Delete(models.Relations{})
}

func TestPostFriendCommonList(t *testing.T) {
	// create initial data
	db := repo.InitDb()
	defer db.Close()

	email1 := "andy@example.com"
	email2 := "john@example.com"
	email3 := "common@example.com"

	addfriends1 := models.Relations{Email1: email1, Email2: email3, Friend: true, Subscribe: false, Block: false}
	addfriends2 := models.Relations{Email1: email3, Email2: email1, Friend: true, Subscribe: false, Block: false}
	addfriends3 := models.Relations{Email1: email2, Email2: email3, Friend: true, Subscribe: false, Block: false}
	addfriends4 := models.Relations{Email1: email3, Email2: email2, Friend: true, Subscribe: false, Block: false}
	db.Create(&addfriends1)
	db.Create(&addfriends2)
	db.Create(&addfriends3)
	db.Create(&addfriends4)

	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"friends\":[\"andy@example.com\",\"john@example.com\"]}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendcommonlist", bodyPost)

	// check the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success": true,
		"friends": []string{"common@example.com"},
		"count":   1,
	}

	// do assert test
	assert.Equal(t, body["success"], result.Success)
	assert.Equal(t, body["count"], result.Count)
	assert.Equal(t, body["friends"], result.Friends)

	// clean data
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email1, email3, email3, email1).Delete(models.Relations{})
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email2, email3, email3, email2).Delete(models.Relations{})
}

func TestPostFriendSubscribe(t *testing.T) {
	// create initial data
	db := repo.InitDb()
	defer db.Close()

	email1 := "lisa@example.com"
	email2 := "john@example.com"

	addfriends1 := models.Relations{Email1: email1, Email2: email2, Friend: true, Subscribe: false, Block: false}
	addfriends2 := models.Relations{Email1: email2, Email2: email1, Friend: true, Subscribe: false, Block: false}
	db.Create(&addfriends1)
	db.Create(&addfriends2)

	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"requestor\":\"lisa@example.com\",\"target\":\"john@example.com\"}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendsubscribe", bodyPost)

	// check the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success": true,
	}

	// get real value
	var relation models.Relations
	db.Where("email1 = ? AND email2 = ?", email1, email2).First(&relation)

	// do assert test
	assert.Equal(t, body["success"], result.Success)
	assert.Equal(t, true, relation.Subscribe)

	// clean data
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email1, email2, email2, email1).Delete(models.Relations{})
}

func TestPostFriendBlock(t *testing.T) {
	// create initial data
	db := repo.InitDb()
	defer db.Close()

	email1 := "andy@example.com"
	email2 := "john@example.com"

	addfriends1 := models.Relations{Email1: email1, Email2: email2, Friend: true, Subscribe: false, Block: false}
	addfriends2 := models.Relations{Email1: email2, Email2: email1, Friend: true, Subscribe: false, Block: false}
	db.Create(&addfriends1)
	db.Create(&addfriends2)

	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"requestor\":\"andy@example.com\",\"target\":\"john@example.com\"}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendblock", bodyPost)

	// check the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success": true,
	}

	// get real value
	var relation models.Relations
	db.Where("email1 = ? AND email2 = ?", email1, email2).First(&relation)

	// do assert test
	assert.Equal(t, body["success"], result.Success)
	assert.Equal(t, true, relation.Block)

	// clean data
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email1, email2, email2, email1).Delete(models.Relations{})
}

func TestPostFriendUpdates(t *testing.T) {
	// create initial data
	db := repo.InitDb()
	defer db.Close()

	email1 := "john@example.com"
	email2 := "lisa@example.com"

	addfriends1 := models.Relations{Email1: email1, Email2: email2, Friend: true, Subscribe: false, Block: false}
	addfriends2 := models.Relations{Email1: email2, Email2: email1, Friend: true, Subscribe: false, Block: false}
	db.Create(&addfriends1)
	db.Create(&addfriends2)

	// create POST body
	bodyPost := bytes.NewBuffer([]byte("{\"sender\":\"john@example.com\",\"text\":\"Hello World, kate@example.com\"}"))
	// get all route
	router := server.LoadRoute()
	// do POST
	w := performRequest(router, "POST", "/api/v1/friendupdates", bodyPost)

	// check the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// get the JSON response
	var result models.Result
	err := json.Unmarshal([]byte(w.Body.String()), &result)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	// get the value & check it exists or not
	// expected result body
	body := gin.H{
		"success":    true,
		"recipients": []string{"lisa@example.com", "kate@example.com"},
	}

	// do assert test
	assert.Equal(t, body["success"], result.Success)
	assert.Equal(t, body["recipients"], result.Recipients)

	// clean data
	db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", email1, email2, email2, email1).Delete(models.Relations{})
}
