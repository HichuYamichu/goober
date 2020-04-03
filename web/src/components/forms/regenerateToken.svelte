<script>
  import { api } from "../../api";
  import { user } from "../../store";

  let result = "";

  let userValue;

  const unsubscribe = user.subscribe(value => {
    userValue = value;
  });

  async function handleSubmit(event) {
    const res = await api.regenerateToken();
    user.set({ ...userValue, token: res.token });
  }
</script>

<main>
  <h2 class="subtitle has-text-centered">Change your token</h2>
  <form on:submit|preventDefault={handleSubmit}>

    <label class="label">Your token:</label>
    <p class="control has-addons has-text-centered">
      <input
        name="token"
        class="input"
        type="text"
        readonly
        value={$user.token} />
      <input
        type="submit"
        value="Regenerate token"
        class="button is-primary " />
    </p>
    <p class="has-text-centered">{result}</p>
  </form>

</main>
