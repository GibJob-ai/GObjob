<script>
	import LabeledToggle from '../labeled_toggle/labeled_toggle.svelte';
	export let items;
</script>
<style src='toggle_list.scss'>
</style>
<div class='toggle_list'>
	{#each items as item}
		<LabeledToggle title="{item.title}" bind:on={item.on}/>
	{/each}
</div>
