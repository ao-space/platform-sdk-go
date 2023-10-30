package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteDevice(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	require.NoError(t, client.SetRequestId("xxx").DeleteDevice())
}
