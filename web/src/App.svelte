<script>
  import "./scss-entrypoint.scss";
  import { Router } from "svelte-router-spa";
  import { routes } from "./routes";
  import { user } from "./store";

  user.useLocalStorage();
</script>

<svelte:head>
  <script defer src="https://use.fontawesome.com/releases/v5.3.1/js/all.js">

  </script>
</svelte:head>

<style>
  :global(a:hover) {
		text-decoration: none;
	}
</style>

<Router {routes} />
