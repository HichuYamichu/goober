<script>
  import { api } from "../../api";

  let result = "";

  async function handleSubmit(event) {
    const password = event.target.password.value;
    const passwordConfirm = event.target.passwordConfirm.value;
    if (password !== passwordConfirm) {
      result = "passwords don't match";
      return;
    }

    const res = await api.changePassword(password);
    result = res.message;
  }
</script>

<style>
  .mg {
    margin: 1em;
  }
</style>

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
        class="button is-primary mg" />
    </p>
    <p class="has-text-centered">{result}</p>
  </form>

</main>
