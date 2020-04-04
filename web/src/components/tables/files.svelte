<script>
  import { createEventDispatcher } from "svelte";
  import { user, files } from "../../store";
  import { api } from "../../api";

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
  <nav class="pagination" role="navigation" aria-label="pagination">
    <button class="pagination-previous is-pulled-left">Previous</button>
    <button class="pagination-next is-pulled-right is-primary">Next</button>
  </nav>
  <table class="table is-fullwidth is-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th />
      </tr>
    </thead>
    <tfoot>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th />
      </tr>
    </tfoot>
    <tbody>
      {#each $files as file, i}
        <tr>
          <th>{i + 1}</th>
          <td>
            <a href="/files/{file.name}" download>{file.name}</a>
          </td>
          <td>{formatBytes(file.size)}</td>
          <td>{file.createdAt}</td>
          <td>
            <a
              href="javascript:;"
              class="button is-small is-danger is-outlined"
              title="Delete album"
              on:click={() => handleDelete(file.name)}>
              <span class="icon is-small">
                <i class="fa fa-trash" />
              </span>
            </a>
          </td>
        </tr>
      {:else}
        <p>loading...</p>
      {/each}
    </tbody>
  </table>
  <nav class="pagination" role="navigation" aria-label="pagination">
    <button class="pagination-previous is-pulled-left">Previous</button>
    <button class="pagination-next is-pulled-right">Next</button>
  </nav>
</main>
