<script>
  import { onMount } from "svelte";
  import { api } from "../api";

  let users = [];

  async function handleUserActivation(id) {
    await api.activateUser(id);
    users = await api.getUsers();
  }

  async function handleUserDeletetion(id) {
    await api.deleteUser(id);
    users = await api.getUsers();
  }

  onMount(async () => {
    users = await api.getUsers();
  });
</script>

<main>
  <h2 class="subtitle has-text-centered">Users</h2>
  <table class="table is-fullwidth">
    <thead>
      <tr>
        <th>#</th>
        <th>Username</th>
        <th>Active</th>
        <th>Quota</th>
        <th>Action</th>
      </tr>
    </thead>
    <tfoot>
      <tr>
        <th>#</th>
        <th>Username</th>
        <th>Active</th>
        <th>Quota</th>
        <th>Action</th>
      </tr>
    </tfoot>
    <tbody>
      {#each users as user, i}
        <tr>
          <th>{i + 1}</th>
          <td>{user.username}</td>
          <td>{user.active}</td>
          <td>{(user.quota * 10e-5).toFixed(0)}MB</td>
          <td>
            <button
              class="button is-small is-success is-inverted"
              disabled={user.active}
              on:click={() => handleUserActivation(user.id)}>
              <span class="icon is-small ">
                <i class="fas fa-check" />
              </span>
            </button>
            <button
              class="button is-small is-danger is-inverted"
              on:click={() => handleUserDeletetion(user.id)}>
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
