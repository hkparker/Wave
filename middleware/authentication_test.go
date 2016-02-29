package middleware

import (
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthSeeded(t *testing.T) {
	user := models.User{}
	helpers.DB().First(&user)
	assert.Equal(t, user.Name, "Joe Hackerman")
}

func TestActiveSessionTrueWithSessionAndUpdatesTime(t *testing.T) {

}
