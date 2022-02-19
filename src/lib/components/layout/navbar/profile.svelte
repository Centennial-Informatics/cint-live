<script>
	// @ts-nocheck
	import { IDToken, userInfo } from '$lib/data/stores/userInfo';
	import { onMount } from 'svelte';
	import jwt_decode from 'jwt-decode';
	import Login from '$lib/utils/networking/login';

	onMount(() => {
		window.google.accounts.id.initialize({
			client_id: '151372106954-hqa4baeucq2n1pajodj31h9e5r141mur.apps.googleusercontent.com',
			callback: async (res) => {
				const decoded = jwt_decode(res.credential);
				IDToken.set(await Login(res.credential));
				userInfo.set({
					ID: decoded.sub,
					name: decoded.name,
					email: decoded.email,
					picture: decoded.picture
				});
			},
			auto_select: true
		});
		if (!$IDToken) window.google.accounts.id.prompt();
	});
</script>

<svelte:head>
	<script src="https://accounts.google.com/gsi/client" async defer></script>
</svelte:head>

<div>
	{#if !$IDToken}
		<div class="g_id_signin" data-type="standard" />
	{:else}
		<div class="hover:cursor-pointer">
			<div class="flex flex-row items-center">
				<a href="/#register" class="flex items-center space-x-3">
					<img src={$userInfo.picture} alt="avatar" class="w-10 h-10 rounded-full" />
					<div class="py-15 flex flex-col">
						<div class="font-medium tracking-normal h-5 text-md text-gray-300">
							{$userInfo.name}
						</div>
						<div class="font-medium text-gray-400 tracking-normal text-sm">Team Name</div>
					</div>
				</a>
			</div>
		</div>
	{/if}
</div>
