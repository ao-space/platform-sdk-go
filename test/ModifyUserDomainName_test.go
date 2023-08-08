package test

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModifyUserDomainName(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.ModifyUserDomain(&platform.ModifyUserDomainRequest{
		UserId:    "xxx",
		Subdomain: "xxx",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "\"success\":true")
	resp, err = client.ModifyUserDomain(&platform.ModifyUserDomainRequest{
		UserId:    "xxx",
		Subdomain: "xxx",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "\"success\":false")
}
