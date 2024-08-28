<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchEntries, deleteEntry, deleteAllEntries, updateEntry } from '$lib/index';
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';

    let entries = [];
    let loading = true;
    let selected = new Set();
    let error = null;

      let columns = {
        id: false,
        data: true,
        actions: true,
      };

  // Fetch entries on component mount

  onMount(async () => {
    try {
      const res = await fetchEntries();
      entries = res ?? [];
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  });

    // Delete specific entry
  async function handleDelete(id) {
    try {
       await deleteEntry(id);
       entries = entries.filter(entry => entry.id !== id); // Remove from local list
       alert('Entry deleted successfully');
    } catch (err) {
       console.error('Failed to delete entry:', err);
    }
  }

  // Select or unselect an entry
  function toggleSelect(id) {
    if (selected.has(id)) {
      selected.delete(id);
    } else {
      selected.add(id);
    }
  }

    // Update entry
    async function handleEdit(id, entryValue) {
      try {
        await updateEntry(id, entryValue);
        alert('Entry updated successfully');
      } catch (err) {
        console.error('Failed to update entry:', err);
      }
    }


</script>

<h1>Entries</h1>
{#if loading}
  <p>Loading entries...</p>
{:else if error}
  <p>Error: {error}</p>
{:else if entries.length === 0}
  <p>No entries found in the database.</p>
{:else}
<!-- Column Visibility Controls -->
<div>
  <h3>Toggle Columns</h3>
  <label>
    <input type="checkbox" bind:checked={columns.id} />
    ID
  </label>
  <label>
    <input type="checkbox" bind:checked={columns.data} />
    Data
  </label>
  <label>
    <input type="checkbox" bind:checked={columns.actions} />
    Actions
  </label>
</div>
<Table hoverable={true}>
  <TableHead>
    <TableHeadCell>Select</TableHeadCell>
    {#if columns.id}
    <TableHeadCell>ID</TableHeadCell>
    {/if}
    <TableHeadCell>Data</TableHeadCell>
    <TableHeadCell>Actions</TableHeadCell>
    <TableHeadCell>
      <span class="sr-only">Edit</span>
    </TableHeadCell>
  </TableHead>
  <TableBody tableBodyClass="divide-y">
      {#each entries as entry (entry.id)}
        <TableBodyRow>
              <TableBodyCell>
                          <input
                            type="checkbox"
                            checked={selected.has(entry.id)}
                            on:change={() => toggleSelect(entry.id)}
                          />
              </TableBodyCell>
              {#if columns.id}
                <TableBodyCell>
                {entry.id}
                </TableBodyCell>
              {/if}
              <TableBodyCell>
                          <input
                            type="text"
                            bind:value={entry.entry_val}
                            on:blur={() => handleEdit(entry.id, entry.entry_val)}
                          />
              </TableBodyCell>
              <TableBodyCell>
              <button on:click={() => handleDelete(entry.id)}>Delete</button>
              </TableBodyCell>
              <TableBodyCell>
                <a href="/tables" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Edit</a>
              </TableBodyCell>
        </TableBodyRow>
      {/each}
  </TableBody>
</Table>
{/if}
