<script>
  import { user } from "../store";
  import { api } from "../api";

  async function handleSubmit(event) {
    const payload = {
      username: event.target.username.value,
      password: event.target.password.value
    };

    const response = await api.post("/api/login", JSON.stringify(payload));
    const data = await response.json();

    if (response.status !== 200) {
      const errorEl = document.getElementById("error");
      console.log(data);
      errorEl.style.color = "red";
      errorEl.innerHTML = data.message;
      return;
    }

    document.cookie = data.token;
    user.set(data.user);
  }
</script>

<main>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-4 is-offset-4">
          <h3 class="title has-text-black">Login</h3>
          <hr class="login-hr has-background-black" />
          <p class="subtitle has-text-black">Please login to proceed</p>
          <div class="box">
            <form on:submit|preventDefault={handleSubmit}>
              <div class="field">
                <div class="control">
                  <input
                    class="input is-large"
                    type="text"
                    name="username"
                    placeholder="Your Name" />
                </div>
              </div>
              <div class="field">
                <div class="control">
                  <input
                    class="input is-large"
                    type="password"
                    name="password"
                    placeholder="Your Password" />
                </div>
              </div>
              <input
                type="submit"
                class="button is-block is-link is-large is-fullwidth"
                value="Login" />
              <p id="error" />
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
