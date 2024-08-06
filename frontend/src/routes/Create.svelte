<script>
  import { createEntry } from '../stores';

  let value = '';

  const submit = async () => {
    await createEntry({ value });
    value = '';
  }
</script>

<h1>Create Entry</h1>
<input type="text" bind:value placeholder="Value">
<button on:click={submit}>Create Entry</button>
