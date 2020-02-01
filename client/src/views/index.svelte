<script>
  import { Navigate } from "svelte-router-spa";
  import { token } from "../store.js";

  let tokenValue;

  const unsubscribe = token.subscribe(token => {
    tokenValue = `Bearer ${token}`;
  });

  async function handleSubmit(event) {
    const formData = new FormData(event.target);
    const response = await fetch("/api/upload", {
      method: "POST",
      headers: {
        Accept: "application/json",
        Authorization: tokenValue
      },
      body: formData
    });
    const data = await response.json();
    console.log(data);
  }
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
</main>
