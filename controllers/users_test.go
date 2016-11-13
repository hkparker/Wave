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
			"{\"name\": \"Foober Doober\"}",
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

func TestUserCannotChangeTheirNameWithNonJSONData(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"this isn't going to work",
		)),
	)

	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	req.AddCookie(&cookie)

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

func TestUserCannotChangeNameWithMissingKey(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/name",
		strings.NewReader(fmt.Sprintf(
			"{\"something\": \"else\"}",
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
			assert.Equal("no username provided", params["error"])
		}
	}
}

func TestUserCanUpdatePassword(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	user.SetPassword("hunter2")
	session_id, _ := user.NewSession()
	user.NewSession()
	assert.Equal(2, len(user.Sessions))
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/password",
		strings.NewReader(fmt.Sprintf(
			"{\"old_password\": \"hunter2\", \"new_password\": \"1234\"}",
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
	user.Reload()
	assert.Equal(true, user.ValidAuthentication("1234"))
	//assert.Equal(1, len(user.Sessions))
}

func TestUserCannotUpdatePasswordWithBadPassword(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/password",
		strings.NewReader(fmt.Sprintf(
			"{\"old_password\": \"wrong\", \"new_password\": \"1234\"}",
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
}

func TestAdminCanAssignUserPassword(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	user.NewSession()
	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/assign-password",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"1234\"}",
			user.Username,
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
	user.Reload()
	assert.Equal(true, user.ValidAuthentication("1234"))
	//assert.Equal(0, len(user.Sessions))
}

func TestUserCannotAssignUserPassword(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	not_admin := models.CreateTestUser([]string{})
	session_id, _ := not_admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/assign-password",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"1234\"}",
			user.Username,
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
}

func TestAdminCanDeleteUser(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/delete",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\"}",
			user.Username,
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
	user.Reload()
	assert.NotNil(user.DeletedAt)
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)
	models.DropUsers()

	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	var names []map[string]string
	var params map[string]string

	req, err := http.NewRequest(
		"GET",
		testing_endpoint+"/users",
		strings.NewReader(fmt.Sprintf("")),
	)
	req.AddCookie(&cookie)
	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&names)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(names) {
		assert.Equal(1, len(names))
		for _, user := range names {
			if name, ok := user["name"]; assert.Equal(true, ok) {
				assert.Equal("Wifi Jackson", name)
			}
			if adm, ok := user["admin"]; assert.Equal(true, ok) {
				assert.Equal("true", adm)
			}
		}
	}

	req, err = http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"number1\"}",
		)),
	)
	req.AddCookie(&cookie)
	assert.Nil(err)
	resp, err = testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&params)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(params) {
		if assert.NotNil(params["success"]) {
			assert.Equal("true", params["success"])
		}
	}

	req, err = http.NewRequest(
		"GET",
		testing_endpoint+"/users",
		strings.NewReader(fmt.Sprintf("")),
	)
	req.AddCookie(&cookie)
	assert.Nil(err)
	resp, err = testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&names)
	if !assert.Nil(err) {
		assert.Nil(err.Error())
	}
	if assert.NotNil(names) {
		assert.Equal(2, len(names))
	}
}
