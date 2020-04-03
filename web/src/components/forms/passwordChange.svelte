<script>
  import { api } from "../../api";

  async function handleSubmit(event) {
    const resultEl = document.querySelector("#result");

    const password = event.target.password.value;
    const passwordConfirm = event.target.passwordConfirm.value;
    if (password !== passwordConfirm) {
      resultEl.innerHTML = "passwords don't match";
      return;
    }

    const res = await api.changePassword(password);
    resultEl.innerHTML = res.message;
  }
</script>

<main>
  <h2 class="subtitle has-text-centered">Change your password</h2>
  <form on:submit|preventDefault={handleSubmit}>
    <label class="label">New password:</label>
    <p class="control has-addons">
      <input
        name="password"
        class="input is-expanded"
        type="password"
        placeholder="Your new password" />
    </p>
    <label class="label">Confirm password:</label>
    <p class="control has-addons has-text-centered">
      <input
        name="passwordConfirm"
        class="input"
        type="password"
        placeholder="Verify your new password" />
      <input
        type="submit"
        value="Set new password"
        class="button is-primary " />
    </p>
    <p id="result" class="has-text-centered" />
  </form>

</main>
