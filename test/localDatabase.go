package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/joho/godotenv"
)

func initLocalDatabase(t *testing.T) *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	databaseConnection := framework.GetDatabaseConnection()
	_, err = databaseConnection.Exec("DELETE from users")

	if err != nil {
		t.Fatal(err)
	}

	return databaseConnection
}
