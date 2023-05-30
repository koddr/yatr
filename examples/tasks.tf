name        = "My tasks set"
description = "This is an tasks set"
tasks = [
  {
    "is_async" = true
    "commands" = [
      {
        "name"        = "print async command 1"
        "description" = "this is an async command 1"
        "is_sudo"     = false
        "exec"        = ["echo", "this is the command 1 from the async task 1"]
      },
      {
        "name"        = "print async command 2"
        "description" = "this is an async command 2"
        "is_sudo"     = false
        "exec"        = ["echo", "this is the command 2 from the async task 1"]
      },
      {
        "name"        = "print async command 3 with error"
        "description" = "this is an async command 3 with error"
        "is_sudo"     = false
        "exec"        = ["error", "some error"]
      }
    ]
  },
  {
    "is_async" = false
    "commands" = [
      {
        "name"        = "print command 1"
        "description" = "this is a command 1"
        "is_sudo"     = false
        "exec"        = ["echo", "this is the command 1 from the task 2"]
      },
      {
        "name"        = "print command 2"
        "description" = "this is a command 2"
        "is_sudo"     = false
        "exec"        = ["echo", "this is the command 2 from the task 2"]
      },
      {
        "name"        = "print command 3 with error"
        "description" = "this is a command 3 with error"
        "is_sudo"     = false
        "exec"        = ["error", "some error in command 3 from the task 2"]
      },
    ]
  }
]
