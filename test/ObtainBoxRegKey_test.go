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

	client, _ = platform.NewClientWithHost(platform.AoSpaceDomain, tr)

	resp, err := client.SetRequestId("1111").ObtainBoxRegKey(&platform.ObtainBoxRegKeyRequest{
		BoxUUID:    "364b553c01dabb2764b2f2c0e721c1e860e308b1c7daed2671507d21434060ed",
		ServiceIds: []string{"10001"},
	})

	require.NoError(t, err)
	require.Contains(t, utils.ToString(resp), "tokenResults")

}
