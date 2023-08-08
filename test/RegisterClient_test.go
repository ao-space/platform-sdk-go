package test

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterClient(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.RegisterClient(&platform.RegisterClientRequest{
		UserId:     "xxx",
		ClientUUID: "xxx",
		ClientType: "xxx",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "userId")
	require.Contains(t, utils.ToString(resp), "clientType")
}
