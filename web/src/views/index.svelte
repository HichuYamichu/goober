<script>
  import Tabs from "../components/tabs.svelte";
  import { api } from "../api";

  async function handleSubmit(event) {
    const formData = new FormData(event.target);
    const response = await api.post("/api/upload", formData);
    const data = await response.json();
  }
</script>

<style>
  .upload-box {
    padding-left: 10em !important;
    padding-right: 10em !important;
  }
</style>

<main>

  <section class="section">
    <div class="container has-text-centered">
      <div class="file is-medium is-primary is-boxed is-centered">
        <label class="file-label">
          <input class="file-input" type="file" name="files" multiple />
          <span class="file-cta upload-box">
            <span class="file-icon">
              <i class="fas fa-upload" />
            </span>
            <span class="file-label">Upload</span>
          </span>
        </label>
      </div>
    </div>
  </section>
  <section class="section">
    <div class="container">
      <Tabs />
    </div>
  </section>
</main>
