package main

import (
	"testing"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/require"
)

func Test_newParser(t *testing.T) {
	p1 := ""
	p2 := "./examples/tasks.json"
	k := koanf.New(".")

	_, err := newParser(&p1, k)
	require.Error(t, err)

	_, err = newParser(&p2, k)
	require.NoError(t, err)
}
