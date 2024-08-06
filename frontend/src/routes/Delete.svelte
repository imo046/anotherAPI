<script>
  import { deleteEntry, deleteAllEntries } from '../stores';

  let id = '';

  const deleteOne = async () => {
    await deleteEntry(id);
    id = '';
  }

  const deleteAll = async () => {
    await deleteAllEntries();
  }
</script>

<h1>Delete Entry</h1>
<input type="text" bind:value={id} placeholder="Entry ID">
<button on:click={deleteOne}>Delete Entry</button>
<button on:click={deleteAll}>Delete All Entries</button>
