<script>
  import "./scss-entrypoint.scss";
  import { files } from "./store";
  import { api } from "./api";
  import Index from "./views/index.svelte";

  async function handlePaste(event) {
    for (const item of event.clipboardData.items) {
      if (item.type.indexOf("image") === 0) {
        const formData = new FormData();
        formData.append("files", item.getAsFile());
        await api.upload(formData);
        files.set(await api.getFiles());
      }
    }
  }
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

<main on:paste|preventDefault={handlePaste}>
  <Index />
</main>
