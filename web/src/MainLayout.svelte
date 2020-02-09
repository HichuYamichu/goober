<script>
  import { Route } from "svelte-router-spa";
  export let currentRoute;
</script>

<style>
  .app {
    width: 100%;
    height: 100vh;
    overflow: hidden;
  }
</style>

<main class="app">
  <Route {currentRoute} />
</main>
