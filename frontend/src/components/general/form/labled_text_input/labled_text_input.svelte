<script>
	export let label;
	export let value;
	export let form;
</script>
<style src='./labled_text_input.scss'></style>

<div class='labeled-text-input'>
	<div class='label'>{label}</div>
	<input type='text' name={label.toLowerCase()} bind:value/>
</div>
