<script>
  import { user } from "../store";
  import { api } from "../api";

  let username = "";
  let password = "";
  let result = "";

  async function handleLogin() {
    const data = await api.login(username, password);
    if (data.user && data.token) {
      document.cookie = data.token;
      user.set(data.user);
    } else {
      result = data.message;
    }
  }

  async function handleRegister() {
    const data = await api.register(username, password);
    result = data.message;
  }

  async function handleKeydown(event) {
    if (event.key === "Enter") {
      await handleLogin();
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<main>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-4 is-offset-4">
          <h3 class="title has-text-black">Login</h3>
          <hr class="login-hr has-background-black" />
          <p class="subtitle has-text-black">Or register new account</p>
          <div class="box">
            <div class="field">
              <div class="control">
                <input
                  bind:value={username}
                  class="input is-large"
                  type="text"
                  placeholder="Your Name" />
              </div>
            </div>
            <div class="field">
              <div class="control">
                <input
                  class="input is-large"
                  bind:value={password}
                  type="password"
                  placeholder="Your Password" />
              </div>
            </div>
            <div class="field is-grouped is-grouped-centered">
              <div class="control">
                <button class="button is-dark is-large" on:click={handleLogin}>
                  Login
                </button>
              </div>
              <div class="control">
                <button
                  class="button is-dark is-large"
                  on:click={handleRegister}>
                  Register
                </button>
              </div>
            </div>
            <p>{result}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
