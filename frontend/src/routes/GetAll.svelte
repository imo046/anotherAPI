<script>
  import { entries, fetchEntries } from '../stores';

  fetchEntries();
</script>

<h1>Entries</h1>
<ul>
  {#each $entries as entry}
    <li>ID: {entry.id}, Value: {entry.value}</li>
  {/each}
</ul>
