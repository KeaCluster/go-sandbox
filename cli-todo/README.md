# Cli-todo

## About

Small cli app to manage todo-tasks

### How

- golang

build then run?

```sh
go build -o cli-todo
```

then run

```sh
./cli-todo add "Finish a project for once"
```

#### Tests

Run tests with `go test ./... -v`

_Should_ run any `*_test.go` file inside the project's dir.

### Code

- Handler for each CRUD operation
- Test for testing
- Task.go
  - The model
- storage
  - Manages the storage inside `/tasks.json`

### Functions

- `AddTask()`
  - Adds a task into a `tasks.json` file
    - Validation
    - stuff
- `showTasks()`
  - Returns a list of tasks if they exist
- `CompleteTask()`
  - Marks the task as completed
  - It doesn't delete it so far after completion
- `DeleteTask()`
  - Returns the completed task
  - Deletes the task from the file
  - Returns the result

### Todo

- [x] Refactor functions
- [x] Functional question mark
