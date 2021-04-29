<script>
	import { quintOut } from 'svelte/easing';
	import { crossfade } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	let isOpen = false;
</script>

<style src='./test_page.scss'>
</style>

<nav class={isOpen ? 'open' : ''}>
	<button on:click={() => {isOpen = !isOpen}} aria-label="Menu">
		<svg id="icon" width="184" height="184" viewBox="0 0 184 184" fill="none" xmlns="http://www.w3.org/2000/svg">
			<g id="close">
				<path id="topLine" d="M36 54H148"/>
				<path id="bottomLine" d="M36 130H148"/>
			</g>
			<path id="middleBarCircle" d="M36 92H147.5C161 92 179.5 83.4 179.5 59C179.5 28.5 144 4 92 4C43.3989 4 4 43.3989 4 92C4 140.601 43.3989 180 92 180C140.601 180 180 140.601 180 92C180 43.3989 141.601 4 93 4"/>
		</svg>
	</button>
	<ul>
		{#each items as m}
			<li>{m}</li>;
		{/each}
	</ul>
</nav>

<!-- <div class='hamburger'> -->
		<!-- {#if thing1} -->
		<!-- [> {#each list1 as i (i)} <] -->
			<!-- <svg width="{hamburgerSize}px" height="{hamburgerSize}px" -->
						<!-- in:receive="{{key: 2}}" -->
						<!-- out:send="{{key: 2}}" -->
			<!-- >  -->
				<!-- <path  -->
						<!-- stroke="#000" -->
						<!-- id="svg_3" -->
						<!-- d="m{0+hamburgerPad},{0+hamburgerPad} l{hamburgerSize-2*hamburgerPad},{hamburgerSize-2*hamburgerPad}" -->
						<!-- opacity="NaN" -->
						<!-- fill-opacity="null" -->
						<!-- stroke-opacity="null" -->
						<!-- stroke-dasharray="null" -->
						<!-- stroke-width="14" -->
						<!-- fill="none" -->
						<!-- stroke-linecap="round"/> -->
			<!-- </svg> -->
		<!-- {/if} -->
		<!-- [> {/each} <] -->
		<!-- {#if !thing1} -->
			<!-- <svg width="{hamburgerSize}px" height="{hamburgerSize}px" -->
						<!-- in:receive="{{key: 2}}" -->
						<!-- out:send="{{key: 2}}" -->
			<!-- >  -->
			<!-- <path  -->
						<!-- stroke="#000" -->
						<!-- id="svg_3" -->
						<!-- d="m{0+hamburgerPad},{0+hamburgerPad} l{hamburgerSize-2*hamburgerPad},0" -->
						<!-- opacity="NaN" -->
						<!-- fill-opacity="null" -->
						<!-- stroke-opacity="null" -->
						<!-- stroke-dasharray="null" -->
						<!-- stroke-width="14" -->
						<!-- fill="none" -->
						<!-- stroke-linecap="round"/> -->
			<!-- </svg> -->
		<!-- {/if} -->
<!-- </div> -->

<!-- <div style='display: flex; flex-flow: row nowrap'> -->
	<!-- <div style='width:200px; height: 200px; background-color: white;'> -->
		<!-- {#if thing1} -->
			<!-- <div in:receive="{{key: 1}}" -->
					 <!-- out:send="{{key: 1}}" -->
					 <!-- style='width: 100px; height: 100px; background-color: orange'> -->
				<!-- hi -->
			<!-- </div> -->
		<!-- {/if} -->
	<!-- </div> -->

	<!-- <div style='width:200px; height: 200px; background-color: white;'> -->
		<!-- {#if !thing1} -->
			<!-- <div in:receive="{{key: 1}}" -->
					 <!-- out:send="{{key: 1}}" -->
					 <!-- style='width: 100px; height: 100px; background-color: orange'> -->
				<!-- hi -->
			<!-- </div> -->
		<!-- {/if} -->
	<!-- </div> -->
<!-- </div> -->

<!-- <button on:click={() => { -->
				<!-- thing1 = !thing1; -->
				<!-- if (thing1){ -->
					<!-- list2 = [1]; list1 = []; -->
				<!-- } -->
				<!-- else{ -->
					<!-- list1 = [1]; list2 = []; -->
				<!-- } -->
				<!-- }}>press</button> -->
