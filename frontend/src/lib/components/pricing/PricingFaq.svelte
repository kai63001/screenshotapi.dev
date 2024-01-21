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
		},
		{
			title: 'How do I sign up for ScreenshotAPI?',
			content:
				'To sign up for ScreenshotAPI, simply visit our website and click on the "Get Started for Free" button. Fill in the required information and you\'ll be ready to start using our service.'
		},
		{
			title: 'What payment methods do you accept?',
			content: 'We accept all major credit cards, including Visa, Mastercard, and American Express.'
		},
		{
			title: 'Can I cancel my subscription at any time?',
			content:
				'Yes, you can cancel your subscription at any time. Simply go to your account settings and click on the "Manage Subscription" button.'
		},
		{
			title: 'How can I contact support?',
			content:
				'You can contact our support team by sending an email to support@screenshotapi.dev or by using the live chat feature on our website.'
		},
		{
			title: 'How long does it take to generate a screenshot?',
			content:
				'The time it takes to generate a screenshot depends on the complexity of the website and the server load. In most cases, it takes just a few seconds.'
		},
		{
			title: 'Can I customize the size and format of the screenshots?',
			content:
				'Yes, you can customize the size and format of the screenshots using our API. You can specify the width, height, and format (PNG, JPEG or WEBP) in the API request.'
		},
		{
			title: 'How does pricing work?',
			content:
				'Our pricing is based on the number of screenshots you generate. We offer flexible plans to suit your needs. You can view our pricing details on our website.'
		},
		{
			title: 'Is there a free plan available?',
			content:
				'Yes, we offer a free plan that allows you to generate up to 150 screenshots per month. Sign up on our website to get started with the free plan.'
		},
		{
			title: 'How secure is my data?',
			content:
				'At ScreenshotAPI, we take data security seriously. We use industry-standard encryption and follow best practices to ensure the security of your data.'
		},
		{
			title: 'Are there any additional charges for monthly or yearly plans?',
			content:
				'Yes, if you choose the per month plan, there will be an additional charge per month. Similarly, if you choose the per year plan, there will be an additional charge per year.'
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
