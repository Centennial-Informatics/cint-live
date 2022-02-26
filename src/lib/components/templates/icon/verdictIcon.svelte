<script lang="ts">
	import type { VerdictStatus } from 'src/types/submission';
	import CheckCircle from 'svelte-material-icons/CheckCircle.svelte';
	import CloseCircle from 'svelte-material-icons/CloseCircle.svelte';
	import AlertCircle from 'svelte-material-icons/AlertCircle.svelte';
	import HelpCircle from 'svelte-material-icons/HelpCircle.svelte';
	import Loading from '$lib/components/templates/loading.svelte';
	import { Colors } from '$lib/data/constants/colors';
	import { BAD, ERROR, GOOD, PENDING, UNSUBMITTED } from '$lib/data/constants/status';

	export let status: VerdictStatus = UNSUBMITTED;
	export let defaultIcon: boolean = true;
	export let size: string = '36';
</script>

{#if status === GOOD}
	<span class="mr-2"><CheckCircle width={size} height={size} color={Colors.GOOD[50]} /></span>
{:else if status === BAD}
	<span class="mr-2"><CloseCircle width={size} height={size} color={Colors.BAD[50]} /></span>
{:else if status === ERROR}
	<span class="mr-2"><AlertCircle width={size} height={size} color={Colors.BAD[50]} /></span>
{:else if status === PENDING}
	<span class="mr-2"><Loading {size} color={Colors.WARNING[50]} /></span>
{:else if defaultIcon && status === UNSUBMITTED}
	<span class="mr-2"><HelpCircle width={size} height={size} color={Colors.NEUTRAL[0]} /></span>
{/if}
