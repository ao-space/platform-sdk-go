package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetAbility(t *testing.T) {
	t.Run("ObtainBoxRegKey", TestObtainBoxRegKey)
	//
	response, err := client.GetAbility()
	require.NoError(t, err)
	fmt.Println(*response)
	for i := 0; i < 100; i++ {
		go func() {
			response, err = client.GetAbility()
			require.NoError(t, err)
			fmt.Println(*response)
		}()
	}
	time.Sleep(time.Second * 10)
	//threading.GoSafe(client.FlushAbilityWithDuration(time.Minute))
}
