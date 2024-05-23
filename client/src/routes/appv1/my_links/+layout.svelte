<script>
	import CreateFolderForm from '$lib/components/addFolderForm.svelte';
	import AddLinkForm from '$lib/components/addLinkForm.svelte';
	import { createFolderMode, addLinkMode } from '../../../stores/stores';
	createFolderMode.set(true); // remove
	addLinkMode.set(true);
</script>

<!-- {#if $createFolderMode} -->
<!-- 	<CreateFolderForm /> -->
<!-- {/if} -->

{#if $addLinkMode}
	<AddLinkForm />
{/if}
<slot />
