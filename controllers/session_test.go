package controllers

import (
	"fmt"
	"github.com/hkparker/Wave/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestNewSessionCreatesSession(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	user.SetPassword("hunter2")
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/sessions/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"hunter2\"}",
			user.Username,
		)),
	)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	if header, ok := resp.Header["Set-Cookie"]; assert.Equal(true, ok) {
		if assert.NotEqual(len(header), 0) {
			components := strings.Split(header[0], " ")
			assert.NotEqual(0, len(components))
			cookie := components[0]
			session_id := cookie[13 : len(cookie)-1]
			found_user, err := models.UserFromSessionCookie(session_id)
			assert.Nil(err)
			assert.Equal(user.Username, found_user.Username)
		}
	}
}

func TestNewSessionErrorWithInvalidData(t *testing.T) {
	assert := assert.New(t)

	models.CreateTestUser([]string{})
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/sessions/create",
		strings.NewReader(fmt.Sprintf(
			"let me in ok?",
		)),
	)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(400, resp.StatusCode)
	_, ok := resp.Header["Set-Cookie"]
	assert.Equal(false, ok)
}

func TestNewSessionErrorWithInvalidCredentials(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	user.SetPassword("hunter2")
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/sessions/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"hunter3\"}",
			user.Username,
		)),
	)

	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(401, resp.StatusCode)
	_, ok := resp.Header["Set-Cookie"]
	assert.Equal(false, ok)
}

func TestDeleteSessionDeletesSession(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/sessions/delete",
		strings.NewReader(""),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal("/login", resp.Request.URL.Path)

	_, err = models.SessionFromID(session_id)
	if assert.NotNil(err) {
		assert.Equal("record not found", err.Error())
	}
}
