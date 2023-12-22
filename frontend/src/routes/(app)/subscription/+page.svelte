<script>
	import Switch from '$lib/components/ui/switch/switch.svelte';
	import { pb, currentUser, axiosInstance } from '$lib/pocketbase';
	import axios from 'axios';
	import Label from '$lib/components/ui/label/label.svelte';
	import {pricingPlans as listPricingPlans} from '$lib/listPricingFeature';
	import { onMount } from 'svelte';

	let pricingPlans = listPricingPlans;

	let isYearly = false;
	let currentIndex = 0;

	onMount(async () => {
		//get subscription plan list
		const data = await pb.collection('subscription_plans').getFullList();

		//map data and pricingPlans with name
		pricingPlans = pricingPlans.map((plan) => {
			const planData = data.find((item) => item.name === plan.name);
			if (planData) {
				plan.id = planData.id;
				plan.pricing.monthly = planData.monthly;
				plan.pricing.yearly = planData.yearly;
				plan.stripe_pricing_id = planData.stripe_pricing_id;
			}
			return plan;
		});

		//get current index of pricing plan
		currentIndex = pricingPlans.findIndex((plan) => plan.id === $currentUser?.subscription_plan);
	});

	let loading = false;

	const subscription = async (planId) => {
		//check if has been subscribed
		if ($currentUser?.subscription_plan != pricingPlans[0].id && $currentUser?.subscription_status == 'active') {
			portalSubscription();
			return;
		}
		if (!planId) return;
		loading = true;
		const { data } = await axiosInstance.post(`/subscription`, {
			plan_id: planId,
			is_yearly: isYearly
		});
		const url = data.url;
		//open stripe checkout with url
		window.open(url, '_blank');
		loading = false;
	};

	const portalSubscription = async () => {
		if (!$currentUser?.stripe_customer_id) return;

		loading = true;
		const { data } = await axiosInstance.post(`/portal`);
		const url = data.portal.url;
		//open stripe checkout with url
		window.open(url, '_blank');
		loading = false;
	};
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Manage Your Subscription</h1>
		<p class="text-mute mt-2">
			View your current plan and explore various upgrade options. Manage your billing information
			and customize your subscription to fit your needs.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md">
		<div class="bg-gray-100 py-12">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div class="text-center">
					<h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">Pricing Plans</h2>
					<p class="mt-4 text-lg text-gray-500">Choose a plan that works best for your business</p>
					<div class="flex items-center justify-center space-x-2 mt-4 relative">
						<Label for="changeYearly">
							<span class="text-gray-500">Monthly</span>
						</Label>
						<Switch id="changeYearly" bind:checked={isYearly} />
						<Label for="changeYearly">
							<span class="text-gray-500">Yearly</span>
							<div
								class="absolute ml-12 text-xs top-0 bg-primary text-white rounded-full px-2 py-1"
							>
								Save 2 months
							</div>
						</Label>
					</div>
				</div>
				<div class="mt-5">
					<div class="flex justify-center">
						<div class="w-full grid grid-cols-1 gap-4 md:grid-cols-3 lg:grid-cols-4">
							{#each pricingPlans as plan, index}
								<div class="bg-white rounded-lg shadow-lg">
									<div class="px-6 py-8">
										<div class="text-center">
											<h3 class="text-lg leading-6 font-medium text-gray-900">
												{plan.name}
											</h3>
											<div class="mt-4 flex items-center justify-center">
												<span
													class="px-1 flex items-start text-6xl leading-none tracking-tight text-gray-900"
												>
													<span class="text-4xl font-medium"
														>${isYearly
															? Math.round(pricingPlans[index].pricing.yearly / 12)
															: pricingPlans[index].pricing.monthly}</span
													>
												</span>
												<span class="text-xl leading-7 font-medium text-gray-500"> /mo </span>
											</div>
										</div>
										<div class="mt-6">
											<ul class="space-y-2">
												{#each plan.features as feature}
													<li class="flex items-start">
														<div class="flex-shrink-0">
															<!-- Heroicon name: check -->
															<svg
																class="h-6 w-6 text-green-500"
																xmlns="http://www.w3.org/2000/svg"
																fill="none"
																viewBox="0 0 24 24"
																stroke="currentColor"
															>
																<path
																	stroke-linecap="round"
																	stroke-linejoin="round"
																	stroke-width="2"
																	d="M5 13l4 4L19 7"
																/>
															</svg>
														</div>
														<p class="ml-3 text-base leading-6 text-gray-500">
															{feature.name}
														</p>
													</li>
												{/each}
											</ul>
										</div>
										<div class="mt-8">
											<div class="rounded-lg shadow-md">
												{#if plan.id == $currentUser?.subscription_plan}
													<button
														on:click={portalSubscription}
														class="block w-full text-center rounded-lg border border-transparent bg-red-700 px-6 py-3 text-base leading-6 font-medium text-white hover:bg-red-500 focus:outline-none focus:red-indigo-700 focus:shadow-outline-indigo transition duration-150 ease-in-out"
													>
														{#if loading}
															<div class="flex justify-center">
																<svg
																	class="animate-spin text-center -ml-1 mr-3 h-5 w-5 text-white"
																	xmlns="http://www.w3.org/2000/svg"
																	fill="none"
																	viewBox="0 0 24 24"
																>
																	<circle
																		class="opacity-25"
																		cx="12"
																		cy="12"
																		r="10"
																		stroke="currentColor"
																		stroke-width="4"
																	/>
																	<path
																		class="opacity-75"
																		fill="currentColor"
																		d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 0012 20c4.411 0 8-3.589 8-8h-2c0 3.309-2.691 6-6 6-3.309 0-6-2.691-6-6H6c0 4.411 3.589 8 8 8a7.962 7.962 0 005.291-2H6z"
																	/>
																</svg>
															</div>
														{:else if $currentUser?.subscription_plan == plan.id && $currentUser?.stripe_customer_id}
															Manage Subscription
														{:else}
															Current Plan
														{/if}
													</button>
												{:else}
													<button
														on:click={() => {
															if (index == 0) {
																portalSubscription();
																return;
															}
															subscription(plan.id);
														}}
														class="block w-full text-center rounded-lg border border-transparent bg-primary px-6 py-3 text-base leading-6 font-medium text-white hover:bg-red-500 focus:outline-none focus:red-indigo-700 focus:shadow-outline-indigo transition duration-150 ease-in-out"
														disabled={loading}
														class:opacity-50={loading}
													>
														{#if loading}
															<div class="flex justify-center">
																<svg
																	class="animate-spin text-center -ml-1 mr-3 h-5 w-5 text-white"
																	xmlns="http://www.w3.org/2000/svg"
																	fill="none"
																	viewBox="0 0 24 24"
																>
																	<circle
																		class="opacity-25"
																		cx="12"
																		cy="12"
																		r="10"
																		stroke="currentColor"
																		stroke-width="4"
																	/>
																	<path
																		class="opacity-75"
																		fill="currentColor"
																		d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 0012 20c4.411 0 8-3.589 8-8h-2c0 3.309-2.691 6-6 6-3.309 0-6-2.691-6-6H6c0 4.411 3.589 8 8 8a7.962 7.962 0 005.291-2H6z"
																	/>
																</svg>
															</div>
														{:else if currentIndex > index}
															Downgrade
														{:else if index == 0 && $currentUser?.subscription_status != 'active'}
															Current Plan
														{:else}
															Upgrade
														{/if}
													</button>
												{/if}
											</div>
										</div>
									</div>
								</div>
							{/each}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
