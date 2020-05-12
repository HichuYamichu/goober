<script>
  import { onMount } from "svelte";
  import { files } from "../store";
  import { api } from "../api";
  import FilesTable from "../components/fileTable.svelte";

  async function handleUpload(event) {
    const formData = new FormData();
    for (const file of event.target.files) {
      formData.append("files", file);
    }
    await api.upload(formData);
    files.set(await api.getFiles());
  }

  onMount(async () => {
    files.set(await api.getFiles());
  });
</script>

<style>
  .upload-box {
    padding-left: 10em !important;
    padding-right: 10em !important;
  }

  .no-pb-bot {
    padding-bottom: 0em !important;
  }
</style>

<main>
  <section class="section no-pb-bot">
    <div class="container has-text-centered">
      <div class="file is-medium is-primary is-boxed is-centered">
        <label class="file-label">
          <input
            class="file-input"
            type="file"
            name="files"
            multiple
            on:change={handleUpload} />
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
      <FilesTable />
    </div>
  </section>
</main>
