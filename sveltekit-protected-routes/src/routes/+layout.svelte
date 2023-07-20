<script lang="ts">
	import './globals.css';
	import { page } from "$app/stores";
</script>

<main>
	{#if $page.route.id !== '/login' && $page.data.user}
		<div class="container mx-auto navbar bg-base-100">
			<div class="flex-1">
				<p class="normal-case text-xl">Protected Routes Practice</p>
			</div>
			<div class="flex-none">
				<a  class="btn btn-secondary" data-sveltekit-preload-data="tap" href="/logout">Logout ({$page.data.user.name})</a>
			</div>
		</div>
	{/if}
	<slot />
</main>
