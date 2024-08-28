<script lang="ts">
  import { createEntry } from '$lib/index';

  let entry_val = '';

  async function handleSubmit() {
    if (entry_val) {
      try {
        const response = await createEntry({ entry_val });
        console.log('Entry added:', response);
        // Reset the form
        entry_val = '';
      } catch (error) {
        console.error('Error adding entry:', error);
      }
    } else {
      alert('Please provide entry value');
    }
  }
</script>

<h1>Create Entry</h1>
<form method="POST" on:submit|preventDefault={handleSubmit}>
  <label for="value">Value:</label>
  <textarea id="text" bind:value={entry_val} required></textarea>

  <button type="submit">Add Entry</button>
</form>