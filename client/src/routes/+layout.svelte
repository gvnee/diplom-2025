<script lang="ts">
	import '$lib/styles/app.css';
	
	let { children } = $props();
</script>

<svelte:head>
	<link rel="preconnect" href="https://rsms.me/">
	<link rel="stylesheet" href="https://rsms.me/inter/inter.css">
</svelte:head>

<div class="bg-[#181926]">
	{@render children()}
</div>