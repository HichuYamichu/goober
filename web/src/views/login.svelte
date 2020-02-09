<script>
  import { navigateTo } from "svelte-router-spa";
  import { user } from "../store.js";
  import { api } from "../api.js";
  export let currentRoute;

  let action = "Login";
  let endpoint = "/api/login";
  let message = "Please login to proceed";
  if (currentRoute.namedParams.inviteID) {
    action = "Register";
    endpoint = "/api/register";
    message = "This invite link will be used up once you register";
  }

  async function handleSubmit(event) {
    const payload = {
      username: event.target.username.value,
      password: event.target.password.value
    };

    if (currentRoute.namedParams.inviteID) {
      payload.invite = currentRoute.namedParams.inviteID;
    }

    const response = await api.post(endpoint, JSON.stringify(payload));
    if (response.status !== 200) {
      const errorEl = document.getElementById("error");
      errorEl.style.color = "red";
      errorEl.innerHTML = "Failed to authenticate";
      return;
    }

    const data = await response.json();
    document.cookie = data.token;
    user.set(data.user);
    navigateTo("/");
  }
</script>

<main>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-4 is-offset-4">
          <h3 class="title has-text-black">{action}</h3>
          <hr class="login-hr has-background-black" />
          <p class="subtitle has-text-black">{ message }</p>
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
                value={action} />
              <p id="error" />
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
