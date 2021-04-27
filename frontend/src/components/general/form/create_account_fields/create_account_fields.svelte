<script>
	import LabledTextInput from '../labled_text_input/labled_text_input.svelte';
	import {email} from '../../../../store/register.js';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Email', 'val': $email},{'name':'Password', 'val': ''}];
	email.set(''); // null out the store after taking the value so that if they reload the page they start over
	let form;

	import { mutation } from 'svelte-apollo';
	import gql from 'graphql-tag'
	const signupQuery = gql`
		mutation{
			signUp(
					email: $email,
					password: $password,
					firstName: $firstName,
					lastName: $lastName
			){
				ok,
				error,
			}
		}
	`;

	const signUp = mutation(signupQuery);
	const submit = async () => {
			try{
					data = await signUp({ variables:
											 {email: "test1@test.com",
												password: "test",
												firstName: "test",
												lastName: "test"}})
					if(data.data){
							console.log(data.data)
					}
			} catch(error){
					console.error("graphql error: " + error)
			}
	};

</script>
<style src='create_account_fields.scss'>
</style>

<form id='registerForm' bind:this={form}>
	<div id=labledInputsBlock>
		{#each inputs as input}
			<LabledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='registerButton' on:click={submit} on:submit|preventDefault={()=>{}}>
			Register
		</button>
	</div>
</form>
