<script lang="ts">
	import Footer from '$lib/components/layout/footer/footer.svelte';
	import TeamEditor from '$lib/components/page/home/teamEditor.svelte';
	import Header from '$lib/components/templates/page/header.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Paragraph from '$lib/components/templates/typography/paragraph.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { contestData, contestInfo } from '$lib/data/stores/contestData';
	import { currentPage } from '$lib/data/stores/currentPage';
	import { contestEnded } from '$lib/data/stores/currentTime';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!$currentPage) currentPage.set('home');
	});

	function formatDate(date: Date) {
		return date.toLocaleDateString('en-US', {
			weekday: 'long',
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}

	function formatTime(date: Date) {
		return date.toLocaleTimeString('en-US', {
			hour: 'numeric',
			minute: '2-digit'
		});
	}

	function formatStrings(startTime: string, stopTime: string) {
		const startDate = new Date(startTime);
		const stopDate = new Date(stopTime);
		const dateString =
			formatDate(startDate) +
			(stopDate.toDateString() !== startDate.toDateString() ? ' – ' + formatDate(stopDate) : '');
		const startTimeString = formatTime(startDate);
		const stopTimeString = formatTime(stopDate);
		return [dateString, startTimeString, stopTimeString];
	}

	let dateString: string, startTimeString: string, stopTimeString: string;

	contestData.subscribe((val) => {
		const [dateStr, startTimeStr, stopTimeStr] = formatStrings(val.StartTime, val.StopTime);
		dateString = dateStr;
		startTimeString = startTimeStr;
		stopTimeString = stopTimeStr;
	});
	let location = 'Centennial High School Cafeteria';
</script>

<Wrapper transparent>
	<Page>
		{#if $contestData.Info}
			<Header>
				<!-- <a
					class="text-xl font-bold text-white rounded-lg p-3 bg-brand hover:bg-brand-light"
					href="https://codeforces.com/contestInvitation/b259e34a22093ddbb25c365d8c92a4a7ca39d914"
				>
					Practice Standard problems
				</a>
				<a
					class="mt-4 text-xl font-bold text-white rounded-lg p-3 bg-alt hover:bg-alt-light"
					href="https://codeforces.com/contestInvitation/b6b614c34df0daf281d0de77f1634267d9336601"
				>
					Practice Advanced problems
				</a> -->
				<Title>{$contestInfo.Name}</Title>
				<Subtitle>{dateString}</Subtitle>
				<Subtitle>{startTimeString} – {stopTimeString}</Subtitle>
				<Subtitle>{location}</Subtitle>
			</Header>
			<Paragraph>{$contestInfo.Description}</Paragraph>
			<Paragraph>
				{#if !$contestEnded}
					Teams may have up to four participants. Each team member’s submission will score for the
					entire team. Enter a team code down below to join a team, or share your code to invite
					others to your team. Team names and members will be locked at the start of the
					competition.
				{:else}
					CInT 2022 has concluded. Standings will be revealed in the Auditorium and updated on this
					site later. Use the links above if you wish to make practice submissions to the contest
					problems. Thank you for attending and we hope to have you back next year!
				{/if}
			</Paragraph>
			{#if !$contestEnded}
				<div id="register">
					<TeamEditor />
				</div>
			{/if}
		{/if}
	</Page>
	<Footer />
</Wrapper>
