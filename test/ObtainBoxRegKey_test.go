package test

import (
	"crypto/tls"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
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
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = platform.NewClientWithHost("xxx", tr)

	resp, err := client.SetRequestID("1111").ObtainBoxRegKey(&platform.ObtainBoxRegKeyRequest{
		BoxUUID:    "xxx",
		ServiceIds: []string{"xxx"},
	})

	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "tokenResults")

}
