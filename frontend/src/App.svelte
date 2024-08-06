<script>
  import { Router, Route } from 'svelte-routing';
  import CreateEntry from './routes/Create.svelte';
  import ListEntries from './routes/GetAll.svelte';
  import DeleteEntry from './routes/Delete.svelte';
</script>

<Router>
  <nav>
    <a href="/">List Entries</a>
    <a href="/create">Create Entry</a>
    <a href="/delete">Delete Entry</a>
  </nav>
  <Route path="/" component={ListEntries} />
  <Route path="/create" component={CreateEntry} />
  <Route path="/delete" component={DeleteEntry} />
</Router>
