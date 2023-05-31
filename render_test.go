package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResults_render(t *testing.T) {
	r := Results{}
	err := r.render()
	require.Error(t, err)

	r = Results{Success: []Result{{ID: "", Description: "", Output: ""}}, Fail: []Result{{ID: "", Description: "", Output: ""}}}
	err = r.render()
	require.NoError(t, err)
}
