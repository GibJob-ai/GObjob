<script>
	import Toggle from '../toggle/toggle.svelte';
	export let on;
	export let title;
</script>
<style src='labled_toggle.scss'>
</style>
<div class='labled_toggle'>
	<div class='title'>{title}</div>
	<Toggle bind:on/>
</div>
