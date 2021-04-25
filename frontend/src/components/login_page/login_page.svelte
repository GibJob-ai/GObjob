<script>
	import CreateAccountFields from '../general/form/create_account_fields/create_account_fields.svelte';
</script>
<style src='login_page.scss'></style>
<div id=loginPageMain>
	<h1 id='loginTitle'>Create your account</h1>
	<div id='newHere'>
		<div id='newHereInner'>
			New here?
		</div>
		<a href='/register' target='_self'>Create an account</a>
	</div>
	<div id='loginMain'>
		<CreateAccountFields/>
	</div>
</div>
