<script>
	import { SwitchOnCreateMode } from '$lib/utils/switchOnCreateMode';
	import { ToggleMenuBar } from '$lib/utils/toggleMenuBar';

	function handlePlusIconClick() {
		SwitchOnCreateMode();
	}
</script>

<svete:head>
	<script src="https://kit.fontawesome.com/d8939cb1a8.js" crossorigin="anonymous"></script>
</svete:head>
<nav>
	<div class="menu">
		<i
			class="las la-bars"
			on:click|preventDefault|stopPropagation={ToggleMenuBar}
			on:blur|preventDefault={() => {
				console.log('blur');
			}}
		/>
	</div>
	<div class="search">Search</div>
	<button on:click={handlePlusIconClick}>Icon</button>
</nav>

<style lang="scss">
	::-webkit-scrollbar {
		width: 0.5rem;
		height: 0.5rem;
	}

	nav {
		position: fixed;
		top: 0;
		left: 0;
		height: 4.5rem;
		width: 100%;
		display: flex;
		align-items: center;
		border-bottom: 0.1rem solid $border-color-regular;

		.menu {
			flex: 1;
			height: 100%;
			display: flex;
			align-items: center;
			padding-left: 1em;

			i {
				font-size: 3rem;
				color: $text-color-regular;
				cursor: pointer;
			}
		}

		.search {
			flex: 2;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;

			form {
				width: 100%;
				height: 3.5rem;
				border-radius: 0.3rem;
				border: 0.1rem solid $border-color-regular;
				display: flex;
				align-items: center;
				transition: all 200ms ease-in-out;

				input[type='search'] {
					width: 90%;
					height: 100%;
					border: none;
					outline: none;
					padding: 0 1em;
					font-family: 'Arial CE', sans-serif;
					border-radius: inherit;

					&::placeholder {
						font-family: 'Arial CE', sans-serif;
						color: $text-color-light;
					}
				}

				input[type='search']::-webkit-search-cancel-button {
					appearance: none;
				}

				.clear {
					width: 10%;
					height: 100%;
					display: flex;
					align-items: center;
					justify-content: center;
					background-color: white;
					padding: 0 1em;

					i {
						font-size: 1.5rem;
					}
				}

				&:focus-within {
					box-shadow:
						rgba(67, 71, 85, 0.27) 0px 0px 0.25em,
						rgba(90, 125, 188, 0.05) 0px 0.25em 1em;
				}
			}
		}

		.buttons {
			flex: 1;
			height: 100%;
			padding-right: 1em;
			display: flex;
			align-items: center;
			justify-content: flex-end;
			gap: 1em;

			.share {
				display: flex;
				align-items: center;
				justify-content: center;
				min-width: max-content;
				border: 0.1rem solid $blue;
				min-width: 8rem;
				background-color: $blue;
				padding: 0.4em;
				border-radius: 0.25rem;
				gap: 1em;
				cursor: pointer;
				transition: all 200ms linear;

				i {
					font-size: 1.8rem;
					color: white;
				}

				span {
					font-size: 1.3rem;
					font-family: 'Arial CE', sans-serif;
					color: white;
				}

				&:hover {
					box-shadow:
						rgba(0, 0, 0, 0.16) 0px 1px 4px,
						$blue 0px 0px 0px 2px;
				}
			}

			.disabled {
				opacity: 0.5;
				pointer-events: none;
			}

			i.la-plus {
				font-size: 3rem;
				cursor: pointer;

				&:hover {
					color: $blue;
				}
			}
		}
	}
</style>
