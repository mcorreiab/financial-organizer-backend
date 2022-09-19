package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/joho/godotenv"
)

type testSuite struct {
	DatabaseConnection *sql.DB
	tests              []testRunner
	t                  *testing.T
}

type testRunner struct {
	Name    string
	Command func(*testing.T)
}

var databaseConnection *sql.DB

func newSuite(t *testing.T, tests []testRunner) testSuite {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	databaseConnection = framework.GetDatabaseConnection()

	return testSuite{databaseConnection, tests, t}
}

func (t testSuite) run() {
	for _, test := range t.tests {
		t.cleanDatabase()
		t.t.Run(test.Name, test.Command)
	}
}

func (t testSuite) cleanDatabase() {
	_, err := t.DatabaseConnection.Exec("DELETE from users")
	if err != nil {
		t.t.Fatal(err)
	}
}
