<script lang="ts">
	import {
		foldersCut,
		foldersFound,
		linksFound,
		query,
		searchInputFocused,
		selectedFolders,
		selectedLinks
	} from '../../stores/stores';
	import Folder from './folder.svelte';
	import Link from './link.svelte';

	function closeSearchResults() {
		query.set('');
		searchInputFocused.set(false);
		foldersFound.set([]);
		linksFound.set([]);
		foldersCut.set([]);
		selectedLinks.set([]);
		selectedFolders.set([]);
	}
</script>

{#if $searchInputFocused && $query !== ''}
	<div class="search_results" on:click|preventDefault|stopPropagation={closeSearchResults}>
		{#if $foldersFound.length > 0 && $linksFound.length > 0}
			{#each $foldersFound as folder}
				<Folder on:click {folder} />
			{/each}
			{#each $linksFound as link}
				<Link {link} />
			{/each}
		{:else if $foldersFound.length > 0 && $linksFound.length < 1}
			{#each $foldersFound as folder}
				<Folder on:click {folder} />
			{/each}
		{:else if $foldersFound.length < 1 && $linksFound.length > 0}
			{#each $linksFound as link}
				<Link {link} />
			{/each}
		{:else if $foldersFound.length === 0 && $linksFound.length === 0}
			<span>Nothing was found!</span>
		{/if}
	</div>
{/if}

<div>Search Results</div>
