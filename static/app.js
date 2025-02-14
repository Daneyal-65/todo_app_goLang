//  get all the todos
async function fetchTasks() {
  try {
    const response = await fetch("/todos");
    if (!response.ok) throw new Error("Failed to fetch tasks");

    const tasks = await response.json();
    const taskList = document.getElementById("task-list");
    taskList.innerHTML = "";

    tasks.forEach((task) => {
      console.log(task);
      const li = document.createElement("li");
      li.className = task.completed ? "completed" : "";
      li.innerHTML = `
        <span class ="text">${task.title} - ${task.description}</span>
        <div class="task-actions">
          <button onclick="toggleTaskCompletion(${
            task.ID
          }, ${!task.completed})">
            ${task.completed ? "Undo" : "Complete"}
          </button>
          <button class="update" onclick="updateTask(${
            task.ID
          })">Update</button>
          <button onclick="deleteTask(${task.ID})">Delete</button>
        </div>
      `;
      taskList.appendChild(li);
    });
  } catch (error) {
    console.error("Error fetching tasks:", error);
  }
}
// create the todos
document.getElementById("task-form").addEventListener("submit", async (e) => {
  // console.log("called");
  e.preventDefault();
  const title = document.getElementById("title").value;
  const description = document.getElementById("description").value;

  try {
    const response = await fetch("/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, description, completed: false }),
    });
    if (response.ok) {
      fetchTasks();
      document.getElementById("task-form").reset();
    } else {
      console.error("Failed to add task");
    }
  } catch (error) {
    console.error("Error adding task:", error);
  }
});
// toggleTaskCompletion
async function toggleTaskCompletion(taskId, completed) {
  try {
    const response = await fetch(`/tasks/${taskId}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ completed }),
    });

    if (response.ok) fetchTasks();
  } catch (error) {
    console.error("Error updating task:", error);
  }
}
// updateTask
async function updateTask(taskId) {
  const title = prompt("Enter new title:");
  const description = prompt("Enter new description:");

  if (!title || !description) return;

  try {
    const response = await fetch(`/tasks/${taskId}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, description }),
    });

    if (response.ok) fetchTasks();
  } catch (error) {
    console.error("Error updating task:", error);
  }
}
// delete task
async function deleteTask(taskId) {
  try {
    const response = await fetch(`/tasks/${taskId}`, { method: "DELETE" });
    if (response.ok) fetchTasks();
  } catch (error) {
    console.error("Error deleting task:", error);
  }
}
// handle logout
document.getElementById("logout-button").addEventListener("click", logout);
async function logout() {
  try {
    // Clear the authentication cookie by making a request to the backend
    const res = await fetch("/logout", { method: "POST" });
    localStorage.clear("auth");
    // Redirect to the login page
    alert("logout success !!!");
    window.location.href = "/";
  } catch (error) {
    console.error("Logout failed:", error);
  }
}
async function secureRoute() {
  if (!localStorage.getItem("auth")) {
    window.location.href = "/";
  }
}
document.addEventListener("DOMContentLoaded", fetchTasks);
document.addEventListener("DOMContentLoaded", secureRoute);
