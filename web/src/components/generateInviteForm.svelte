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

  let maxFileSize = 0;
  let totalQuota = 0;
</script>

<style>
  input[type="range"] {
    width: 100%;
  }
</style>

<main>
  <label class="label">Max file size: {maxFileSize}MB</label>
  <input
    name="fileSize"
    class="is-primary"
    type="range"
    min="0"
    max="100"
    bind:value={maxFileSize} />
  <br />
  <label class="label">Total upload quota: {totalQuota}MB</label>
  <input
    name="totalSize"
    class="is-primary"
    type="range"
    min="0"
    max="1000"
    bind:value={totalQuota} />
  <p class="has-text-centered">
    <button on:click={handleSubmit} class="button is-primary ">
      Generate invite
    </button>
  </p>
</main>
