<script>
  import { navigateTo } from "svelte-router-spa";
  import { user } from "../store.js";

  async function handleSubmit(event) {
    const payload = {
      username: event.target.username.value,
      password: event.target.password.value
    };

    const response = await fetch("/api/login", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    });

    if (response.status !== 200) {
      const errorEl = document.getElementById("error");
      errorEl.style.color = "red";
      errorEl.innerHTML = "Failed to authenticate";
      return;
    }

    const data = await response.json();
    document.cookie = data.token;
    user.update(user => {
      user.username = data.user.username;
      user.quota = data.user.quota;
      user.admin = data.user.admin;
    });
    console.log(data.user);
    navigateTo("/");
  }
</script>

<main>
  <form
    on:submit|preventDefault={handleSubmit}
    class="loginForm"
    action="/login"
    method="post">
    <input type="text" name="username" id="username" />
    <br />
    <input type="password" name="password" id="password" />
    <p>
      <input type="submit" value="login" />
    </p>
    <p id="error" />
  </form>
</main>
