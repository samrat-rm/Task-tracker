# Task Tracker CLI

A robust command-line application built in Golang for managing tasks. This tool enables users to add, update, delete, list, and manage the status of tasks, with all data stored persistently in a JSON file.

---

### Roadmap Project Challenge

This project was developed as part of the Task Tracker Challenge from Roadmap.sh, designed to enhance skills in building practical applications.
---

## Key Features

	•	Add Tasks: Quickly create a new task by specifying a description.
	•	Update Tasks: Edit the description or status of existing tasks.
	•	Delete Tasks: Permanently remove unwanted tasks from your list.
	•	Mark Task Status: Set tasks as “in-progress” or “done” to track their progress.
	•	List Tasks: Display all tasks, along with their descriptions and current statuses, in an organized format.

---

## Installation

To install and run the Task Tracker CLI, ensure you have Golang installed on your system. Then follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/samrat-rm/Task-tracker.git
   ```

2. Navigate to the project directory:

   ```bash
   cd Task-tracker
   ```

3. Build the application:

   ```bash
   go build -o task-cli
   ```

---

## Usage

Here’s how to use the various commands provided by the Task Tracker CLI:

### Add a Task

```bash
./task-cli add "Your task description"
```

### Update a Task

```bash
./task-cli update <task_id> "Updated task description"
```

### Delete a Task

```bash
./task-cli delete <task_id>
```

### Marking a Task as "in-progress" or "done"

```bash
./task-cli mark-in-progress <task_id>
./task-cli mark-done <task_id>
```

### Listing All Tasks

```bash
./task-cli list
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
