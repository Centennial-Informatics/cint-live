<script lang="ts">
	import NavLink from './navLink.svelte';
	import Profile from './profile.svelte';
</script>

<nav class="fixed top-1 right-0 left-60 bg-gray-800 text-gray-300 p-6 invisible md:visible">
	<div class="flex space-x-10">
		<NavLink to="/" tabId="home">Home</NavLink>
		<NavLink to="/standings" tabId="standings">Scoreboard</NavLink>
		<div class="flex-grow" />
		<Profile />
	</div>
</nav>
