package main

import (
	"testing"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newParser(t *testing.T) {
	p := ""
	pJson := "./examples/tasks.json"
	pYaml := "./examples/tasks.yaml"
	pToml := "./examples/tasks.toml"
	pHcl := "./examples/tasks.tf"
	k := koanf.New(".")

	_, err := newParser(&p, k)
	require.Error(t, err)

	_, err = newParser(&pJson, k)
	require.NoError(t, err)

	_, err = newParser(&pYaml, k)
	require.NoError(t, err)

	_, err = newParser(&pToml, k)
	require.NoError(t, err)

	_, err = newParser(&pHcl, k)
	require.NoError(t, err)

	tasks := Tasks{
		Name:        "My tasks set",
		Description: "This is my set of tasks",
		Tasks: []Task{
			{
				IsAsync: true, IsSudo: false,
				Name: "async task 1", Description: "print string from the async task 1",
				Exec: []string{"echo", "hello, async task 1!"},
			},
			{
				IsAsync: true, IsSudo: false,
				Name: "async task 2 with error", Description: "print string from the async task 2 with error",
				Exec: []string{"error", "some error in async task 2"},
			},
			{
				IsAsync: false, IsSudo: false,
				Name: "sequential task 1", Description: "print string from the sequential task 1",
				Exec: []string{"echo", "hello, sequential task 1!"},
			},
		},
	}

	result, _ := newParser(&pJson, k)
	assert.EqualValues(t, &tasks, result)

	result, _ = newParser(&pYaml, k)
	assert.EqualValues(t, &tasks, result)

	result, _ = newParser(&pToml, k)
	assert.EqualValues(t, &tasks, result)

	result, _ = newParser(&pHcl, k)
	assert.EqualValues(t, &tasks, result)
}
