package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newMigrationConfig(t *testing.T) {
	t.Run("success on provided", func(t *testing.T) {
		t.Setenv("POSTGRES_USER", "user")
		t.Setenv("POSTGRES_PASSWORD", "password")
		t.Setenv("POSTGRES_DB", "db")

		config, err := newConfig()

		require.NoError(t, err)
		assert.Equal(t, "db", config.DB)
		assert.Equal(t, "password", config.Password)
		assert.Equal(t, "user", config.User)
	})

	t.Run("fails on missing", func(t *testing.T) {
		_, err := newConfig()

		require.Error(t, err)
	})
}
