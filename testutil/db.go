package testutil

import (
	"ZachIgarz/test-beer/config"
	"ZachIgarz/test-beer/infrastructure/datastore"
	"testing"

	"github.com/stretchr/testify/assert"
)

//ConfigDbTest set the test database to zero initialize
func ConfigDbTest(t *testing.T) {
	//environment variables are set to connect to a test database
	err := config.SetConfigs("localhost", "testingdb", "postgres", "secure_pass_here", true)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	//connects to test database
	err = datastore.NewDBConn()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	datastore.ResetDatabase()
	if err != nil {
		assert.Fail(t, err.Error())
	}
}
