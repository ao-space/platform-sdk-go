package test

import (
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/big-dust/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterClient(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.RegisterClient(&platform.RegisterClientRequest{
		UserId:     "1",
		ClientUUID: "1",
		ClientType: "client_auth",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "userId")
	require.Contains(t, utils.ToString(resp), "clientType")
}
