<script>
  import TopBar from '../general/menu/topbar/topbar.svelte';
  $: register_link = '/register'
</script>
<style src='./homepage.scss'></style>
<div id='landing-page-main'>
	<TopBar/>
  <div id="landing-page-content">
    <div id="landing-page-horizontal-split">
      <div id="landing-page-text">
        <div id="we-believe-in-you">
          <div id="we-believe-in">
            We believe in
          </div>
          <div id="you">
            you
          </div>
        </div>
        <div id='jk-text'>
          jk; you're probably fucked even with GibJob, but we'll do our best
        </div>
        <div id='register-bar'>
          <input type="text"/>
          <a href='{register_link}' id="register-button" target="_blank">
            Register
          </a>
        </div>
      </div>
      <div id='demo-animation-placeholder'>
        demo animation coming later
      </div>
    </div>
  </div>
</div>
