<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
	import { page } from '$app/stores';
	afterNavigate(() => {
		goto(`${$page.url.origin}/appv1/my_links`);
	});
</script>

<div class="home" />

<style lang="scss">
	.home {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: white;
	}
</style>
