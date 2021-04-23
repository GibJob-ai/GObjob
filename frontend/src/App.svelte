<script>
	let test=1;

	import { Router } from "@roxi/routify";
	import { routes } from "../.routify/routes"; // this is a build time import
</script>

<style src="./app.scss" global></style>

<Router {routes}/>
