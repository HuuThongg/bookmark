<script lang="ts">
	import { containsSpecialChars } from '$lib/utils/checkForSpecialCharacters.js';
	import { folderName, moveItemsMode } from '../../stores/stores';
	import { CreateFolder } from '$lib/utils/createFolder';
	import { page } from '$app/stores';
	function showCreateFolderForm() {}
	async function submitForm() {
		if ($moveItemsMode) {
			console.log('move items mode on');
		} else {
			await CreateFolder($folderName, $page.params.folder_id);
		}
	}
</script>

<div
	class="create-folder"
	on:click|preventDefault|stopPropagation={showCreateFolderForm}
	on:keydown
>
	<form
		action=""
		on:submit={() => {
			console.log('submitted');
		}}
	>
		<div class="top">
			<span>Create collection</span>
			<i class="las la-times" />
		</div>
		<div class="middle">
			<div
				class="input"
				class:outline-red={$folderName === '' || containsSpecialChars($folderName)}
			>
				<i class="las la-folder-plus" />
				<input
					type="text"
					id="create-folder-input"
					bind:value={$folderName}
					placeholder="Enter folder name..."
					autocomplete="off"
					spellcheck="false"
				/>
			</div>
		</div>
		<div class="errors">
			{#if $folderName === ''}
				<span>Folder name is required</span>
			{/if}
			{#if containsSpecialChars($folderName)}
				<span>Folder name must not contain special chracters</span>
			{/if}
		</div>
		<div class="bottom">
			<div class="buttons">
				<button
					on:click|preventDefault|stopPropagation|once={submitForm}
					type="submit"
					class:disabled={$folderName === '' || containsSpecialChars($folderName)}
				>
					<span>Create</span>
				</button>
				<button on:click|preventDefault|stopPropagation={showCreateFolderForm}
					><span>Cancel</span></button
				>
			</div>
		</div>
	</form>
</div>
<div>Hello</div>

<style lang="scss">
	.create-folder {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		z-index: 6000;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: rgba(0, 0, 0, 0.4);
		backdrop-filter: blur(8px);

		form {
			min-height: 20rem;
			min-width: 40rem;
			background-color: white;
			display: flex;
			flex-direction: column;
			box-shadow:
				rgba(0, 0, 0, 0.2) 0px 12px 28px 0px,
				rgba(0, 0, 0, 0.1) 0px 2px 4px 0px,
				rgba(255, 255, 255, 0.05) 0px 0px 0px 1px inset;
			padding: 2em;
			transform: scale(1);
			border-radius: 0.6rem;
			animation: zoomin 0.5s ease-in-out;
			gap: 1em;

			.top {
				height: 10%;
				display: flex;
				align-items: center;
				justify-content: space-between;

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.5rem;
					color: $text-color-regular;
				}

				i {
					font-size: 1.8rem;
					color: $text-color-regular;
					cursor: pointer;
					background-color: transparent;
					border-radius: 100vh;
					padding: 0.1em;
					transition: background-color 150ms ease-in-out;

					&:hover {
						background-color: $gray;
					}
				}
			}

			.middle {
				flex: 1;
				width: 100%;
				display: flex;
				align-items: center;

				.input {
					display: flex;
					align-items: center;
					width: 100%;
					border: 0.2rem solid $border-color-regular;
					height: 3.5rem;
					border: none;
					outline: 0.1rem solid $border-color-regular;
					border-radius: 0.3rem;

					i {
						font-size: 2rem;
					}

					// .folder-icon {
					// 	width: 10%;
					// 	height: 100%;
					// 	display: flex;
					// 	align-items: center;
					// 	justify-content: center;

					// 	svg {
					// 		path {
					// 			stroke: transparent;
					// 			fill: rgba(0, 0, 0, 0.5);
					// 		}
					// 	}
					// }

					input[type='text'] {
						width: 90%;
						border: none;
						outline: none;
						height: 100%;
						padding: 0 0.5em;
						font-family: 'Arial CE', sans-serif;

						&::placeholder {
							font-family: 'Arial CE', sans-serif;
						}
					}

					&:focus-within {
						outline-color: $yellow;
					}
				}

				.outline-red {
					outline-color: $form_error_red;

					&:hover {
						outline-color: $form_error_red;
					}

					&:focus-within {
						outline-color: $form_error_red;
					}
				}
			}

			.errors {
				display: flex;
				flex-direction: column;

				span {
					font-family: 'Arial CE', sans-serif;
					font-size: 1.1rem;
					color: $form_error_red;
				}
			}

			.bottom {
				flex: 1;
				width: 100%;
				display: flex;
				align-items: center;
				justify-content: flex-end;

				.buttons {
					display: flex;
					align-items: center;
					gap: 0.5em;

					button {
						height: 3.5rem;
						min-width: 7rem;
						border: none;
						background-color: $gray;
						cursor: pointer;
						display: flex;
						align-items: center;
						justify-content: center;
						// background-image: linear-gradient(to top, $red 50%, $gray 50%);
						// background-size: 100% 200%;
						// background-position: top;
						// transition: background-position 100ms ease-in-out;
						border-radius: 0.3rem;
						transition: all 200ms ease-in-out;

						span {
							font-family: 'Product Sans Medium', sans-serif;
							font-size: 1.3rem;
							color: $text-color-regular;
						}

						&:hover {
							filter: brightness(90%);
						}
					}

					button[type='submit'] {
						// background-color: $main-blue;
						// background-image: linear-gradient(to top, $green 50%, $main-blue 50%);
						background-color: $yellow;

						span {
							color: $text-color-regular;
							font-family: 'Arial CE', sans-serif;
						}

						&:hover {
							//background-color: $secondary-blue;
							//background-position: bottom;
							filter: brightness(90%);
						}
					}

					.disabled {
						opacity: 0.5;
						pointer-events: none;
					}
				}
			}

			@keyframes zoomin {
				0% {
					transform: translateY(-50px);
					opacity: 0;
				}

				100% {
					transform: translateY(0);
					opacity: 1;
				}
			}

			// @media screen and (max-width: 440px) {
			// 	min-width: 98%;
			// 	min-height: 30%;
			// }

			// @media screen and (min-width: 441px) and (max-width: 768px) {
			// 	min-width: 80%;
			// 	min-height: 30%;
			// }
		}

		// @media screen and (max-width: 440px) {
		// 	form {
		// 		min-width: 98%;
		// 		min-height: 30%;
		// 	}
		// }

		// @media screen and (min-width: 441px) and (max-width: 1200px) {
		// 	form {
		// 		min-width: 90%;
		// 		min-height: 30%;
		// 	}
		// }

		// @media screen and (min-width: 769px) and (max-width: 1200px) {
		// 	form {
		// 		min-width: 0%;
		// 		min-height: 30%;
		// 	}
		// }
	}
</style>
