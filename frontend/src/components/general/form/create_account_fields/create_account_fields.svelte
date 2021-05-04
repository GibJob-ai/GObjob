<script>
	import LabledTextInput from '../labled_text_input/labled_text_input.svelte';
	import {email} from '../../../../store/register.js';
	let username = '';
	let password = '';
	let email = `${$email}`; // copy val
	let inputs = [{'name':'Username', 'val': username},{'name':'Email', 'val': email},{'name':'Password', 'val': password}];
	email.set(''); // null out the store after taking the value so that if they reload the page they start over
	let form;

	import { mutation } from '@urql/svelte';


	const signupMutation = mutation({
			query: `
			mutation($email: String!, $password: String!, username: String!){
				signUp(
					email: $email,
					password: $password,
					username: $username
				){
					ok,
					error,
				}
			}
		`,
	});

	// const signUp = mutation(signupMutation);
	const submit = () => {
		signupMutation({
			email: email,
			password: password,
			username: username,
		}).then(result => {
			if(result.data){
				console.log(result.data);
			} else if (result.error){
				console.error(result.error);
			}
		});
	};

</script>
<style src='create_account_fields.scss'>
</style>

<form id='registerForm' bind:this={form} on:submit|preventDefault={submit}>
	<div id=labledInputsBlock>
		{#each inputs as input}
			<LabledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='registerButton'>
			Register
		</button>
	</div>
</form>
