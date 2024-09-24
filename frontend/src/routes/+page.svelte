<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchEntries, deleteEntry, deleteAllEntries, updateEntry } from '$lib/index';
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';

    let entries = [];
    let columnNames = ["UuId","app_id","Answer time","Created","Elaborate","Elaborate timestamp end","Elaborate timestamp start","Greeting","Greeting timestamp end","Greeting timestamp start","Language medication",
    "Language medication timestamp end","Language medication timestamp start","Language1","Language1 timestamp end","Language1 timestamp start","Language2","Language2 timestamp end","Language2 timestamp start",
    "Name","Name ink", "Name ink timestamp end","Name ink timestamp start","Name timestamp end","Name timestamp start","Picture","Picture timestamp end","Picture timestamp start","Radio1","Radio1 timestamp start",
    "Radio2","Radio2 timestamp start","Retell","Retell timestamp end","Retell timestamp start","Slider1","Slider1 timestamp start","Slider2","Slider2 timestamp start","Slider3","Slider3 timestamp start",
    "Submission ID","Suggestion","Suggestion timestamp end","Suggestion timestamp start","TapRate","TapRate timestamp end","TapRate timestamp start","data_version","notification_datetime","session_id","study_id","Annotations"];

    let columnNameSlice = columnNames.slice(1);

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

  function handleValue(entry, columnName): any {
    return entry[columnName];
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
</div>
<Table hoverable={true}>
  <TableHead>
    <TableHeadCell>Select</TableHeadCell>
    {#if columns.id}
    <TableHeadCell>ID</TableHeadCell>
    {/if}
    {#each columnNameSlice as columnName}
    <TableHeadCell>{columnName}</TableHeadCell>
    {/each}
  </TableHead>
  <TableBody tableBodyClass="divide-y">
      {#each entries as entry}
        <TableBodyRow>
              <TableBodyCell>
                          <input
                            type="checkbox"
                            checked={selected.has(entry.UuId)}
                            on:change={() => toggleSelect(entry.UuId)}
                          />
              </TableBodyCell>
              {#if columns.id}
                <TableBodyCell>
                {entry.UuId}
                </TableBodyCell>
              {/if}
              {#each columnNameSlice as columnName}
              <TableBodyCell>
              {handleValue(entry, columnName)}
              </TableBodyCell>
              {/each}
              <TableBodyCell>
              <button on:click={() => handleDelete(entry.app_id)}>Delete</button>
              </TableBodyCell>
              <TableBodyCell>
                <a href="/tables" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Edit</a>
              </TableBodyCell>
        </TableBodyRow>
      {/each}
  </TableBody>
</Table>
{/if}
