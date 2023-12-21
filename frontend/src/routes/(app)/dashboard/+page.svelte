<script>
	import { onMount } from 'svelte';
	import { pb, currentUser, axiosInstance } from '$lib/pocketbase';
	import { Progress } from '$lib/components/ui/progress';
	import { Switch } from '$lib/components/ui/switch';
	import { Label } from '$lib/components/ui/label';
	import { Toaster, toast } from 'svelte-french-toast';

	onMount(() => {
		getQuotaScreenshot();
	});
	let quota = {
		screenshots_taken: 0,
		included_screenshots: 0
	};

	let disableExtra = false;

	$: {
		if (disableExtra || !disableExtra) {
			axiosInstance.post('/update_disable_extra', {
				status: disableExtra
			}).then((res) => {
				toast.success('Success', {
					duration: 2000,
					position: 'top-right'
				});
			}).catch((err) => {
				toast.error('Error', {
					duration: 2000,
					position: 'top-right'
				});
			});
		}
    }

	const getQuotaScreenshot = async () => {
		const userId = $currentUser?.id;
		const quotaCollection = await pb
			.collection('screenshot_usage')
			.getFirstListItem(`user_id = '${userId}'`, {
				expand: 'user_id,user_id.subscription_plan',
				fields: 'screenshots_taken,expand.user_id.expand.subscription_plan.included_screenshots'
			});
		// const quotaCollection = await pb
		// 	.collection('users')
		// 	.getFirstListItem(`id = '${userId}'`, {
		// 		expand: 'subscription_plan,screenshot_usage',

		// 		fields:
		// 			'expand.screenshot_usage,expand.subscription_plan.included_screenshots'
		// 	});
		quota = {
			screenshots_taken: quotaCollection.screenshots_taken,
			included_screenshots:
				quotaCollection?.expand?.user_id?.expand?.subscription_plan?.included_screenshots
		};
	};
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Welcome to Your Dashboard!</h1>
		<p class="text-mute mt-2">
			Quickly access key metrics, recent activities, and manage your settings. All your essential
			tools and insights, streamlined in one place.
		</p>
	</div>
	<div class="flex space-x-3 w-full">
		<div class="bg-white p-5 rounded-md w-full">
			<h2 class="font-bold text-2xl">Your Monthly Quota</h2>
			<p class="text-mute mt-2">
				You have taken <span class="font-bold">{quota.screenshots_taken}</span> screenshots this
				month. Your plan includes
				<span class="font-bold">{quota.included_screenshots}</span> screenshots.
				<a href="/subscription" class="text-red-600 hover:text-red-800 hover:underline"
					>Consider upgrading</a
				> for more!
			</p>
			<div class="my-1">
				<Progress value={quota.screenshots_taken} max={quota.included_screenshots} />
			</div>
		</div>
		<!-- * SWITCH DISABLE OR ENABLE EXTRA SCREENSHOT TAKEND -->
		<div class="bg-white p-5 rounded-md w-4/12">
			<h2 class="font-bold text-2xl">Activate Extra Usage</h2>
			<div class="flex items-center space-x-2 h-full">
				<Switch id="extra" bind:checked={disableExtra} />
				<Label for="extra">
					<span class="font-bold">Disable</span> extra screenshot taken
				</Label>
			</div>
		</div>
	</div>
</div>
<Toaster />