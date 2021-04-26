<script>
	import LabledTextInput from '../labled_text_input/labled_text_input.svelte';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Password', 'val': ''}];
	const submit = () => {};
	let form;
</script>
<style src='./login_fields.scss'>
</style>

<form id='loginForm' bind:this={form}>
	<div id=labledInputsBlock>
		{#each inputs as input}
			<LabledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='loginButton' on:click={submit}>
			Login
		</button>
	</div>
</form>
