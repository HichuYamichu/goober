<script>
  import PasswordChangeForm from "./forms/passwordChange.svelte";
  import RegenerateTokenForm from "./forms/regenerateToken.svelte";
  import UsersTable from "./tables/users.svelte";
  import { user } from "../store";

  let activeTab = 0;

  function handleClick(idx) {
    const tabs = document.querySelectorAll("#panel a");
    tabs[activeTab].classList.remove("is-active");
    tabs[idx].classList.add("is-active");
    activeTab = idx;
  }

  const config = {
    Name: "Goober",
    DestinationType: "ImageUploader, FileUploader",
    RequestType: "POST",
    RequestURL: `http://${location.hostname}/api/files`,
    FileFormName: "files",
    Headers: {
      Authorization: $user.token
    },
    ResponseType: "Text",
    URL: "$json:files[0].url$",
    ThumbnailURL: "$json:files[0].url$"
  };
  const sharexBlob = new Blob([JSON.stringify(config)], {
    type: "application/octet-binary"
  });
  const shareX = URL.createObjectURL(sharexBlob);

  function handleLogout() {
    user.set({});
  }
</script>

<style>
  .adjusted {
    margin: auto;
    width: 70%;
  }

  .pd {
    padding: 3em;
  }
</style>

<main>
  <div class="columns" id="panel">
    <div class="column is-2">
      <aside class="menu">
        <p class="menu-label">General</p>
        <ul class="menu-list">
          <li>
            <a
              href="javascript:;"
              on:click|preventDefault={() => handleClick(0)}
              class="is-active">
              Change password
            </a>
          </li>
          <li>
            <a
              href="javascript:;"
              on:click|preventDefault={() => handleClick(1)}>
              Change token
            </a>
          </li>
          <li>
            <a
              href="javascript:;"
              on:click|preventDefault={() => handleClick(2)}>
              Sharex
            </a>
          </li>
          <li>
            <a
              href="javascript:;"
              on:click|preventDefault={() => handleClick(3)}>
              Logout
            </a>
          </li>
        </ul>
        {#if $user.admin}
          <p class="menu-label">Admin</p>
          <ul class="menu-list">
            <li>
              <a
                href="javascript:;"
                on:click|preventDefault={() => handleClick(4)}>
                Users
              </a>
            </li>
          </ul>
        {/if}
      </aside>
    </div>
    <div class="column">
      <div class="adjusted">
        {#if activeTab == 0}
          <PasswordChangeForm />
        {:else if activeTab == 1}
          <RegenerateTokenForm />
        {:else if activeTab == 2}
          <div class="has-text-centered pd">
            <a
              href={shareX}
              download={`${location.hostname}.sxcu`}
              class="button is-primary is-large">
              Get ShareX config
            </a>
          </div>
        {:else if activeTab == 3}
          <div class="has-text-centered pd">
            <button class="button is-primary is-large" on:click={handleLogout}>
              Logout
            </button>
          </div>
        {:else if activeTab == 4}
          <UsersTable />
        {/if}
      </div>
    </div>
  </div>
</main>
