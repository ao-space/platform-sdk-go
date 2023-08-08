package test

import (
	"github.com/ao-space/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateUserDomainName(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.GenerateUserDomain(&platform.GenerateUserDomainRequest{
		EffectiveTime: "XXX",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}
