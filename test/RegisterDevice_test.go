package test

import (
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterDevice(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.RegisterDevice()
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "networkClient")
}
