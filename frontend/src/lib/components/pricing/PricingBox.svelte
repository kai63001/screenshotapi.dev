<script>
	import { pricingPlans as listPricingPlans } from '$lib/listPricingFeature';
	import Icon from '@iconify/svelte';
	import { Label } from '../ui/label';
	import { Switch } from '../ui/switch';

	let pricingPlans = listPricingPlans;

	let isYearly = true;
</script>

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
			<div class="absolute ml-12 text-xs top-0 bg-primary text-white rounded-full px-2 py-1">
				Save 2 months
			</div>
		</Label>
	</div>
</div>
<div class="mt-5">
	<div class="flex justify-center">
		<div class="w-full grid grid-cols-1 gap-4 md:grid-cols-3 lg:grid-cols-4">
			{#each pricingPlans as plan, index}
				<div class="bg-white rounded-lg shadow-[0_0px_10px_0.1px_rgba(0,0,0,0.1)]"  class:shadow-red-600={plan.name == 'Business'} >
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
								<a
									href="/subscription"
									type="submit"
									class="w-full bg-primary text-white rounded py-3.5 text-sm group flex items-center justify-center"
								>
									Start Free
									<Icon
										icon="akar-icons:arrow-right"
										class="w-5 h-5 ml-2 transform duration-300 group-hover:translate-x-1"
									/>
								</a>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
