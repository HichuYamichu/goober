<script>
  import { createEventDispatcher } from "svelte";
  import { files, page } from "../store";
  import { api } from "../api";
  import moment from "moment";

  async function handleDelete(id) {
    await api.deleteFile(id);
    files.set(await api.getFiles());
  }

  async function next() {
    page.update(n => n + 1);
    files.set(await api.getFiles());
  }

  async function prev() {
    page.update(n => n - 1);
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

<style>

</style>

<main>
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
          <td width="60%" class="fileName">
            <a href="/files/{file.id}" download>{file.name}</a>
          </td>
          <td>{formatBytes(file.size)}</td>
          <td>{moment.unix(file.createdAt).format('YYYY-MM-DD HH:mm:ss')}</td>
          <td>
            <a
              href="javascript:;"
              class="button is-small is-danger is-outlined"
              title="Delete album"
              on:click={() => handleDelete(file.id)}>
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
    <button class="pagination-previous is-pulled-left" on:click={prev}>
      Previous
    </button>
    <button class="pagination-next is-pulled-right is-primary" on:click={next}>
      Next
    </button>
  </nav>
</main>
