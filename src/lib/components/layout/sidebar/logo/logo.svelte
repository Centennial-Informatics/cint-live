<script lang="ts">
	export let imgUrl: string;
	export let imgSrc: string;
</script>

<div class="space-y-3 px-5">
	<div
		class="pt-3 text-3xl font-semibold text-black no-underline tracking-tighter flex flex-row items-center space-x-2 justify-center"
	>
		<a href={imgUrl}>
			<img src={imgSrc} alt="contest-logo" class="h-18 py-15 object-contain" />
		</a>
	</div>
</div>
