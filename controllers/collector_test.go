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

func TestNewCollectorCreatesAndDeletes(t *testing.T) {
	assert := assert.New(t)

	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/create",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 6\"}",
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
		if assert.NotNil(params["name"]) {
			assert.Equal("living room channel 6", params["name"])
		}
		if assert.NotNil(params["certificate"]) {
			assert.NotEqual("", params["certificate"])
		}
		if assert.NotNil(params["private_key"]) {
			assert.NotEqual("", params["private_key"])
		}
	}

	req, err = http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/delete",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 6\"}",
		)),
	)
	req.AddCookie(&cookie)
	assert.Nil(err)
	resp, err = testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)
}

func TestNewCollectorFailsWithoutAdmin(t *testing.T) {
	assert := assert.New(t)

	user := models.CreateTestUser([]string{})
	session_id, _ := user.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/create",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 1\"}",
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

func TestGetCollectors(t *testing.T) {
	assert := assert.New(t)

	admin := models.CreateTestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	session, err := models.SessionFromID(session_id)
	assert.Nil(err)
	cookie := session.HTTPCookie()
	var names []string
	var params map[string]string

	req, err := http.NewRequest(
		"GET",
		testing_endpoint+"/collectors",
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
		assert.Equal([]string{}, names)
	}

	req, err = http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/create",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 11\"}",
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
		if assert.NotNil(params["name"]) {
			assert.Equal("living room channel 11", params["name"])
		}
		if assert.NotNil(params["certificate"]) {
			assert.NotEqual("", params["certificate"])
		}
		if assert.NotNil(params["private_key"]) {
			assert.NotEqual("", params["private_key"])
		}
	}

	req, err = http.NewRequest(
		"GET",
		testing_endpoint+"/collectors",
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
		assert.Equal([]string{"living room channel 11"}, names)
	}

	req, err = http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/create",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 1\"}",
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
		if assert.NotNil(params["name"]) {
			assert.Equal("living room channel 1", params["name"])
		}
		if assert.NotNil(params["certificate"]) {
			assert.NotEqual("", params["certificate"])
		}
		if assert.NotNil(params["private_key"]) {
			assert.NotEqual("", params["private_key"])
		}
	}

	req, err = http.NewRequest(
		"GET",
		testing_endpoint+"/collectors",
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
		assert.Equal([]string{"living room channel 11", "living room channel 1"}, names)
	}

	req, err = http.NewRequest(
		"POST",
		testing_endpoint+"/collectors/delete",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"living room channel 11\"}",
		)),
	)
	req.AddCookie(&cookie)
	assert.Nil(err)
	resp, err = testing_client.Do(req)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	req, err = http.NewRequest(
		"GET",
		testing_endpoint+"/collectors",
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
		assert.Equal([]string{"living room channel 1"}, names)
	}
}
