package test

import (
	"github.com/ao-space/platform-sdk-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateUserDomainName(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	resp, err := client.GenerateUserDomain(&platform.GenerateUserDomainRequest{
		EffectiveTime: 12,
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}
