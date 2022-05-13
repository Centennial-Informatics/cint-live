<script>
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	// import { REGISTRATION_FORM_LINK } from '$lib/data/constants/contest';
</script>

<Wrapper margin={false} padding={false}>
	<div class="py-8 space-y-4 flex flex-col justify-center items-center px-4 text-white font-bold">
		<!-- <div class="text-white text-xl p-4">
			<strong>If you haven't registered:</strong>
			<div>1. Click the <strong>Register</strong> button down below.</div>
			<div>
				2. Enter your preferred <strong>Google Email</strong> (can be a school account) and additional
				information.
			</div>
			<div>
				3. Reload this page and sign in through the popup or the sign-in button <strong
					>using the email you registered with</strong
				>.
			</div>
		</div>
		<a
			href={REGISTRATION_FORM_LINK}
			target="_blank"
			class="text-white font-bold cursor-pointer hover:bg-brand-light text-3xl bg-brand py-4 px-16 rounded-xl"
			>Register</a
		> -->
		Sign-in first using your Google Email.
	</div>
</Wrapper>
