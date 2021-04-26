<script>
	import LabledTextInput from '../labled_text_input/labled_text_input.svelte';
	import {email} from '../../../../store/register.js';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Email', 'val': $email},{'name':'Password', 'val': ''}];
	email.set(''); // null out the store after taking the value so that if they reload the page they start over
	const submit = () => {};
	let form;
</script>
<style src='create_account_fields.scss'>
</style>

<form id='registerForm' bind:this={form}>
	<div id=labledInputsBlock>
		{#each inputs as input}
			<LabledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='registerButton' on:click={submit}>
			Register
		</button>
	</div>
</form>
