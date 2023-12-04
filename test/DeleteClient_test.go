package test

import (
	"github.com/ao-space/platform-sdk-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteClient(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	require.NoError(t, client.DeleteClient(&platform.DeleteClientRequest{
		UserId:     "xxx",
		ClientUUID: "xxx",
	}))
}
