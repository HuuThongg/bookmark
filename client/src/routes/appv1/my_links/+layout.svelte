<script>
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import CreateFolderForm from '$lib/components/addFolderForm.svelte';
	import AddLinkForm from '$lib/components/addLinkForm.svelte';
	import { onMount } from 'svelte';
	import {
		createFolderMode,
		addLinkMode,
		currentFolder,
		loading,
		folders,
		links
	} from '../../../stores/stores';
	import Navbar from '$lib/components/navbar.svelte';
	import New from '$lib/components/new.svelte';
	import ActionsMenu from '$lib/components/actionsMenu.svelte';
	import Loader from '$lib/components/loader.svelte';
	import Link from '$lib/components/link.svelte';
	import Folder from '$lib/components/folder.svelte';
	afterNavigate(async () => {
		console.log('afterNavigate');
		setCurrentFolder();
	});
	function setCurrentFolder() {
		if ($page.params.folder_id === undefined) {
			console.log('$page.params.folder_id', $page.params.folder_id);
			currentFolder.set('');
		} else {
			console.log('$page.params.folder_id 2 ', $page.params.folder_id);
			currentFolder.set($page.params.folder_id);
		}
	}
	console.log('currentFolder ', $currentFolder);
	console.log('c', $page.params.folder_id);
	onMount(() => {
		getCurrentFolderAfterNavigate();
	});
	function getCurrentFolderAfterNavigate() {
		console.log('onMount');
		setCurrentFolder();
		console.log('folder', $folders);
		console.log('links', $links);
	}

	// createFolderMode.set(true); // remove
	// addLinkMode.set(true);
	function handleLinksSectionContextMenu() {}
	function handleAddItemWhenParentCollectionIsEmpty() {}
</script>

{#if $createFolderMode}
	<CreateFolderForm />
{/if}

{#if $addLinkMode}
	<AddLinkForm />
{/if}
<a href="my_links/hello-world">to </a>
<New />
<div class="container">
	<ActionsMenu />
	<Navbar />
	<div
		class="links"
		id="links"
		on:contextmenu|preventDefault|stopPropagation={handleLinksSectionContextMenu}
		on:keyup
	>
		{#if $loading}
			<Loader />
		{:else if !$loading}
			{#if $folders.length > 0 && $links.length > 0}
				{#each $links as link}
					<Link {link} />
				{/each}
				{#each $folders as folder}
					<Folder on:click {folder} />
				{/each}
			{:else if $folders.length > 0 && $links.length < 1}
				{#each $folders as folder}
					<Folder on:click {folder} />
				{/each}
			{:else if $links.length > 0 && $folders.length < 1}
				{#each $links as link}
					<Link {link} />
				{/each}
			{:else if $folders.length === 0 && $links.length === 0}
				{#if $page.url.pathname}
					<div class="no_items_container">
						{#if $page.url.pathname === '/appv1/my_links/trash'}
							<span>No items in trash</span>
						{:else if $page.url.pathname === '/appv1/my_links/shared_with_me'}
							<p>Nothing has been shared with you yet!</p>
						{:else}
							<!-- <div
								class="button"
								id="addLinkOrCreateFolderBtn"
								on:click|preventDefault|stopPropagation={SwitchOnCreateMode}
								on:keyup
							>
								<span>New</span>
								<i class="las la-plus" />
							</div> -->
							<!-- <span
								class="add_link"
								on:click|preventDefault|stopPropagation={handleAddItemWhenParentCollectionIsEmpty}
								on:keyup>Click to add</span
							> -->
							<img {src} alt="no-items" />
							<div class="text">
								<p>This collection is empty...</p>
								<button
									on:click|preventDefault|stopPropagation={handleAddItemWhenParentCollectionIsEmpty}
								>
									<i class="las la-plus" />
									<span>Add link or create collection</span>
									<i class="las la-angle-down" />
								</button>
							</div>
						{/if}
					</div>
				{/if}
			{/if}
		{/if}
	</div>
	<slot />
</div>

<style lang="scss">
	.container {
		width: 100%;
		height: 100%;
		position: relative;
		//background-color: green;

		.links {
			position: absolute;
			top: 4.5rem;
			left: 0;
			width: 100%;
			height: calc(100% - 4.5rem);
			max-height: calc(100% - 4.5rem);
			padding: 1em;
			display: flex;
			gap: 1.5em;
			overflow: auto;
			flex-wrap: wrap;
			align-content: flex-start;

			.no_items_container {
				height: 100%;
				width: 100vw;
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;

				// span {
				// 	font-size: 1.3rem;
				// 	font-family: 'Product Sans Medium', sans-serif;
				// 	color: $text-color-regular-2;
				// 	transition: all 200ms ease-in-out;
				// }

				// span.add_link {
				// 	cursor: pointer;

				// 	&:hover {
				// 		text-decoration: underline;
				// 	}
				// }

				// .button {
				// 	height: 3.5rem;
				// 	width: 9rem;
				// 	display: flex;
				// 	align-items: center;
				// 	justify-content: center;
				// 	gap: 1em;
				// 	background-color: $blue;
				// 	cursor: pointer;
				// 	border-radius: 0.3rem;

				// 	span {
				// 		font-family: 'Arial CE', sans-serif;
				// 		color: white;
				// 		font-size: 1.3rem;
				// 	}

				// 	i {
				// 		font-size: 1.5rem;
				// 		background-color: white;
				// 		border-radius: 100vh;
				// 	}
				// }

				img {
					width: 30rem;
					object-fit: contain;
				}

				.text {
					display: flex;
					flex-direction: column;
					gap: 1em;
					align-items: center;
					justify-content: center;

					p {
						font-size: 1.3rem;
						font-family: 'Arial CE', sans-serif;
						color: $text-color-dropbox;
					}

					button {
						display: flex;
						align-items: center;
						gap: 0.5em;
						border: 0.1rem solid #ffea20;
						outline-color: transparent;
						background-color: #ffea20;
						min-width: max-content;
						padding: 0.5em 1em;
						cursor: pointer;
						border-radius: 0.2rem;

						i {
							font-size: 1.8rem;
							color: $text-color-dropbox;
						}

						span {
							font-size: 1.3rem;
							font-family: 'Arial CE', sans-serif;
							color: $text-color-dropbox;
						}

						&:hover {
							outline: 0.1rem dashed $border-color-regular;
							outline-offset: 0.2rem;
						}
					}
				}
			}
		}
	}
</style>
