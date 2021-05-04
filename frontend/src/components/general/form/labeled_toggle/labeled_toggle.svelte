<script>
	import Toggle from '../toggle/toggle.svelte';
	export let on;
	export let title;
</script>
<style src='labeled_toggle.scss'>
</style>
<div class='labeled_toggle'>
	<div class='title'>{title}</div>
	<Toggle bind:on/>
</div>
