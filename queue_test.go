package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newQueue(t *testing.T) {
	_, err := newQueue(&Tasks{Name: ""})
	require.Error(t, err)

	_, err = newQueue(&Tasks{Name: "Task set 1", Tasks: []Task{}})
	require.Error(t, err)

	_, err = newQueue(&Tasks{Name: "Task set 1", Tasks: []Task{{Name: ""}}})
	require.Error(t, err)

	_, err = newQueue(&Tasks{Name: "Task set 1", Tasks: []Task{{Name: "Task 1", Exec: []string{}}}})
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

	results, err := newQueue(&tasks)
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(results.AsyncQueue))
	assert.EqualValues(t, 1, len(results.SequentialQueue))
}
