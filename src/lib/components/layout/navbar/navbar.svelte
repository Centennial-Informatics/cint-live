<script lang="ts">
	import NavLink from './navLink.svelte';
	import Profile from './profile.svelte';
	let scrollY: number;
	function navOpacity(y: number): number {
		const threshold = 50; // fade out on scroll until this point
		return Math.floor(y >= threshold ? 0 : 100 * ((threshold - y) / threshold));
	}
</script>

<svelte:window bind:scrollY />

<nav
	class="fixed top-1 right-0 left-60 bg-gray-800 text-gray-300 p-6 invisible md:visible"
	style="opacity: {navOpacity(scrollY)}%;"
>
	<div class="flex space-x-10">
		<NavLink to="/" tabId="home">Home</NavLink>
		<NavLink to="/standings" tabId="standings">Scoreboard</NavLink>
		<div class="flex-grow" />
		<Profile />
	</div>
</nav>
