package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStatus(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	response, err := client.GetStatus()
	require.NoError(t, err)
	fmt.Println(*response)
}