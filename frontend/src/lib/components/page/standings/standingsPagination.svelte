<script lang="ts">
	import ArrowButton from './arrowButton.svelte';

	export let itemsPerPage = 10;
	export let totalItems = 0;
	export let displayPage: (pageStart: number, pageEnd: number) => void;

	let page = 0;

	let pageStart = 0;
	let pageEnd = pageStart + itemsPerPage;

	displayPage(pageStart, pageEnd); // initialize pages

	let hasPrevPage = pageStart > 0;
	let hasNextPage = pageEnd < totalItems;

	function onNextPage() {
		if (pageEnd < totalItems) {
			page++;
			pageStart = page * itemsPerPage;
			pageEnd = pageStart + itemsPerPage;
			hasPrevPage = pageStart > 0;
			hasNextPage = pageEnd < totalItems;
			displayPage(pageStart, pageEnd);
		}
	}

	function onPrevPage() {
		if (page > 0) {
			page--;
			pageStart = page * itemsPerPage;
			pageEnd = pageStart + itemsPerPage;
			hasPrevPage = pageStart > 0;
			hasNextPage = pageEnd < totalItems;
			displayPage(pageStart, pageEnd);
		}
	}
</script>

<div class="flex flex-row w-full items-center space-x-2 justify-end">
	<div class="text-gray-200">
		{pageStart + 1} to {Math.min(pageEnd, totalItems)} of {totalItems}
	</div>
	<div class="flex">
		<ArrowButton direction="L" onClick={onPrevPage} disabled={!hasPrevPage} size="30" />
		<ArrowButton direction="R" onClick={onNextPage} disabled={!hasNextPage} size="30" />
	</div>
</div>
