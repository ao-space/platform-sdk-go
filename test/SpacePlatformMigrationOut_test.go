package test

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpacePlatformImmigrate(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.SpacePlatformMigrationOut(&platform.SpacePlatformMigrationOutRequest{
		UserDomainRouteInfos: []platform.UserDomainRouteInfo{
			platform.UserDomainRouteInfo{
				UserId:             "1",
				UserDomainRedirect: "userone.ao.space",
			},
		}})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "userDomainRouteInfos")
}
