package infrastructure_test

import (
	"go_service/internal/rateservice/infrastructure"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchEnvStrictVariableSet(t *testing.T) {
	err := os.Setenv("TEST_ENV", "value")
	assert.NoError(t, err)
	defer func() {
		err = os.Unsetenv("TEST_ENV")
		assert.NoError(t, err)
	}()

	result := infrastructure.FetchEnv("TEST_ENV", true)
	assert.Equal(t, "value", result)
}

func TestFetchEnvStrictVariableNotSet(t *testing.T) {
	assert.Panics(t, func() {
		infrastructure.FetchEnv("NON_EXISTENT_ENV", true)
	})
}

func TestFetchEnvNonStrictVariableSet(t *testing.T) {
	err := os.Setenv("TEST_ENV", "value")
	assert.NoError(t, err)
	defer func() {
		err = os.Unsetenv("TEST_ENV")
		assert.NoError(t, err)
	}()

	result := infrastructure.FetchEnv("TEST_ENV", false)
	assert.Equal(t, "value", result)
}

func TestFetchEnvNonStrictVariableNotSet(t *testing.T) {
	result := infrastructure.FetchEnv("NON_EXISTENT_ENV", false)
	assert.Equal(t, "", result)
}
