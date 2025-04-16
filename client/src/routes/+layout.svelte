<script lang="ts">
	import 'iconify-icon';
	import '$lib/styles/app.css';
	import Header from '$lib/components/header.svelte';
	
	let { children } = $props();
</script>

<svelte:head>
	<link rel="preconnect" href="https://rsms.me/">
	<link rel="stylesheet" href="https://rsms.me/inter/inter.css">
</svelte:head>

<div>
	<Header />
	<main>
	{@render children()}
	</main>
</div>

<style>
	main {
		height: 100vh;
		width: 100%;
		background: var(--crust);
		padding-top: var(--header-height);
	}
</style>