package test

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.RegisterUser(&platform.RegisterUserRequest{
		UserID:     "xxx",
		Subdomain:  "xxx",
		UserType:   "user_admin",
		ClientUUID: uuid.New().String(),
	})
	resp, err = client.RegisterUser(&platform.RegisterUserRequest{
		UserID:     "xxx",
		Subdomain:  "xxx",
		UserType:   "user_member",
		ClientUUID: uuid.New().String(),
	})
	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "userId")
	require.Contains(t, utils.ToString(resp), "userType")
	require.Contains(t, utils.ToString(resp), "userDomain")
}
