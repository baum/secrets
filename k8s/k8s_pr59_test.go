// +build k8s

package k8s

import (
	"fmt"
	"testing"

	"github.com/libopenstorage/secrets"
	"github.com/stretchr/testify/require"
)

// TestGetInvalidSecretId tests PR59
func TestGetInvalidSecretId(t *testing.T) {
	s, _ := New(nil)
	keyContext := map[string]string{
		SecretNamespace: "default",
	}
	_, err := s.GetSecret("invalid", keyContext)
	invalid := (err == secrets.ErrInvalidSecretId)
	fmt.Printf("GetSecret err %v", err)
	require.True(t, invalid, "GetSecret for non-existent secret should return ErrInvalidSecretId")
}
