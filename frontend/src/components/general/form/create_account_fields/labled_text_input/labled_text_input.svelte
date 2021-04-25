<script>
	export let label;
	export let value;
</script>
<style src='./labled_text_input.scss'></style>

<div class='labeled-text-input'>
	<div class='label'>{label}</div>
	<input type='text' bind:value/>
</div>
