name        = "My tasks set"
description = "This is my set of tasks"

tasks = [
  {
    name            = "async task 1"
    description     = "print string from the async task 1"
    is_async        = true
    is_sudo         = false
    is_print_output = true
    exec            = ["echo", "hello, async task 1!"]
  },
  {
    name            = "sequential task 1"
    description     = "print string from the sequential task 1"
    is_async        = false
    is_sudo         = false
    is_print_output = true
    exec            = ["echo", "hello, sequential task 1!"]
  }
]
