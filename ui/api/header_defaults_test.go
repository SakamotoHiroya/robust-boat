package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PingOKMessage(t *testing.T) {
	t.Parallel()
	h := &PingOK{}
	h.setDefaults()
	require.Equal(t, h.Message, PingOKMessage)
}
