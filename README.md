# Go Todo CLI

This is a simple command-line interface (CLI) application for managing a todo list, written in Go. It supports adding, listing, and deleting todo items, and stores the data in a CSV file.

## Code Structure

`cmd/`

- `add.go`: Implements the `add` command.
- `delete.go`: Implements the `delete` command.
- `list.go`: Implements the `list` command.

`data/`

- `data.csv`: Stores the todo items.

`models/`

- `todo.go`: Defines the `TodoItem` struct.

`utils/`

- `csv_utils.go`: Contains utility functions for reading from and writing to the CSV file.

`main.go`: The entry point of the application.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/chetanr25/go_todo.git
   cd go_todo
   ```

2. Install the dependencies:

   ```sh
   go mod tidy
   ```

3. Build the application:
   ```sh
   go build -o todo
   ```

## Usage

The CLI supports the following commands:

### Add a Todo Item

Add a new todo item to the list.

```sh
./todo add "Your todo item title"
```

### List Todo Items

List all the todo items.

```sh
./todo list
```

### Delete a Todo Item

Delete a todo item by its ID.

```sh
./todo delete <id>
```
