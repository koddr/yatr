package main

import (
	"os"
	"testing"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newParser(t *testing.T) {
	k := koanf.New(".")

	p := ""
	_, err := newParser(&p, k)
	require.Error(t, err)

	_ = os.WriteFile("./bin/error.json", []byte("error"), 0x0755)
	_ = os.WriteFile("./bin/error.yaml", []byte("error"), 0x0755)
	_ = os.WriteFile("./bin/error.toml", []byte("error"), 0x0755)
	_ = os.WriteFile("./bin/error.tf", []byte("error"), 0x0755)

	eJson := "./bin/error.json"
	_, err = newParser(&eJson, k)
	require.Error(t, err)

	eYaml := "./bin/error.yaml"
	_, err = newParser(&eYaml, k)
	require.Error(t, err)

	eToml := "./bin/error.toml"
	_, err = newParser(&eToml, k)
	require.Error(t, err)

	eHcl := "./bin/error.tf"
	_, err = newParser(&eHcl, k)
	require.Error(t, err)

	_ = os.RemoveAll(eJson)
	_ = os.RemoveAll(eYaml)
	_ = os.RemoveAll(eToml)
	_ = os.RemoveAll(eHcl)

	pJson := "./examples/tasks.json"
	_, err = newParser(&pJson, k)
	require.NoError(t, err)

	pYaml := "./examples/tasks.yaml"
	_, err = newParser(&pYaml, k)
	require.NoError(t, err)

	pToml := "./examples/tasks.toml"
	_, err = newParser(&pToml, k)
	require.NoError(t, err)

	pHcl := "./examples/tasks.tf"
	_, err = newParser(&pHcl, k)
	require.NoError(t, err)

	pUnknown := "./examples/tasks.exe"
	_, err = newParser(&pUnknown, k)
	require.Error(t, err)

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
