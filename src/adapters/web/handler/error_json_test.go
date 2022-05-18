package handler_test

import (
	"testing"

	"github.com/sergio/go-hexagonal/adapters/web/handler"
	"github.com/stretchr/testify/require"
)

func TestJsonError(t *testing.T) {
	error := handler.JsonError("error")

	require.Equal(t, []byte(`{"message":"error"}`), error)
}
