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

	const startDate = new Date($contestData.StartTime);
	const stopDate = new Date($contestData.StopTime);

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

	let dateString =
		formatDate(startDate) +
		(stopDate.toDateString() !== startDate.toDateString() ? ' – ' + formatDate(stopDate) : '');
	let startTimeString = formatTime(startDate);
	let stopTimeString = formatTime(stopDate);

	let location = 'Centennial High School Cafeteria';
</script>

<Wrapper transparent>
	<Page>
		<Header>
			<Title>{$contestInfo.Name}</Title>
			<Subtitle>{dateString}</Subtitle>
			<Subtitle>{startTimeString} – {stopTimeString}</Subtitle>
			<Subtitle>{location}</Subtitle>
		</Header>
		<Paragraph
			>Note all website issues here <a
				href="https://docs.google.com/document/d/1JTN0ENdYsaMEsUhO96iPPwOQWm6ZQQ6IsmW4sdvperM/edit"
				>https://docs.google.com/document/d/1JTN0ENdYsaMEsUhO96iPPwOQWm6ZQQ6IsmW4sdvperM/edit</a
			></Paragraph
		>
		<Paragraph>{$contestInfo.Description}</Paragraph>
		<Paragraph>
			Teams may have up to four participants. Each team member’s submission will score for the
			entire team. Enter a team code down below to join a team, or share your code to invite others
			to your team. Team names and members will be locked at the start of the competition.
		</Paragraph>
		<div id="register">
			<TeamEditor />
		</div>
	</Page>
</Wrapper>
