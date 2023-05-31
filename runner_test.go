package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_runTasks(t *testing.T) {
	q := Queue{}
	_, err := q.runTasks()
	require.Error(t, err)

	q = Queue{
		AsyncQueue: []Task{
			{Name: "Task 1", IsSudo: true, Exec: []string{"echo", "hello, world!"}},
			{Name: "Task 2", IsSudo: false, Exec: []string{"error", "error"}},
		},
		SequentialQueue: []Task{
			{Name: "Task 1", IsSudo: true, Exec: []string{"echo", "hello, world!"}},
			{Name: "Task 2", IsSudo: false, Exec: []string{"error", "error"}},
		},
	}
	_, err = q.runTasks()
	require.NoError(t, err)
}
