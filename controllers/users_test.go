package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/hkparker/Wave/database"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestAdminCanCreateUser(t *testing.T) {
	assert := assert.New(t)
	admin := database.TestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"email\": \"newuser@example.com\"}",
			session_id,
		)),
	)
	req.Header.Set("wave_session", session_id)
	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var params map[string]string
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["success"]) {
			assert.Equal("true", params["success"])
		}
	}
	var created_user database.User
	db_err := database.DB().First(&created_user, "Email = ? ", "newuser@example.com")
	assert.Nil(db_err.Error)
}

func TestUserCannotCreateUser(t *testing.T) {
	assert := assert.New(t)
	user := database.TestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"email\": \"newuser@example.com\"}",
			session_id,
		)),
	)
	req.Header.Set("wave_session", session_id)
	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(401, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var params map[string]string
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["error"]) {
			assert.Equal("permission denied", params["error"])
		}
	}
}

func UserCanChangeTheirName(t *testing.T) {
	assert := assert.New(t)

	user := database.TestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"Foober Doober\"}",
			session_id,
		)),
	)
	req.Header.Set("wave_session", session_id)
	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var params map[string]string
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["success"]) {
			assert.Equal("true", params["success"])
		}
	}
	assert.Equal("Foober Doober", user.Name)
}
