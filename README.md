![TODO CLI Icon](assets/todo-cli-icon.png)
![TODO CLI Demo](assets/todo-cli-demo.gif)

# TODO CLI

**A simple and lightweight command-line task manager written in Go.**

یک ابزار ساده و کاربردی برای مدیریت وظایف روزانه از طریق ترمینال، توسعه‌یافته با زبان برنامه‌نویسی Go.

---

## About The Project | درباره پروژه

TODO CLI یک برنامه خط فرمان (Command Line Interface) برای مدیریت وظایف روزانه است.

این پروژه با هدف تمرین مفاهیم زبان Go و آشنایی با توسعه نرم‌افزارهای CLI طراحی شده است.

The project focuses on practicing Go fundamentals, including:

- Structs and Methods
- Pointers and Receivers
- Slices and Arrays
- Error Handling
- Packages and Modules
- JSON Encoding and Decoding
- File I/O
- Unit Testing

## Features | امکانات

- **Add Tasks:** افزودن وظایف جدید
- **List Tasks:** نمایش تمام وظایف
- **Complete Tasks:** علامت‌گذاری وظایف انجام‌شده
- **Edit Tasks:** ویرایش عنوان وظایف
- **Delete Tasks:** حذف وظایف
- **Task Filtering:** فیلتر وظایف انجام‌شده و انجام‌نشده
- **JSON Storage:** ذخیره‌سازی دائمی اطلاعات
- **Automatic ID Management:** مدیریت شناسه وظایف

## Project Structure | ساختار پروژه

```text
todo-cli/
│
├── assets/
│   ├── todo-cli-icon.png
│   └── todo-cli-demo.gif
│
├── cmd/
│   └── todo/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   ├── todo.go
│   │   └── todo_test.go
│   │
│   ├── cli/
│   │   ├── cli.go
│   │   └── cli_test.go
│   │
│   ├── storage/
│   │   ├── json.go
│   │   └── json_test.go
│   │
│   └── task/
│       └── task.go
│
├── bin/
├── .gitignore
├── go.mod
└── README.md
```

### Directory Description

| Directory / File | Description |
|---|---|
| `cmd/todo` | Application entry point |
| `internal/app` | Core application logic |
| `internal/cli` | CLI command handling |
| `internal/storage` | JSON data storage |
| `internal/task` | Task data structure |
| `assets` | Project images and GIF |

## Requirements | پیش‌نیازها

- Go
- Terminal or PowerShell
- Git (optional)

## Installation | نصب و راه‌اندازی

Clone the repository:

```bash
git clone https://github.com/abh2020-B2/todo-cli.git
```

Navigate to the project directory:

```bash
cd todo-cli
```

Run the application:

```bash
go run ./cmd/todo help
```

## Usage | نحوه استفاده

تمام دستورات برنامه از طریق ترمینال قابل اجرا هستند.

### 1. Add a Task

افزودن یک وظیفه جدید:

```bash
go run ./cmd/todo add "Learn Go"
```

### 2. List All Tasks

نمایش تمام وظایف:

```bash
go run ./cmd/todo list
```

### 3. List Completed Tasks

نمایش وظایف انجام‌شده:

```bash
go run ./cmd/todo list --completed
```

### 4. List Pending Tasks

نمایش وظایف انجام‌نشده:

```bash
go run ./cmd/todo list --pending
```

### 5. Complete a Task

تکمیل یک وظیفه با استفاده از شناسه:

```bash
go run ./cmd/todo done 1
```

### 6. Edit a Task

ویرایش عنوان وظیفه:

```bash
go run ./cmd/todo edit 1 "Learn Advanced Go"
```

### 7. Delete a Task

حذف یک وظیفه:

```bash
go run ./cmd/todo delete 1
```

### 8. Display Help

مشاهده راهنمای دستورات:

```bash
go run ./cmd/todo help
```

## Data Storage | ذخیره‌سازی اطلاعات

The application uses JSON for persistent data storage.

وظایف پس از هر عملیات تغییردهنده در فایل JSON ذخیره می‌شوند و با بسته‌شدن برنامه از بین نمی‌روند.

Default storage location:

**Windows:**
```text
%USERPROFILE%\.todo-cli\tasks.json
```

**Linux:**
```text
~/.todo-cli/tasks.json
```

## Testing | تست پروژه

The project uses Go