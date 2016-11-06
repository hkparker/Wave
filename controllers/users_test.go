package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/hkparker/Wave/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestAdminCanCreateUser(t *testing.T) {
	assert := assert.New(t)
	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"samsepi0l\"}",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

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
	_, db_err := models.UserByUsername("samsepi0l")
	assert.Nil(db_err)
}

func TestUserCannotCreateUser(t *testing.T) {
	assert := assert.New(t)
	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"notallowed@example.com\"}",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

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

func TestCreateUserWithInvalidData(t *testing.T) {
	assert := assert.New(t)
	user := models.CreateTestUser([]string{"admin"})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"this is not valid json",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(400, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var params map[string]string
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["error"]) {
			assert.Equal("error parsing json", params["error"])
		}
	}

}

func TestCreateUserWithNoData(t *testing.T) {
	assert := assert.New(t)
	user := models.CreateTestUser([]string{"admin"})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(400, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var params map[string]string
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["error"]) {
			assert.Equal("error parsing json", params["error"])
		}
	}

}

func TestUserCanChangeTheirName(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"Foober Doober\"}",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

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
	user.Reload()
	assert.Equal("Foober Doober", user.Name)
}

func TestUserNameChangeWithInvalidData(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"this isn't even json",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(400, resp.StatusCode)
}

func TestUserNameChangeWithMissingKey(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"{\"hey\": \"what's up\"}",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(400, resp.StatusCode)
}

// test user can use password reset link
// test password reset links expires in 24 hours
// test err when password reset token data is bad

// test user can password reset
// test password reset with non-json
// test password reset with missing password key

// test destroy user removes user
// test user cannot destroy other accounts
// test admin can destroy other accounts
// test err when admin provides bad data

// test current user with valid session
// test current user with bad session data
// test current user with no session cookie
