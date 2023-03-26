<script lang="ts">
	import { eraseCookie } from '$lib/utils/cookies';

	function logout() {
		eraseCookie('jwt_login');
		window.location.reload();
	}
</script>

<button
	type="button"
	on:click={logout}
	class="bg-red-500 w-full rounded-xl py-4 text-white font-bold shadow-xl hover:bg-opacity-70"
	>Sign Out</button
>
