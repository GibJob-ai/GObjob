<script>
	import { Router } from "@roxi/routify";
	import { routes } from "../.routify/routes"; // this is a build time import

	// graphql stuff
	import { initClient, query, operationStore } from "@urql/svelte"

	initClient({
			url: 'http://localhost:3000/graphql'
	})

	const test = operationStore(`
	 {
			test
	 }
	`);

	query(test)

	test.subscribe(value => {
		if(value.fetching){
				console.log("loading");
		}
		else if(value.error){
				console.error("graphql error: " + value.error)
		}
		else{
			console.log(`graphql value: ${value.data.test}`);
		}
	});
</script>

<style src="./app.scss" global></style>

<Router {routes}/>

