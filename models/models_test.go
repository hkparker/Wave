package models

import (
	"database/sql/driver"
	"flag"
	"github.com/erikstmartin/go-testdb"
	"github.com/hkparker/Wave/helpers"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	seedModelsTests()
	exitcode := m.Run()
	os.Exit(exitcode)
}

func seedModelsTests() {
	helpers.TestDB().CreateTable(&User{})
	helpers.TestDB().CreateTable(&Session{})
	user := User{
		Name:  "Joe Hackerman",
		Email: "modeltest1@example.com",
	}
	helpers.TestDB().Create(&user)

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "name"}
		result := `
		1,Tim
		2,Joe
		3,Bob
		`
		return testdb.RowsFromCSVString(columns, result), nil
	})
}

func TestDatabaseSeeded(t *testing.T) {
	var users []User
	helpers.TestDB().Find(&users)
	if (users[0].Name != "Tim") || len(users) != 3 {
		t.Errorf("Unexcepted result returned")
	}
}
