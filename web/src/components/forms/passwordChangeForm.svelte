<script>
  import { api } from "../../api";

  async function handleSubmit(event) {
    let password = event.target.password.value;
    let passwordConfirm = event.target.passwordConfirm.value;
    if (password !== passwordConfirm) {
      document.getElementById("result").innerHTML = "passwords don't match";
      return;
    }

    const payload = { password };
    const response = await api.post(
      "/api/password/change",
      JSON.stringify(payload)
    );
    const data = await response.json();
    console.log(data);
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
    <p id="result" />
  </form>

</main>
