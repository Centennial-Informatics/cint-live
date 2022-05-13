<script lang="ts">
	import TeamEditor from '$lib/components/page/home/teamEditor.svelte';
	import Header from '$lib/components/templates/page/header.svelte';
	import Page from '$lib/components/templates/page/page.svelte';
	import Wrapper from '$lib/components/templates/page/wrapper.svelte';
	import Paragraph from '$lib/components/templates/typography/paragraph.svelte';
	import Subtitle from '$lib/components/templates/typography/subtitle.svelte';
	import Title from '$lib/components/templates/typography/title.svelte';
	import { contestData, contestInfo } from '$lib/data/stores/contestData';
	import { currentPage } from '$lib/data/stores/currentPage';
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
				<a
					class="text-xl font-bold text-white rounded-lg p-3 bg-alt-darkest hover:bg-alt-dark"
					href="https://docs.google.com/document/d/1JTN0ENdYsaMEsUhO96iPPwOQWm6ZQQ6IsmW4sdvperM/edit"
				>
					Note all website issues <a
						class="text-alt-light hover:underline"
						href="https://docs.google.com/document/d/1JTN0ENdYsaMEsUhO96iPPwOQWm6ZQQ6IsmW4sdvperM/edit"
						>Here</a
					>
				</a>
				<Title>{$contestInfo.Name}</Title>
				<Subtitle>{dateString}</Subtitle>
				<Subtitle>{startTimeString} – {stopTimeString}</Subtitle>
				<Subtitle>{location}</Subtitle>
			</Header>

			<Paragraph>{$contestInfo.Description}</Paragraph>
			<Paragraph>
				Teams may have up to four participants. Each team member’s submission will score for the
				entire team. Enter a team code down below to join a team, or share your code to invite
				others to your team. Team names and members will be locked at the start of the competition.
			</Paragraph>
			<div id="register">
				<TeamEditor />
			</div>
		{/if}
	</Page>
</Wrapper>
