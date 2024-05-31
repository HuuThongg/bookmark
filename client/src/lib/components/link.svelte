<script lang="ts">
	import DefaultFavicon from '$lib/images/favicon.png';
	import {
		selectedLinks,
		selectedFolders,
		linksCut,
		controlKeyIsPress,
		currentCollectionMember
	} from './../../stores/stores';
	import type { Link } from '$lib/types/link';
	import { dragLink } from '../utils/dragLink';
	export let link: Partial<Link>;

	let selectedLink: Partial<Link>;

	let link_title: string | undefined = '';
	let link_id: string | undefined = '';
	let account_id: string | undefined = '';
	let folder_id: string | undefined = '';
	function handleClickOnCheckbox() {
		console.log('handleClickOnCheckBox');
	}
	function removeLinkFromSelectedLinks(link: Partial<Link>) {
		$selectedLinks = $selectedLinks.filter((li) => li.link_id != link.link_id);
	}

	function appendLinkToSelectedLinks(link: Partial<Link>) {
		$selectedLinks = [...$selectedLinks, link];
	}
	function handleClickOnLink() {
		console.log('handleClickOnLink');
	}
	function handleLinkContextMenu(e: MouseEvent) {
		console.log('handleLinkContextMenu');
	}
	function openLink() {
		console.log('openLink');
	}
</script>

<div
	class="link"
	class:link_cut={$linksCut.map((li) => li.link_id).includes(link.link_id)}
	data-linkid={link.link_id}
	data-linktitle={link.link_title}
	data-accountid={link.account_id}
	data-folderid={link.folder_id?.String}
	data-linkurl={link.link_url}
	on:contextmenu|preventDefault|stopPropagation={handleLinkContextMenu}
	on:click|preventDefault|stopPropagation={handleClickOnLink}
	on:keyup
	class:link-selected={$selectedLinks.map((li) => li.link_id).includes(link.link_id)}
	draggable="true"
	on:drag={dragLink}
>
	<div class="top">
		<div class="img-container">
			<img src={link.link_thumbnail} alt="card-hero" loading="lazy" />
		</div>
	</div>
	<div class="bottom">
		<div class="text-content">
			<div class="title">
				<span on:click|preventDefault|stopPropagation={openLink} on:keyup>
					{link.link_title}
				</span>
			</div>
			<div class="flavicon_and_menu">
				<div class="flavicon">
					<div class="flavicon-container">
						<img
							src={link.link_favicon !== '' ? link.link_favicon : DefaultFavicon}
							alt="flavicon"
							loading="lazy"
						/>
					</div>
					<div class="host">
						<span>{link.link_hostname}</span>
					</div>
				</div>
				<div class="menu">
					<div class="icon">
						<svg
							width="24px"
							height="24px"
							stroke-width="1.5"
							viewBox="0 0 24 24"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							color="#000000"
							class="more_svg"
							on:click|preventDefault|stopPropagation={handleLinkContextMenu}
							on:keyup
							><path
								d="M12 12.5a.5.5 0 100-1 .5.5 0 000 1zM12 18.5a.5.5 0 100-1 .5.5 0 000 1zM12 6.5a.5.5 0 100-1 .5.5 0 000 1z"
								fill="#000000"
								stroke="#000000"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</svg>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div
		class="check-box"
		on:click|preventDefault|stopPropagation={handleClickOnCheckbox}
		on:keyup
		data-linkid={link.link_id}
		data-linktitle={link.link_title}
		data-accountid={link.account_id}
		data-folderid={link.folder_id?.String}
	>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M5 13l4 4L19 7"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
	</div>
</div>

<style lang="scss">
	.link {
		height: 30rem;
		width: 30rem;
		border-radius: 0.3rem;
		border: 0.1rem solid $border-color-regular;
		background-color: white;
		transition: box-shadow 0.2s ease-in-out;
		position: relative;
		border-radius: 0.6rem;

		.top {
			height: 70%;
			width: 100%;

			.img-container {
				width: 100%;
				height: 100%;
				background-color: $gray;

				img {
					max-inline-size: 100%;
					height: 100%;
					width: 100%;
					object-fit: fill;
					filter: brightness(95%);
					transition: filter 200ms ease-in-out;
				}
			}
		}

		.bottom {
			height: calc(100% - 70%);
			min-height: calc(100% - 70%);
			width: 100%;

			.text-content {
				width: 100%;
				height: 100%;
				display: flex;
				flex-direction: column;

				.title {
					width: 100%;
					transition: height 200ms ease-in-out;
					display: flex;
					align-items: center;
					transition: height 0.4s linear;
					flex: 2;
					padding: 0 0.5em 0 0.5em;
					overflow: auto;

					span {
						font-family: 'Arial CE', sans-serif;
						font-size: 1.3rem;
						color: $text-color-regular;
						text-decoration: none;
						line-height: 1.6;
						text-decoration: underline;
						text-decoration-color: transparent;
						transition: text-decoration-color 0.2s linear;
						cursor: pointer;
						display: -webkit-box;
						-webkit-line-clamp: 2;
						-webkit-box-orient: vertical;
						overflow: hidden;
						text-overflow: ellipsis;
						text-transform: capitalize;
					}
				}

				.flavicon_and_menu {
					width: 100%;
					display: flex;
					align-items: center;
					justify-content: space-between;
					flex: 1;
					padding: 0 0.5em 0 0.5em;
					overflow: auto;

					.flavicon {
						display: flex;
						align-items: center;
						gap: 0.5em;
						width: 90%;
						height: 100%;

						.flavicon-container {
							height: 1.8rem;
							width: 1.8rem;
							border-radius: 100vh;
							display: flex;
							align-items: center;
							justify-content: center;

							img {
								max-inline-size: 100%;
								height: 100%;
								width: 100%;
								object-fit: cover;
								border-radius: 100vh;
								border-radius: inherit;
							}
						}

						.host span {
							font-size: 1.2rem;
							color: $text-color-medium;
							font-family: 'Arial CE', sans-serif;
							cursor: pointer;
						}
					}

					.menu {
						width: 10%;
						height: 100%;
						display: flex;
						align-items: center;
						justify-content: flex-end;

						.icon {
							width: 2.5rem;
							height: 2.5rem;
							display: flex;
							align-items: center;
							justify-content: flex-end;

							svg {
								cursor: pointer;

								path {
									stroke: $stroke;
									pointer-events: none;
								}
							}
						}
					}
				}
			}
		}

		.check-box {
			height: 1.7rem;
			width: 1.7rem;
			display: flex;
			align-items: center;
			justify-content: center;
			border: 0.15rem solid $text-color-regular;
			position: absolute;
			top: 1%;
			right: 1%;
			background-color: white;
			opacity: 0;
			cursor: pointer;
			border-radius: 100vh;
			transform: scale(1);
			box-shadow: rgba(0, 0, 0, 0.15) 0px 2px 8px;
			transition:
				opacity 0.2s linear,
				transform 0.2s linear;

			svg {
				visibility: hidden;
				pointer-events: none;

				path {
					stroke: $stroke;
					pointer-events: none;
				}
			}

			&:hover {
				transform: scale(1.1);

				svg {
					visibility: visible;
				}
			}
		}

		&:hover {
			box-shadow: 0 0.125rem 0.5rem 0 rgb(0 0 0 / 22%);

			.top {
				.img-container {
					img {
						filter: brightness(80%);
					}
				}
			}

			.bottom {
				//backdrop-filter: brightness(80%);

				.text-content {
					.title {
						span {
							text-decoration-color: $text-color-regular;
						}
					}

					.flavicon_and_menu {
						.flavicon {
							.host {
								text-decoration: underline;
								color: $text-color-regular;
								text-decoration-color: $text-color-regular;
							}
						}
					}
				}
			}

			.check-box {
				opacity: 1;
			}
		}
	}

	.link_cut {
		border-style: dashed;
		opacity: 0.5;
		pointer-events: none;
	}

	:global(.link-selected) {
		//box-shadow: rgba(3, 102, 214, 0.3) 0px 0px 0px 1px !important;
		background-color: rgb(214, 230, 255) !important;

		.top {
			background-color: rgb(214, 230, 255) !important;

			.img-container {
				img {
					filter: brightness(80%);
				}
			}
		}

		.bottom {
			background-color: rgb(214, 230, 255) !important;

			.text-content .title span {
				color: $main-blue !important;
				text-decoration-color: $main-blue !important;
			}

			.flavicon_and_menu {
				.flavicon {
					.host {
						span {
							color: $main-blue !important;
							text-decoration: underline;
						}
					}
				}

				.menu {
					.icon {
						svg {
							path {
								stroke: $main-blue !important;
							}
						}
					}
				}
			}
		}

		.check-box {
			background-color: $main-blue !important;
			opacity: 1 !important;
			border-color: $main-blue;

			svg {
				visibility: visible;

				path {
					stroke: white !important;
				}
			}
		}

		&:hover {
			box-shadow: none !important;

			.top {
				.img-container {
					img {
						filter: brightness(80%);
					}
				}
			}
		}
	}
</style>
