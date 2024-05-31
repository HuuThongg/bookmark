<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import Breadcrumb from '$lib/components/breadcrumb.svelte';
	import { folders, links, loading } from '../../../stores/stores';
	import type { PageData } from './$types';
	export let data: PageData;

	$: data.folders ? folders.set(data.folders) : folders.set([]);
	$: data.links ? links.set(data.links) : links.set([]);
	afterNavigate(async () => {
		// currentCollectionMember.set({});

		// breadcrumbRoot.set('Home');

		loading.set(false);

		// query.set('');

		// hideContextMenu();

		// resetAncestorsOfCurrentFolder();

		//await getLinksAndFolders('');
	});
</script>

<svelte:head>
	<title>Home - Linkspace</title>
</svelte:head>
<Breadcrumb />
<!-- <div>my link</div> -->
<!-- <div>{JSON.stringify(data.folders)}</div> -->
