package test

import (
	"github.com/ao-space/platform-sdk-go"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModifyUserDomainName(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.ModifyUserDomain(&platform.ModifyUserDomainRequest{
		UserId:    "1",
		Subdomain: "user_one",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "\"success\":true")
	resp, err = client.ModifyUserDomain(&platform.ModifyUserDomainRequest{
		UserId:    "2",
		Subdomain: "user-two",
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "\"success\":false")
}
