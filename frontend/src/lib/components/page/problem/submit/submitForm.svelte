<script lang="ts">
	import Page from '$lib/components/templates/page/page.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import { PENDING } from '$lib/data/constants/status';
	import { contestLanguages } from '$lib/data/stores/contestData';
	import { contestEnded } from '$lib/data/stores/currentTime';
	import { IDToken } from '$lib/data/stores/userInfo';
	import Submit from '$lib/utils/networking/submit';
	import Button from './button.svelte';
	import Editor from './editor.svelte';
	import LanguageSelect from './languageSelect.svelte';

	export let title = 'A. Problem';
	export let ID = 'A';
	export let submissionStatus: Element;
	export let status = 'Final';

	let langs = $contestLanguages;
	let selectedLanguage = langs[0].ID;
	let submissionText = '';

	async function handleSubmit(submission: string, file?: File) {
		if (file) {
			const ext = file.name.split('.').pop();
			if (ext !== 'cpp' && ext !== 'java' && ext !== 'py') {
				// TODO: File type invalid error
				console.log('File invalid');
				return;
			}
		} else if (submission.length > 0) {
			submissionStatus.scrollIntoView({ behavior: 'smooth' });
			status = PENDING;
			await Submit($IDToken, {
				file: file,
				submission: submission,
				problem: ID,
				language: selectedLanguage
			});
			submissionText = '';
		}
	}
	//TODO: no submission if contest over
</script>

<Page>
	<div class="text-center text-gray-100 text-2xl font-bold mt-24">
		<span class="font-normal text-gray-400">Submit solution to</span>
		{title}
	</div>
	<div class="w-full text-center pt-2">
		<Subtitle
			>For Java submissions, make sure your solution is a public class (not just class).</Subtitle
		>
		<Subtitle
			>Note that you may need to use the "long" data type in Java or "long long" in C++ if the
			problem or your solution involves large (> 100 million) numbers.</Subtitle
		>
	</div>
	<form
		on:submit|preventDefault={() => handleSubmit(submissionText)}
		class="flex flex-col items-center justify-center"
	>
		<div class="flex flex-row w-full mb-6 mt-5">
			<div class="w-full text-center">
				<div class="my-3 w-full overflow-y-auto shadow-lg rounded-xl bg-gray-600">
					<LanguageSelect {langs} bind:selectedLanguage />
					<Editor bind:submissionText />
				</div>
			</div>
		</div>
		{#if $IDToken && !$contestEnded}
			<Button
				disabled={status === PENDING || !$IDToken}
				onClick={() => {
					console.log('Submitting');
				}}>Submit</Button
			>
			<div class="text-gray-400 mt-4 text-sm text-center">
				Only the most recent submission will be counted. There are no penalties for wrong answers.
			</div>
		{:else}
			<div class="text-gray-400 mt-4 text-xl text-center">
				{$contestEnded ? 'Contest is over.' : 'Sign in to submit.'}
			</div>
		{/if}
	</form>
</Page>
