<script lang="ts">
	import { showCheckMarkLottie } from '../../stores/stores';

	import checkmark from '$lib/gifs/checkmark.gif';
</script>

{#if $showCheckMarkLottie}
	<div class="wrapper">
		<div class="gif_container">
			<img src={checkmark} alt="checkmark gif" />
		</div>
	</div>
{/if}

<style lang="scss">
	.wrapper {
		position: fixed;
		top: 4.5rem;
		left: 0;
		width: 100vw;
		height: calc(100% - 4.5rem);
		z-index: 6000;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: white;

		.gif_container {
			height: 15rem;
			width: 15rem;
			display: flex;
			align-items: center;
			justify-content: center;

			img {
				object-fit: contain;
				max-inline-size: 100%;
			}
		}
	}
</style>
