# To-Do List Web App

## Overview

This is a **To-Do List Web Application** built with **HTML, CSS, JavaScript, Go (Gin framework), and SQLite**. It allows users to **add, mark complete, update, delete, and list tasks**. The app also includes **user authentication** with login and registration functionalities.

## Features

- **User Authentication** (Register, Login, Logout)
- **Task Management**
  - Add new tasks
  - Mark tasks as completed/incomplete
  - Update tasks
  - Delete tasks
  - Fetch and display tasks dynamically
- **Secure Routes** (Users must be logged in to manage tasks)
- **Cookie-based authentication**
- **Local storage support**
- **Basic styling using CSS**

## Tech Stack

- **Frontend:** HTML, CSS, JavaScript
- **Backend:** Go (Gin Framework)
- **Database:** SQLite
- **Authentication:** Cookies, LocalStorage

## Installation

### Prerequisites

Make sure you have **Go** installed on your system.

### Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/todo-app.git
   cd todo-app
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
  ```

3. **Run the server:**
   
   go run main.go
   ```

4. **Open the app in your browser:**

   ```bash
   http://localhost:8080
   ```

## API Endpoints

### Authentication

| Method | Endpoint  | Description              |
| ------ | --------- | ------------------------ |
| POST   | /register | Register a new user      |
| POST   | /login    | Login and get user token |
| POST   | /logout   | Logout user              |

### Tasks

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /todos      | Fetch all tasks |
| POST   | /tasks      | Add a new task  |
| PUT    | /tasks/:id  | Update a task   |
| DELETE | /tasks/:id  | Delete a task   |

## Usage

1. **Register** a new user.
2. **Login** to access the task dashboard.
3. **Add, update, complete, or delete tasks** as needed.
4. **Logout** to end the session.

## Screenshots

(Include relevant screenshots of the UI here)

## Contributing

1. Fork the repository
2. Create a new branch (\`git checkout -b feature-branch\`)
3. Commit your changes (\`git commit -m 'Add new feature'\`)
4. Push to the branch (\`git push origin feature-branch\`)
5. Create a Pull Request


