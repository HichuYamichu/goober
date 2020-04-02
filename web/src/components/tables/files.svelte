<script>
  import { createEventDispatcher } from "svelte";
  import { user, files } from "../store";

  async function handleDelete(fileName) {
    await api.deleteFile(fileName);
    files.set(await api.getFiles());
  }

  function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return "0 Bytes";

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ["Bytes", "KB", "MB", "GB"];

    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
  }
</script>

<main>
  <table class="table is-fullwidth">
    <thead>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th>Action</th>
      </tr>
    </thead>
    <tfoot>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th>Action</th>
      </tr>
    </tfoot>
    <tbody>
      {#each $files as file, i}
        <tr>
          <th>{i + 1}</th>
          <td>{file.name}</td>
          <td>{formatBytes(file.size)}</td>
          <td>{new Date(file.createdAt).toLocaleDateString('pl-PL')}</td>
          <td>
            <button class="button is-small is-success is-inverted">
              <a class="is-small" download href="/api/download/{file.name}">
                <span class="icon is-small ">
                  <i class="fas fa-download" />
                </span>
              </a>
            </button>
            <button
              class="button is-small is-danger is-inverted"
              disabled={!$user.admin}
              on:click={() => handleDelete(file.name)}>
              <span class="icon is-small ">
                <i class="fas fa-times" />
              </span>
            </button>
          </td>
        </tr>
      {:else}
        <p>loading...</p>
      {/each}
    </tbody>
  </table>
</main>
