<script>
	import { createEventDispatcher, onDestroy } from 'svelte';
	export let modalWidth;

	const dispatch = createEventDispatcher();
	export const close = () => dispatch('close');

	let modal;
	

	const handle_keydown = e => {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			const nodes = modal.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter(n => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	};

	const previously_focused = typeof document !== 'undefined' && document.activeElement;

	if (previously_focused) {
		onDestroy(() => {
			previously_focused.focus();
		});
	}
</script>
<style src='./sidebar.scss'>
</style>

<svelte:window on:keydown={handle_keydown}/>

<div class="modal-background" on:click={close}
		 bind:clientWidth={modalWidth}>
</div>

<div class="modal" role="dialog" aria-modal="true" bind:this={modal}>
	<!-- <slot></slot> -->
	<h1>CONTENT GO HERE</h1>
	<!-- svelte-ignore a11y-autofocus -->
	<button autofocus on:click={close}>close modal</button>
</div>
