# ttcli — Task Tracker CLI tool

This project is an implementation of the **[Task Tracker Challenge from roadmap.sh](https://roadmap.sh/projects/task-tracker)**.

**ttcli** is a fast, dependency-free command-line interface (CLI) tool designed for tracking and managing your daily tasks. Built entirely using the **Go core standard library**, it requires no external frameworks (like Cobra or Urfave), resulting in instant startup times and high portability.

---

## 🚀 Features

- **Zero Third-Party Dependencies:** Leverages only Go's built-in core libraries.
- **Task Lifecycle Management:** Add, update, delete, and track tasks dynamically.
- **Progress Tracking:** Easily move tasks through states (`todo`, `in-progress`, `done`).
- **Flexible List Filters:** Filter tasks by status and toggle between detailed long formats (`--long`/`-l`) or compact views (`--short`/`-s`).
- **Persistent Local Storage:** Automatically saves all your data locally into a structured `tasks.json` file.

---

## 🛠️ Architecture & Project Structure

The project follows a modular and clean architecture separating core logic from data processing:

```text
taskTrackerCLI/
├── main.go          # CLI entry point, intercepts os.Args and routes subcommands
├── go.mod           # Go module definition
├── tasks.json       # Generated local JSON database for task persistence
├── constants/       # Centralized enumerations and status definitions
├── types/           # Core structures (e.g., Task definitions)
├── utils/           # Helper packages for file I/O operations
└── commands/        # Isolated business logic handlers for every action
```

### Core Implementation Details

1. **`main.go` (Routing Layer):** Instead of bulky external frameworks, routing is implemented natively using Go's `os.Args` slicing. It validates input parameters via the built-in `slices` package, preventing invalid command executions before passing arguments to handlers.
2. **`commands/` (Logic Layer):** Contains distinct, isolated functions for each subcommand to maximize code readabilty and maintainability.
3. **`utils/` (Storage Layer):** Handles data marshaling and unmarshaling using `encoding/json` combined with the `os` core file system utilities to ensure safe concurrent-like local writes to `tasks.json`.

---

## 📦 Installation & Global Configuration

To use `ttcli` globally from any folder in your terminal just like regular terminal tools, follow these installation steps.

### Windows (PowerShell)
1. Ensure your Go SDK compiler is available, then navigate to your project directory and configure your explicit binary output path:
   ```powershell
   go env -w GOBIN=C:\Users\Murad\go\bin
   ```
2. Build and install the binary natively into your system:
   ```powershell
   go install .
   ```
3. Add the target path to your User System Environment Variables via PowerShell (Admin):
   ```powershell
   [Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "User") + ";C:\Users\Murad\go\bin", "User")
   ```
*Restart your PowerShell terminal after applying this to refresh your environment.*

### Linux & macOS
```bash
go build -o ttcli main.go
sudo mv ttcli /usr/local/bin/
```

---

## 💻 Usage & Commands Reference

Once installed, use the following syntax anywhere in your terminal: `ttcli <command> [arguments]`

### 1. Global Help Menu
Displays clear usage instructions and arguments layouts:
```bash
ttcli help
```

### 2. Add a New Task
Creates a fresh task initialized under the `todo` status flag:
```bash
ttcli add "Complete building Go core CLI project"
```

### 3. Update Existing Tasks
Modify the textual content description of a task specifying its distinct identifier:
```bash
ttcli update 1 "Complete building and document Go core CLI project"
```

### 4. Delete a Task
Removes a specific record permanently from the tracking ecosystem:
```bash
ttcli delete 1
```

### 5. Status Management
Transition your tasks dynamically between workflow lifecycle phases:
```bash
# Mark as In-Progress
ttcli mark-in-progress 1

# Mark as Completed
ttcli mark-done 1
```

### 6. Advanced Filtering & Listing
Output your tracking database to the terminal using specific layout conditions:
```bash
# List all tasks
ttcli list

# Filter by a specific status
ttcli list todo
ttcli list in-progress
ttcli list done

# Toggle Detailed/Long layout formats
ttcli list --long
ttcli list todo -l

# Toggle Compact/Short layout formats
ttcli list --short
```

---

## 🔒 License
Distributed under the MIT License. Feel free to fork and build upon it!