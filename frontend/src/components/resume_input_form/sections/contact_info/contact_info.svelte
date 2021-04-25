<script>
	import Section from '../section.svelte'
	let contactInfo = {
		phone:'',
		email:'',
		linkedIn:'',
		github:'',
		website:''
	}
</script>

<style src='contact_info.scss'>
</style>

<Section title='Contact Information'>
	<div class='inputs'>
		<h3>Phone Number</h3>
		<input type='text' bind:value={contactInfo.phone}/>
		<h3>Email Address</h3>
		<input type='text' bind:value={contactInfo.email}/>
		<h3>LinkedIn</h3>
		<input type='text' bind:value={contactInfo.linkedIn}/>
		<h3>Github</h3>
		<input type='text' bind:value={contactInfo.github}/>
		<h3>Website</h3>
		<input type='text' bind:value={contactInfo.website}/>
	</div>
</Section>
