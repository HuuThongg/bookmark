<script lang="ts">
	import GoogleButton from '$lib/components/googleButton.svelte';
	import Mainnav from '$lib/components/mainnav.svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	// import { decodeJwt } from '$lib/utils/decodejwt';
	// globalThis.handleCredentialResponse = async (response: any) => {
	// 	const v = decodeJwt(response.credential);
	//
	// 	console.log('v :', v);
	//
	// 	//await continueWithGoogle(v);
	// };

	$: $page.data.s
		? goto(`${$page.url.origin}/appv1/my_links`)
		: console.log('no session found in local storage');
</script>

<h1>Welcome to SvelteKit</h1>
<svelte:head>
	<script src="https://accounts.google.com/gsi/client"></script>
	<title>Linkspace - Free online standalone bookmark manager</title>
</svelte:head>

<Mainnav />
<div>Hello Main nav</div>
<div class="google_btn_wrapper">
	<GoogleButton />
</div>
<div class="copy">
	<div class="above_the_fold">
		<div class="heading_and_subheading">
			<h1>Free online standalone bookmark manager.</h1>
			<p>Want to save and share links online? Liskpspace is where to.</p>
		</div>
		<div class="cta">
			<div class="g_button_wrapper">
				<GoogleButton />
			</div>
			<!-- <a
				href="/accounts/email"
				class="sign_up"
				on:click|preventDefault|stopPropagation={simulateGoogleBtnClick}
			>
				<span>Create my free account</span>
			</a> -->
			<div class="benefits">
				<div class="free_forever">
					<i class="las la-check" />
					<span>It's free, forever.</span>
				</div>
				<div class="no_instalation">
					<i class="las la-check" />
					<span>No installation required.</span>
				</div>
				<div class="no_credit_required">
					<i class="las la-check" />
					<span>No credit card required.</span>
				</div>
			</div>
		</div>
	</div>
	<div class="screenshot">
		<div class="screenshot_container">
			<img
				src="https://ewr1.vultrobjects.com/app-assets/a07148d5-54d1-43ec-a2db-332dc28a1404"
				alt="site dashboard screenshot"
			/>
		</div>
	</div>
	<div class="features">
		<div class="heading">
			<h2>Brief overview of what's inside.</h2>
			<p>Here's what you get when you create your free account.</p>
		</div>
		<div class="bullets">
			<div class="nested_folders">
				<i class="las la-sitemap" />
				<span>Nested collections</span>
				<p>Create unlimited categories and subcategories so you can better organize your links.</p>
			</div>
			<div class="link_thumbnails">
				<i class="las la-image" />
				<span>Link thumbnails</span>
				<p>We automatically add thumbnails to your links for faster skimming on your end.</p>
			</div>
			<div class="lightning_fast_search">
				<i class="lab la-searchengin" />
				<span>Fast search</span>
				<p>Fast search-as-you-type experince. As in, find what you're looking for. Fast.</p>
			</div>
			<div class="recycle_bin">
				<i class="las la-recycle" />
				<span>Recycle bin</span>
				<p>Accidentally deleted a link? Easily restore it from the recycle bin!</p>
			</div>
			<div class="bulk_edit">
				<i class="las la-edit" />
				<span>Bulk edit</span>
				<p>
					Select and edit multiple links at once... so you don't have to manage each one by one.
				</p>
			</div>
			<div class="drag_to_move">
				<i class="las la-arrows-alt" />
				<span>Drag to move</span>
				<p>Move items by dragging them from one collection to another.</p>
			</div>
		</div>
		<div class="g_button_wrapper">
			<GoogleButton />
		</div>
		<div class="button">
			<a href="/accounts/email">
				<span>Create my free account</span>
			</a>
		</div>
	</div>
	<div class="footer">
		<div class="footer_links">
			<a href="/about">
				<span>About us</span>
			</a>
			<a href="mailto:support@linkspace.com" target="_blank">
				<span>Email us</span>
			</a>
		</div>
		<span>Copyright © 2023 linkspace</span>
	</div>
</div>

<style lang="scss">
	.google_btn_wrapper {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		overflow: auto;
		display: flex;
		background-color: rgb(253, 253, 189);
		align-items: center;
		justify-content: center;
		z-index: 7000;
		display: none;
		background-color: rgb(246, 250, 255);
	}

	.copy {
		width: 100%;
		height: 100%;
		overflow: auto;
		display: flex;
		flex-direction: column;
		background-color: rgb(253, 253, 189);
		align-items: center;
		gap: 3em;
		background-color: rgb(246, 250, 255);

		.above_the_fold {
			min-height: max-content;
			display: flex;
			align-items: center;
			justify-content: center;
			flex-direction: column;
			gap: 2em;
			padding: 5em 0;

			.heading_and_subheading {
				display: flex;
				flex-direction: column;
				align-items: center;
				text-align: center;
				gap: 1.5em;

				h1 {
					font-size: 2.5rem;
					line-height: 1.6;
					color: #2b3a55;
					font-family: 'Product Sans Medium', sans-serif;
					font-weight: 500;
				}

				p {
					font-size: 1.7rem;
					color: #263159;
					font-family: 'Arial CE', sans-serif;
					font-weight: 500;
					max-width: 60%;
				}
			}

			.cta {
				display: flex;
				flex-direction: column;
				width: 100vw;
				display: flex;
				align-items: center;
				gap: 2em;

				// a.sign_up {
				// 	text-decoration: none;
				// 	width: 40rem;
				// 	height: 5rem;
				// 	display: flex;
				// 	align-items: center;
				// 	justify-content: center;
				// 	background-color: $blue;
				// 	border-radius: 0.3rem;
				// 	transition: all 200ms ease-in-out;

				// 	span {
				// 		font-size: 1.5rem;
				// 		color: white;
				// 		font-family: 'Arial CE', sans-serif;
				// 		font-weight: 500;
				// 	}

				// 	&:hover {
				// 		filter: brightness(90%);
				// 	}
				// }

				.benefits {
					display: flex;
					align-items: center;
					justify-content: space-between;
					gap: 1em;

					div {
						display: flex;
						align-items: center;
						gap: 0.5em;

						i {
							font-size: 1.6rem;
							color: rgb(12, 169, 64);
						}

						span {
							font-size: 1.5rem;
							font-family: 'Arial CE', sans-serif;
							color: $text-color-regular-2;
						}
					}
				}
			}

			// small to medium mobile screens
			@media only screen and (max-width: 375px) {
				.heading_and_subheading {
					p {
						font-size: 2rem;
						max-width: 98%;
					}
				}

				.cta {
					padding: 1em 0.5em;
					width: 100%;
					align-items: center;

					// a.sign_up {
					// 	width: 100%;
					// }

					.benefits {
						div {
							flex-direction: column;
							text-align: center;

							i {
								font-size: 2rem;
							}
						}
					}
				}
			}

			// large mobile screens
			@media only screen and (min-width: 376px) and (max-width: 568px) {
				.heading_and_subheading {
					p {
						font-size: 2rem;
					}
				}

				.cta {
					padding: 1em 0.5em;
					width: 100vw;
					align-items: center;

					// a.sign_up {
					// 	width: 100%;
					// }

					.benefits {
						div {
							flex-direction: column;
							text-align: center;

							i {
								font-size: 2rem;
							}
						}
					}
				}
			}

			// table screens
			@media only screen and (min-width: 569px) and (max-width: 768px) {
				.heading_and_subheading {
					p {
						font-size: 2rem;
					}
				}

				.cta {
					padding: 1em 0.5em;
					width: 100vw;
					align-items: center;

					// a.sign_up {
					// 	width: 70%;
					// }

					.benefits {
						min-width: max-content;

						div {
							flex-direction: row;

							i {
								font-size: 2rem;
							}
						}
					}
				}
			}
		}
	}

	.screenshot {
		width: 100vw;
		min-height: max-content;
		display: flex;
		align-items: center;
		justify-content: center;

		.screenshot_container {
			width: 90%;
			height: auto;

			img {
				width: 100%;
				height: 100%;
				object-fit: contain;
				border-radius: 0.6rem;
				box-shadow: rgba(0, 0, 0, 0.15) 0px 2px 8px;
			}
		}

		// small to medium mobile screens
		@media only screen and (max-width: 375px) {
			.screenshot_container {
				width: 98%;
			}
		}

		// large mobile screens
		@media only screen and (min-width: 376px) and (max-width: 568px) {
			.screenshot_container {
				width: 98%;
			}
		}

		// tablet screens
		@media only screen and (min-width: 569px) and (max-width: 768px) {
			.screenshot_container {
				width: 98%;
			}
		}
	}

	.features {
		min-height: max-content;
		width: 90%;
		border-radius: 0.6rem;
		display: flex;
		flex-direction: column;
		text-align: center;
		align-items: center;
		justify-content: center;
		gap: 3em;
		padding: 1em 0;

		.heading {
			flex: 1;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 0.5em;
			width: 100%;

			h2 {
				font-size: 2.3rem;
				color: #2b3a55;
				font-family: 'Product Sans Medium', sans-serif;
				font-weight: 500;
			}

			p {
				font-size: 1.7rem;
				color: #263159;
				font-family: 'Arial CE', sans-serif;
				font-weight: 500;
			}
		}

		.bullets {
			width: 80%;
			flex: 2;
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			align-items: center;
			align-content: center;
			gap: 2em;

			div {
				display: flex;
				flex-direction: column;
				align-items: center;
				gap: 1em;

				i {
					font-size: 3.5rem;
					color: $blue;
				}

				span {
					font-size: 1.5rem;
					font-family: 'Arial CE', sans-serif;
					font-weight: 500;
				}

				p {
					font-family: 'Arial CE', sans-serif;
					font-weight: 500;
					font-size: 1.4rem;
					color: $text-color-regular-2;
				}
			}
		}

		.button {
			flex: 1;
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			display: none;

			a {
				height: 5rem;
				width: 50%;
				background-color: $blue;
				display: flex;
				align-items: center;
				justify-content: center;
				border-radius: 0.3rem;
				text-decoration: none;
				transition: all 200ms ease-in-out;

				span {
					font-size: 1.5rem;
					color: white;
					font-family: 'Arial CE', sans-serif;
					font-weight: 500;
				}

				&:hover {
					filter: brightness(90%);
				}
			}
		}

		// small to medium mobile screens
		@media only screen and (max-width: 375px) {
			width: 100%;

			.bullets {
				width: 100%;
			}

			.button {
				a {
					width: 98%;
				}
			}
		}

		// medium mobile screens
		@media only screen and (min-width: 376px) and (max-width: 568px) {
			width: 100%;

			.bullets {
				width: 99%;
			}

			.button {
				a {
					width: 98%;
				}
			}
		}

		// tablet screens
		@media only screen and (min-width: 569px) and (max-width: 768px) {
			width: 100%;

			.bullets {
				width: 99%;
			}

			.button {
				a {
					width: 98%;
				}
			}
		}
	}

	.footer {
		min-height: max-content;
		width: 100vw;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 1em;
		background-color: rgb(0, 95, 153);
		// padding: 1em 0;

		.footer_links {
			display: flex;
			align-items: center;
			gap: 1em;
		}

		a {
			text-decoration: underline;
			text-decoration-color: white;

			span {
				font-size: 1.3rem;
				font-family: 'Arial CE', sans-serif;
				font-weight: 500;
				color: white;
			}

			&:hover {
				text-decoration: underline;
				text-decoration-color: white;
			}
		}

		span {
			font-size: 1.3rem;
			font-family: 'Arial CE', sans-serif;
			font-weight: 500;
			color: white;
		}
		@media only screen and (max-width: 375px) {
			//background-color: orangered;
		}
	}
</style>
