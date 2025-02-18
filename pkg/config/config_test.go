package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_UnmarshalKeys(t *testing.T) {
	conf, err := NewConfig("")
	require.NoError(t, err)

	require.NoError(t, conf.unmarshalKeys("key1: secret1"))
	require.Equal(t, "secret1", conf.Keys["key1"])
}
