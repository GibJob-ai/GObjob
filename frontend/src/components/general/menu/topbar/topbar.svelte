<script>
  import GibjobLogo from '../../logo/logo.svelte'
</script>

<style src='./topbar.scss'></style>

<div id="topbar-container">
	<div id="topbar">
		<GibjobLogo/>
		<div class='topbar-group'>
			<div class="topbar-button">
				Overview
			</div>
			<div class="topbar-button">
				Demos
			</div>
			<div class="topbar-button">
				Requirements
			</div>
		</div>
		<div class='topbar-group'>
			<div class="topbar-button" id="signin">
				Sign in
			</div>
			<div class="topbar-button" id="signup">
				Sign up
			</div>
		</div>
	</div>
</div>
