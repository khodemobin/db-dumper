package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("test load from yaml successfully", func(t *testing.T) {
		configPath := "../config.yaml"
		envPath := "../.env"
		cfg, err := LoadConfig(configPath, envPath)
		assert.NoError(t, err)

		assert.Equal(t, cfg.Tasks[0].Archive.Driver, "zip")
	})

	t.Run("test load from env successfully", func(t *testing.T) {
		configPath := "../config.yaml"
		envPath := "../.env"
		cfg, err := LoadConfig(configPath, envPath)
		assert.NoError(t, err)

		assert.Equal(t, cfg.Tasks[0].Archive.Password, "123456")
	})

	t.Run("test failed load", func(t *testing.T) {
		configPath := "../config_test.yaml"
		envPath := "../.env_test"
		_, err := LoadConfig(configPath, envPath)
		assert.Error(t, err)
	})
}
