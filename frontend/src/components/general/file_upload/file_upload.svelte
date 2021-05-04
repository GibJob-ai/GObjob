<script>
	import ToggleList from '../form/toggle_list/toggle_list.svelte';

	let show_generated_resumes = true;
	let show_uploaded_resumes = true;
	let toggle_items = [
		{'title': 'Show generated resumes', 'on': show_generated_resumes},
		{'title': 'Show uploaded resumes', 'on': show_uploaded_resumes}
	];
</script>

<style src='file_upload.scss'>
</style>

<div class='file_upload'>
	<h1>Your resume PDFs</h1>
	<ToggleList bind:items={toggle_items}>
</div>
