<script>
	const faqs = [
		{
			title: 'What is ScreenshotAPI?',
			content:
				'ScreenshotAPI is a screenshot service that allows you to take screenshots of any website. You can use our API to take screenshots of websites at scale.'
		},
		{
			title: 'Do you store my credit card details?',
			content:
				'We do not store your credit card details. We use Stripe, a secure payment gateway, to handle all payment transactions.'
		}
	];
</script>

<section class="py-4 flex flex-col mt-10">
	<div class="m-auto text-center">
		<h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">Frequently Asked Questions</h2>
		<p class="mt-4 text-lg text-gray-500">
			Find answers to commonly asked questions about our service.
		</p>
	</div>
	<div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
		{#each faqs as faq}
			<div class="bg-gray-100 p-4">
				<h3 class="text-lg font-semibold">{faq.title}</h3>
				<p class="mt-2 text-gray-600">{faq.content}</p>
			</div>
		{/each}
	</div>
</section>
