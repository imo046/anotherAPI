<script>
  import { Router, Route } from 'svelte-routing';
  import ListEntries from './+page.svelte';
</script>

<nav>
  <a href="/">List Entries</a>
</nav>
<nav>
  <a href="/upload">Upload</a>
</nav>
<nav>
  <a href="/delete">Delete</a>
</nav>
<slot />
