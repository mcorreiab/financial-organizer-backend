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

	return framework.GetDatabaseConnection()
}
