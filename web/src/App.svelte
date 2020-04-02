<script>
  import "./scss-entrypoint.scss";
  import { onMount } from "svelte";
  import { user, files } from "./store";
  import { api } from "./api";
  import Index from "./views/index.svelte";
  import Auth from "./views/auth.svelte";

  user.useSessionStorage();

  async function handlePaste(event) {
    for (const item of event.clipboardData.items) {
      if (item.type.indexOf("image") != -1) {
        const formData = new FormData();
        formData.append("files", item.getAsFile());
        await api.upload(formData);
      }
    }
  }

  onMount(async () => {
    files.set(await api.getFiles());
  });
</script>

<style>
  :global(a:hover) {
    text-decoration: none;
  }
</style>

<svelte:head>
  <script defer src="https://use.fontawesome.com/releases/v5.3.1/js/all.js">

  </script>
</svelte:head>

<main>
  {#if $user.username}
    <Index />
  {:else}
    <Auth />
  {/if}
</main>
