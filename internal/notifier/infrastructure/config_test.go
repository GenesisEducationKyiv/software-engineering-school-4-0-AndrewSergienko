package infrastructure_test

import (
	"go_service/internal/notifier/infrastructure"
	"os"
	"path/filepath"
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

func TestGetServicesAPISettings(t *testing.T) {
	projectRoot, err := os.Getwd()
	assert.NoError(t, err)

	configPath := filepath.Join(projectRoot, "..", "..", "..", "conf", "config.toml")

	_, err = infrastructure.GetServicesAPISettings(configPath)
	assert.NoError(t, err)
}

func TestGetServicesAPISettingsError(t *testing.T) {
	_, err := infrastructure.GetServicesAPISettings("wrong_way")
	assert.Error(t, err)
}
