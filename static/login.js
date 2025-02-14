// handle login
document.addEventListener("DOMContentLoaded", function () {
  document
    .getElementById("login-form")
    .addEventListener("submit", async (e) => {
      e.preventDefault();
      const username = document.getElementById("username").value;
      const password = document.getElementById("password").value;

      const response = await fetch("/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        window.location.href = "/app";
        localStorage.setItem("auth", "true");
        document.getElementById("username").value = "";
        document.getElementById("password").value = "";
      } else {
        alert("Invalid credentials");
      }
    });
});
