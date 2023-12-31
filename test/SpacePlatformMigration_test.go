package test

import (
	"github.com/ao-space/platform-sdk-go"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpacePlatformMigrate(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.SpacePlatformMigration(&platform.SpacePlatformMigrationRequest{
		NetworkClientId: "xxx",
		UserInfos: []platform.UserMigrationInfo{
			platform.UserMigrationInfo{
				UserId:     "1",
				UserDomain: "user-one",
				UserType:   "user_admin",
				ClientInfos: []platform.ClientInfo{
					platform.ClientInfo{
						ClientUUID: "1",
						ClientType: "client_auth",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "netWorkClient")
	require.Contains(t, utils.ToString(resp), "userInfos")
}
