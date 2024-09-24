<script lang="ts">
  import { uploadEntries } from '$lib/index';

  let entry_val = '';
  let file: File | null = null;

  function handleFileChange(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      file = input.files[0];
    }
  }

  // Function to upload the file to the Go backend
  async function uploadCsv() {
    if (!file) {
      alert("Please select a file first.");
      return;
    }

     try {
          const data = await uploadEntries(file); // Use the uploadCsv function
          console.log("CSV Data:", data); // Print the CSV content
          alert(JSON.stringify(data, null, 2)); // Display JSON data as an alert
        } catch (error) {
          alert("Error uploading file. Please try again.");
        }
  }
</script>

<h1>Upload CSV</h1>
<!-- Simple File Upload Form -->
<div>
  <input type="file" accept=".csv" on:change={handleFileChange} />
  <button on:click={uploadCsv}>Upload CSV</button>
</div>