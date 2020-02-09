<script>
  import { api } from "../api.js";

  async function handleSubmit() {
    const payload = {
      maxFileSize: maxFileSize,
      totalQuota: totalQuota
    };
    const response = await api.post("/api/invite", JSON.stringify(payload));
    const data = await response.json();
    console.log(data);
  }
</script>

<main>
  <h2 class="subtitle has-text-centered">Change your password</h2>
  <label class="label">New password:</label>
  <p class="control has-addons">
    <input
      class="input is-expanded"
      type="password"
      placeholder="Your new password" />
  </p>
  <label class="label">Confirm password:</label>
  <p class="control has-addons has-text-centered">
    <input
      class="input"
      type="password"
      placeholder="Verify your new password" />
    <button on:click={handleSubmit} class="button is-primary ">
      Set new password
    </button>
  </p>
</main>
