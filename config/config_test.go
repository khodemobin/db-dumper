package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("test load successfully", func(t *testing.T) {
		path := "../config.yaml"
		cfg, err := LoadConfig(path)
		assert.NoError(t, err)

		assert.Equal(t, cfg.Tasks[0].ArchiveDriver, "zip")
	})

	t.Run("test failed load", func(t *testing.T) {
		path := "../config_test.yaml"
		_, err := LoadConfig(path)
		assert.Error(t, err)
	})
}
