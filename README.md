# CLI Task Tracker

A simple and intuitive Command-Line Interface (CLI) application for tracking and managing your daily tasks. Built with Go and the [Cobra](https://github.com/spf13/cobra) CLI framework, this tool stores your tasks locally using a structured JSON format.

## Features

- **Quick Task Entry**: Add new tasks directly from your terminal.
- **Status Management**: Effortlessly transition tasks between `todo`, `in-progress`, and `done`.
- **Filtering**: List all tasks or filter them by their active status.
- **Persistent Storage**: Safely stores all data in a local `tasktracker.json` file as a standard JSON Array.

## Installation

Ensure you have [Go](https://go.dev/) installed on your machine.

1. Clone this repository:
   ```powershell
   git clone https://github.com/Akshat79-art/task-tracker-cli
   cd task-tracker-cli
   ```

2. Compile the binary:
   ```powershell
   go build -o task-cli.exe
   ```

## Usage

You can run the executable directly (`.\task-cli.exe`) or run it on the fly using Go (`go run main.go`).

### Managing Tasks

**Add a task:**
```powershell
task-cli add "Buy groceries"
```
*(New tasks are instantiated automatically with a `todo` status).*

**Update a task description:**
Provide the task ID followed by the new description.
```powershell
task-cli update 1 "Buy organic groceries"
```

**Delete a task:**
```powershell
task-cli delete 1
```

### Managing Statuses

Update the progress state of any task using its ID:
```powershell
task-cli updateInProgress 1
task-cli updateDone 1
task-cli updateToDo 1
```

### Viewing Tasks

**List all tasks:**
```powershell
task-cli list
```

**Filter list by status:**
```powershell
task-cli list todo
task-cli list in-progress
task-cli list done
```

## Data Structure

Data is maintained in `tasktracker.json` at the root of the project. The schema for a Task holds the following properties:
- `Id` (Integer): Unique incremental identifier.
- `Task` (String): Description of the task.
- `Status` (String): Current progress.
- `CreatedAt` (Timestamp)
- `UpdatedAt` (Timestamp)

## License
MIT License.

---
*Built with the help of [roadmap.sh - Task Tracker Project](https://roadmap.sh/projects/task-tracker).*