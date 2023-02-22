package helper

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMustGetEnv(t *testing.T) {
	t.Run("EnvOrDefault", func(t *testing.T) {
		os.Setenv("TEST_KEY", "TEST_VALUE")
		assert.Equal(t, "TEST_VALUE", EnvOrDefault("TEST_KEY", "DEFAULT_VALUE"))
		assert.Equal(t, "DEFAULT_VALUE2", EnvOrDefault("TEST_KEY2", "DEFAULT_VALUE2"))
	})

	t.Run("TestMustGetEnv", func(t *testing.T) {
		// set env variable
		err := os.Setenv("TEST_KEY", "TEST_VALUE")
		assert.Nil(t, err)

		// test
		v, err := GetEnv("TEST_KEY")
		assert.Nil(t, err)
		assert.Equal(t, "TEST_VALUE", v)

		// unset env variable
		err = os.Unsetenv("TEST_KEY")
		assert.Nil(t, err)
	})

	t.Run("TestShort", func(t *testing.T) {
		assert.Equal(t, "1234567", Short("123456789"))
		assert.Equal(t, "1234567", Short("1234567"))
		assert.Equal(t, "1234", Short("1234"))
		assert.Equal(t, "", Short(""))
	})
}
