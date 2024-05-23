<script lang="ts">
	import { goto, afterNavigate } from '$app/navigation';
	afterNavigate(async () => {
		await goto('http://localhost:5173/accounts/sign_in');
	});
</script>
