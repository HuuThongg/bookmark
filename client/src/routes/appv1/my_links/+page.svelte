<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<div>my link</div>
<div>{JSON.stringify(data.folders)}</div>
