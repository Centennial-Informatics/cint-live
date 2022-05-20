<script lang="ts">
	import { IDToken, TeamID, TeamInfoData, userInfo } from '$lib/data/stores/userInfo';
	import { onMount } from 'svelte';
	import jwt_decode from 'jwt-decode';
	import Login from '$lib/utils/networking/login';
	import { contestEnded } from '$lib/data/stores/currentTime';

	interface DecodedCredential {
		sub: string;
		name: string;
		email: string;
		picture: string;
	}
	let ref: HTMLElement;

	onMount(() => {
		// if (!$contestEnded) {
		window.google.accounts.id.initialize({
			client_id: import.meta.env.VITE_OAUTH_CLIENT_ID,
			callback: async (res) => {
				const decoded: DecodedCredential = jwt_decode(res.credential);
				const loginInfo = await Login(res.credential, decoded.name);
				if (loginInfo.error) {
					alert(loginInfo.error);
					window.google.accounts.id.prompt();
				} else {
					IDToken.set(loginInfo.token);
					TeamID.set(loginInfo.team_id);
					userInfo.set({
						ID: decoded.sub,
						name: decoded.name,
						email: decoded.email,
						picture: decoded.picture
					});
				}
			}, // @ts-expect-error itp_support is not yet in the type
			itp_support: true,
			cancel_on_tap_outside: false,
			auto_select: false
		});
		window.google.accounts.id.renderButton(ref, {
			type: 'standard'
		});
		if (!$IDToken) {
			window.google.accounts.id.prompt();
		}
		// }
	});
</script>

<div>
	{#if !$IDToken}
		<div bind:this={ref} />
	{:else}
		<div class="hover:cursor-pointer">
			<div class="flex flex-row items-center">
				<a href="/#register" class="flex items-center space-x-3">
					<img src={$userInfo.picture} alt="avatar" class="w-10 h-10 rounded-full" />
					<div class="py-15 flex flex-col">
						<div class="font-medium tracking-normal h-5 text-md text-gray-300">
							{$userInfo.name}
						</div>
						{#if $TeamInfoData.team}
							<div class="font-medium text-gray-400 tracking-normal text-sm">
								{$TeamInfoData.team.name}
							</div>
						{/if}
					</div>
				</a>
			</div>
		</div>
	{/if}
</div>
