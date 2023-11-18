<script lang="ts">
	import { currentUser, pb } from '$lib/pocketbase';

	import * as Table from '$lib/components/ui/table';
	import { onMount } from 'svelte';
	$: invoices = [];

	onMount(async () => {
		//get invoices
		pb.collection('payments')
			.getFullList()
			.then((res) => {
				invoices = res.map((item) => {
					return {
						invoice: item.id,
						paymentStatus: item.status,
						totalAmount: item.total_amount,
						productDescription: item.product_description,
						invoice_link: item.invoice_pdf
					};
				});
			});
	});
</script>

<div class="gap-4 grid">
	<div class="bg-white p-5 rounded-md">
		<h1 class="font-bold text-2xl">Payment History</h1>
		<p class="text-mute mt-2">
			Review your past transactions, including dates, amounts, and statuses, all in one place.
		</p>
	</div>
	<div class="bg-white p-5 rounded-md">
		<Table.Root>
			<Table.Caption>A list of your recent invoices.</Table.Caption>
			<Table.Header>
				<Table.Row>
					<Table.Head class="w-[100px]">Invoice</Table.Head>
					<Table.Head>Status</Table.Head>
					<Table.Head>Method</Table.Head>
					<Table.Head class="text-right">Amount</Table.Head>
					<Table.Head class="text-right">Amount</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each invoices as invoice, i (i)}
					<Table.Row>
						<Table.Cell class="font-medium">{invoice.invoice}</Table.Cell>
						<Table.Cell>{invoice.paymentStatus}</Table.Cell>
						<Table.Cell>{invoice.productDescription}</Table.Cell>
						<Table.Cell class="text-right">${(invoice.totalAmount / 100).toFixed(2)}</Table.Cell>
						<Table.Cell class="text-right">
              <a href={invoice.invoice_link} target="_blank" class="text-primary">
                View Invoice
              </a>
            </Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
</div>
