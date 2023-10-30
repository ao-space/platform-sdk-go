package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	require.NoError(t, client.DeleteUser("xxx"))
}
