package helper

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("EnvOrDefault", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "TEST_VALUE")
		assert.Nil(t, err)
		assert.Equal(t, "TEST_VALUE", EnvOrDefault("TEST_KEY", "DEFAULT_VALUE"))
		assert.Equal(t, "DEFAULT_VALUE2", EnvOrDefault("TEST_KEY2", "DEFAULT_VALUE2"))
	})

	t.Run("EnvOrDefaultString", func(t *testing.T) {
		assert.Equal(t, "test", EnvOrDefaultString("NOT_EXISTING", "test"))
		err := os.Setenv("TEST_KEY", "apud")
		assert.Nil(t, err)
		assert.Equal(t, "apud", EnvOrDefaultString("TEST_KEY", "test"))
	})

	t.Run("EnvOrDefaultInt64", func(t *testing.T) {
		assert.Equal(t, int64(4321), EnvOrDefaultInt64("NOT_EXISTING", 4321))
		err := os.Setenv("TEST_KEY", "1234")
		assert.Nil(t, err)
		assert.Equal(t, int64(1234), EnvOrDefaultInt64("TEST_KEY", 1234))
	})

	t.Run("TestGetEnv", func(t *testing.T) {
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

	t.Run("TestMustGetEnv", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "TEST_VALUE")
		assert.Nil(t, err)
		assert.Equal(t, "TEST_VALUE", MustGetEnv("TEST_KEY"))
	})

	t.Run("TestMustGetEnv Panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		MustGetEnv("TEST_KEY_PANIC")
	})

	t.Run("TestShort", func(t *testing.T) {
		assert.Equal(t, "1234567", Short("123456789"))
		assert.Equal(t, "1234567", Short("1234567"))
		assert.Equal(t, "1234", Short("1234"))
		assert.Equal(t, "", Short(""))
	})
}
