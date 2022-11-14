package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TestSuite interface {
	Init()
	TearDownAllSuite()
	T() *testing.T
	GetRouter() *gin.Engine
}

type ConcreteTestSuite struct {
	suite.Suite
	router             *gin.Engine
	jwtToken           string
	databaseConnection *sql.DB
}

func (suite *ConcreteTestSuite) Init() {
	suite.databaseConnection = initLocalDatabase(suite.T())
	suite.router = framework.CreateRoutes("mockKey", suite.databaseConnection)
	suite.jwtToken = "mockKey"
}

func (suite *ConcreteTestSuite) TearDownAllSuite() {
	_, err := suite.databaseConnection.Exec("DELETE from users")
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = suite.databaseConnection.Exec("DELETE from expenses")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *ConcreteTestSuite) GetRouter() *gin.Engine {
	return suite.router
}
