<script>
  import PasswordChangeForm from "./forms/passwordChangeForm.svelte";
  import GenerateInviteForm from "./forms/createUserForm.svelte";
  import { user } from "../store";

  let userValue;

  const unsubscribe = user.subscribe(user => {
    userValue = user;
  });

  let activeTab = 0;

  function handleClick(idx) {
    const tabs = document.querySelectorAll("#panel a");
    tabs[activeTab].classList.remove("is-active");
    tabs[idx].classList.add("is-active");
    activeTab = idx;
  }
</script>

<style>
  .adjusted {
    margin: auto;
    width: 70%;
  }
</style>

<main>
  <div class="columns" id="panel">
    <div class="column is-2">
      <aside class="menu">
        <p class="menu-label">General</p>
        <ul class="menu-list">
          <li>
            <a on:click|preventDefault={() => handleClick(0)} class="is-active">
              Change password
            </a>
          </li>
          <li>
            <a on:click|preventDefault={() => handleClick(1)}>Logout</a>
          </li>
        </ul>
        {#if userValue.admin}
          <p class="menu-label">Admin</p>
          <ul class="menu-list">
            <li>
              <a on:click|preventDefault={() => handleClick(2)}>
                Generate invite
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
          <div class="has-text-centered">
            <button class="button is-primary is-large">Logout</button>
          </div>
        {:else if activeTab == 2}
          <GenerateInviteForm />
        {/if}
      </div>
    </div>
  </div>
</main>
