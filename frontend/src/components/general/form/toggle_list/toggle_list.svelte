<script>
	import LabledToggle from '../labled_toggle/labled_toggle.svelte';
	export let items;
</script>
<style src='toggle_list.scss'>
</style>
<div class='toggle_list'>
	{#each items as item}
		<LabledToggle title="{item.title}" bind:on={item.on}/>
	{/each}
</div>
