package controllers

import (
	"fmt"
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
	assert.NotNil(req)
	// expect user in the db
}
