<script lang="ts">
  import { updateEntry } from '$lib/index';

  let id = '';
  let value = '';

  const submit = async () => {
    await updateEntry({ id, value });
    id = '';
    value = '';
  }
</script>

<h1>Update Entry</h1>
<input type="text" bind:value={id} placeholder="Entry ID">
<input type="text" bind:value={value} placeholder="New Value">
<button on:click={submit}>Update Entry</button>
