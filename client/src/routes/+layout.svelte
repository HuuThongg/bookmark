<script lang="ts">
	import Checkmark from '$lib/components/checkmark.svelte';
	import Errors from '$lib/components/errors.svelte';
	import Loading from '$lib/components/loading.svelte';
	import Sucess from '$lib/components/sucess.svelte';
	import ThankYou from '$lib/components/thank-you.svelte';
	import { emptyFoldersCut } from '$lib/utils/cutFolders';
	import { emptyLinksCut } from '$lib/utils/cutLinks';
	import { hideContextMenu } from '$lib/utils/hideContextMenu';
	import { hideSelectShowCategoryMenu } from '$lib/utils/hideSelectSearchCategory';
	import { hideShowOptionsMenu } from '$lib/utils/hideShowOptionsMenu';
	import { hideMenuBar } from '$lib/utils/toggleMenuBar';
	import { hideProfileMenu } from '$lib/utils/toggleProfileMenu';
	import { createButtonToggled, errors, selectedFolders, selectedLinks } from '../stores/stores';

	function handleBodyClick() {
		hideMenuBar();
		hideProfileMenu();
		hideShowOptionsMenu();
		hideSelectShowCategoryMenu();
		createButtonToggled.set(false);
		selectedFolders.set([]);
		selectedLinks.set([]);
		hideContextMenu();
		emptyFoldersCut();
		emptyLinksCut();
	}
	function handleRightClickOnBody() {
		hideMenuBar();
		hideShowOptionsMenu();
		createButtonToggled.set(false);
	}
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://maxst.icons8.com/vue-static/landings/line-awesome/line-awesome/1.3.0/css/line-awesome.min.css"
	/>
</svelte:head>
<Checkmark />
<ThankYou />
<Sucess />
<!-- <Loading /> -->
{#if $errors.length > 0}
	<Errors />
{/if}

<div
	class="app"
	on:click={handleBodyClick}
	on:keydown
	on:contextmenu|preventDefault|stopPropagation={handleRightClickOnBody}
>
	<main>
		<slot />
	</main>
</div>

<style lang="scss" global>
	@import url('https://fonts.cdnfonts.com/css/google-sans');
	@import url('https://fonts.cdnfonts.com/css/arial');
	@import url('https://fonts.cdnfonts.com/css/arial-mt');

	.app {
		padding: 0;
		margin: 0;
		box-sizing: border-box;
		height: 100vh;
		width: 100vw;
		max-height: 100vh;

		main {
			position: fixed;
			top: 4.5rem;
			left: 0;
			height: calc(100% - 4.5rem);
			max-height: calc(100% - 4.5rem);
			width: 100vw;
		}
	}

	* {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
		-webkit-user-select: none; /* Safari */
		-ms-user-select: none; /* IE 10 and IE 11 */
		user-select: none; /* Standard syntax */
		line-height: 1.6;
		text-rendering: optimizeLegibility;
		-webkit-font-smoothing: antialiased;
		font-stretch: expanded;

		&::selection {
			background: $main-blue;
			color: white;
		}
	}

	html {
		font-size: 62.5%;

		xml {
			font-family: 'Product Sans Black', sans-serif !important;

			svg {
				font-family: 'Product Sans Black', sans-serif !important;

				path {
					font-family: 'Product Sans Black', sans-serif !important;
				}
			}
		}
	}

	/* width */
	::-webkit-scrollbar {
		width: 0.8rem;
		height: 0.5rem;
		display: none;
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}

	/* Track */
	::-webkit-scrollbar-track {
		background-color: transparent;
		background-color: #eeeeee;
		background-color: transparent;
	}

	/* Handle */
	::-webkit-scrollbar-thumb {
		background-color: #c8c6c6;

		&:hover {
			background-color: #748da6;
		}
	}

	::-webkit-scrollbar-track-piece {
		background-color: yellow;
		background-color: transparent;
	}
</style>
