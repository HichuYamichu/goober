<script>
  import { onMount } from "svelte";
  import { Navigate } from "svelte-router-spa";
  import { user } from "../store.js";

  let userValue;

  const unsubscribe = user.subscribe(user => {
    userValue = user;
  });

  async function handleSubmit(event) {
    const formData = new FormData(event.target);
    const response = await fetch("/api/upload", {
      method: "POST",
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${document.cookie}`
      },
      body: formData
    });
    const data = await response.json();
  }

  let files = [];

  onMount(async () => {
    const res = await fetch("/api/status", {
      headers: {
        Authorization: `Bearer ${document.cookie}`
      }
    });
    files = await res.json();
  });
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  <form
    on:submit|preventDefault={handleSubmit}
    class="uploadForm"
    action="/upload"
    method="POST"
    enctype="multipart/form-data">
    <input type="file" name="files" multiple />
    <button type="submit">Upload</button>
  </form>
  <Navigate to="login">link to login</Navigate>

  <div>
    {userValue.admin}
  </div>

  <div>
    <ul>
      {#each files as file}
        <li>{file}</li>
      {:else}
        <p>loading...</p>
      {/each}
    </ul>
  </div>
</main>
