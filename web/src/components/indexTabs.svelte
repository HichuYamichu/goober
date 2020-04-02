<script>
  import FilesTable from "./tables/files.svelte";
  import Panel from "./adminPanel.svelte";

  let activeTab = 0;

  function handleClick(idx) {
    const tabs = document.querySelectorAll("li");
    tabs[activeTab].classList.remove("is-active");
    tabs[idx].classList.add("is-active");
    activeTab = idx;
  }
</script>

<main>
  <div class="tabs is-centered is-medium">
    <ul>
      <li class="is-active">
        <a href="javascript:;" on:click|preventDefault={() => handleClick(0)}>
          Uploads
        </a>
      </li>
      <li>
        <a href="javascript:;" on:click|preventDefault={() => handleClick(1)}>
          Administration
        </a>
      </li>
    </ul>
  </div>
  <div>
    {#if activeTab == 0}
      <FilesTable />
    {:else if activeTab == 1}
      <Panel />
    {/if}
  </div>

</main>
