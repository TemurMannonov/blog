package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "asdf1234"

	hashedPassword, err := HashPassword(password)
	fmt.Println(hashedPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}
