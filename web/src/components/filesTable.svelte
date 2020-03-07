<script>
  import { createEventDispatcher } from "svelte";
  import { user } from "../store";

  const dispatch = createEventDispatcher();

  function handleDelete(fileName) {
    dispatch("remove", { fileName });
  }

  export let files = [];

  let userValue;

  const unsubscribe = user.subscribe(user => {
    userValue = user;
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
      {#each files as file, i}
        <tr>
          <th>{i + 1}</th>
          <td>{file.name}</td>
          <td>{(file.size * 10e-5).toFixed(3)}MB</td>
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
              disabled={!userValue.admin}
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
