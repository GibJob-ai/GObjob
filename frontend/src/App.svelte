<script>
	import { Router } from "@roxi/routify";
	import { routes } from "../.routify/routes"; // this is a build time import


	// graphql stuff
	// importing from core is a workaround because the good people at apollo
	// are stupid and use react in a frontend agnostic framework
	import { ApolloClient, InMemoryCache } from "@apollo/client/core";
	import { setClient, query } from "svelte-apollo";
	import gql from "graphql-tag";

	const client = new ApolloClient({
		uri: "http://localhost:3000/graphql",
		cache: new InMemoryCache(),
	})

	setClient(client);

	const TEST = gql`
	 {
			test
	 }
	`

	const value = query(TEST);
	console.log("value: " + value.data.test);

</script>

<style src="./app.scss" global></style>
<Router {routes}/>

