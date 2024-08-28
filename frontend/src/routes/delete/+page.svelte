<script lang="ts">
  import { deleteEntry, deleteAllEntries } from '$lib/index';

  let id: string = '';

  const deleteOne = async () => {
    if (id) {
      try {
        const response = await deleteEntry(id);
        console.log('Entry deleted:', response);
        id = ''; // Reset the ID field
      } catch (error) {
        console.error('Error deleting entry:', error);
      }
    } else {
      alert('Please provide an ID to delete');
    }
  }

  async function deleteAll() {
    if (confirm('Are you sure you want to delete all entries? This action cannot be undone.')) {
      try {
        const response = await deleteAllEntries();
        console.log('All entries deleted:', response);
      } catch (error) {
        console.error('Error deleting all entries:', error);
      }
    }
  }
</script>

<h1>Delete Entry</h1>
<form on:submit|preventDefault={deleteOne}>
  <label for="id">ID of Entry to Delete:</label>
  <input id="id" bind:value={id} required />

  <button type="submit">Delete Entry</button>
</form>

<button on:click={deleteAll} class="delete-all">Delete All Entries</button>
