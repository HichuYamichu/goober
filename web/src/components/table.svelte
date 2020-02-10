<script>
  import { onMount } from "svelte";
  import { api } from "../api";

  let files = [];

  onMount(async () => {
    const response = await api.get("/api/status");
    files = await response.json();
  });
</script>

<main>
  <table class="table is-fullwidth">
    <thead>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th>Uploader</th>
      </tr>
    </thead>
    <tfoot>
      <tr>
        <th>#</th>
        <th>Filename</th>
        <th>Size</th>
        <th>Date added</th>
        <th>Uploader</th>
      </tr>
    </tfoot>
    <tbody>
      {#each files as file, i}
        <tr>
          <th>{i + 1}</th>
          <td>{file.name}</td>
          <td>{file.size}</td>
          <td>{file.createdAt}</td>
          <td>{file.owner}</td>
        </tr>
      {:else}
        <p>loading...</p>
      {/each}
    </tbody>
  </table>
</main>
