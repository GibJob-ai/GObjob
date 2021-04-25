<script>
	import LabledTextInput from './labled_text_input/labled_text_input.svelte';
	import {email} from '../../../../store/register.js';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Email', 'val': $email},{'name':'Password', 'val': ''}];
	email.set(''); // null out the store after taking the value so that if they reload the page they start over
	const submit = () => {};
</script>
<style src='create_account_fields.scss'>
</style>

<div id=labledInputsBlock>
	{#each inputs as input}
		<LabledTextInput label={input.name} bind:value={input.val}/>
	{/each}
	<button id='registerButton' on:click={submit}>
		Register
	</button>
</div>
