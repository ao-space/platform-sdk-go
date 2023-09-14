package test

import (
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/big-dust/platform-sdk-go/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

var client *platform.Client

func TestObtainBoxRegKey(t *testing.T) {

	tr := &http.Transport{
		MaxConnsPerHost: 10,
		IdleConnTimeout: 30 * time.Second,
		MaxIdleConns:    20,
	}

	client, _ = platform.NewClientWithHost("platform.muxixyz.com", tr)

	resp, err := client.ObtainBoxRegKey(&platform.ObtainBoxRegKeyRequest{
		BoxUUID:    "1",
		ServiceIds: []string{"10001"},
	})

	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "tokenResults")

}
