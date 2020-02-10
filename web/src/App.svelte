<script>
  import "./scss-entrypoint.scss";
  import Index from "./views/index.svelte";
  import Login from "./views/login.svelte";
  import { user } from "./store";

  user.useSessionStorage();

  let userValue;

  const unsubscribe = user.subscribe(user => {
    userValue = user;
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
  {#if userValue.username}
    <Index />
  {:else}
    <Login />
  {/if}
</main>
