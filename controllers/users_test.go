package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestCreateUserCreatesUser(t *testing.T) {
	assert := assert.New(t)
	admin := models.TestUser([]string{"admin"})
	session_id, _ := admin.NewSession()
	req, err := http.NewRequest(
		"POST",
		testing_endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"wave_session\": \"%s\", \"email\": \"newuser@example.com\"}",
			session_id,
		)),
	)
	assert.Nil(err)
	resp, err := testing_client.Do(req)
	assert.Nil(err)
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
	var created_user models.User
	database.DB().Where(models.User{Email: "fixit"}).First(&created_user)
	assert.Equal(true, created_user.OTPReset)
}
