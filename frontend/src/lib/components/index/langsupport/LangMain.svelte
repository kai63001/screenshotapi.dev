<script>
	import { CodeBlock } from '@skeletonlabs/skeleton';
	import hljs from 'highlight.js/lib/core';
	import xml from 'highlight.js/lib/languages/xml'; // for HTML
	import javascript from 'highlight.js/lib/languages/javascript'; // for HTML
	import go from 'highlight.js/lib/languages/go'; // for HTML
	import php from 'highlight.js/lib/languages/php'; // for HTML
	import java from 'highlight.js/lib/languages/java'; // for HTML
	hljs.registerLanguage('xml', xml); // for HTML
	hljs.registerLanguage('javascript', javascript); // for HTML
	hljs.registerLanguage('go', go); // for HTML
	hljs.registerLanguage('php', php); // for HTML
	hljs.registerLanguage('java', java); // for HTML
	import 'highlight.js/styles/github-dark.css';
	import { storeHighlightJs } from '@skeletonlabs/skeleton';
	import Icon from '@iconify/svelte';
	storeHighlightJs.set(hljs);

	let languages = [
		{
			name: 'Node.js',
			value: 'nodejs',
			lang: 'javascript',
			icon: 'mdi:nodejs',
			code: `// $ npm install screenshotone-api-sdk --save

import * as fs from 'fs';
import * as screenshotone from 'screenshotone-api-sdk';

// create API client 
const client = new screenshotone.Client("<access key>", "<secret key>");

// set up options
const options = screenshotone.TakeOptions
    .url("https://example.com")
    .delay(3)
    .blockAds(true);    

// generate URL 
const url = client.generateTakeURL(options);
console.log(url);
// expected output: https://api.screenshotone.com/take?url=...

// or download the screenshot
const imageBlob = await client.take(options);
const buffer = Buffer.from(await imageBlob.arrayBuffer());
fs.writeFileSync("example.png", buffer)
// the screenshot is stored in the example.png file`
		},
		{
			name: 'PHP',
			value: 'php',
			icon: 'material-symbols:php',
			code: `<?php`,
			lang: 'php'
		},
		{
			name: 'Go',
			value: 'go',
			icon: 'file-icons:go-old',
			code: `import "net/http"`,
			lang: 'go'
		},
		{
			name: 'Java',
			value: 'java',
			icon: 'bxl:java',
			code: `import java.net.http.HttpClient;`,
			lang: 'java'
		}
	];
	let selectedLanguage = 'nodejs';

	function selectLanguage(language) {
		selectedLanguage = language;
	}

	function getLanguageCode(language) {
		return {
			code: languages.find((item) => item.value === language)?.code,
			lang: languages.find((item) => item.value === language)?.lang
		};
	}
</script>

<section class="py-4 flex flex-col md:space-x-10 w-full">
	<div class="text-center mt-10 m-auto">
		<h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">Use the language you prefer</h2>
		<p class="mt-4 text-lg text-gray-500">
			Send simple HTTP requests or leverage native libraries in your preferred programming language.
		</p>
		<div class="flex items-center justify-center space-x-2 mt-4 relative" />
	</div>
	<div class="flex items-center justify-center mt-4 space-x-3">
		{#each languages as language}
			<button
				class="hover:bg-red-600 hover:text-white duration-300 font-bold py-2 px-4 rounded flex space-x-2 items-center"
				class:bg-red-600={selectedLanguage == language.value}
				class:text-white={selectedLanguage == language.value}
				on:click={() => selectLanguage(language.value)}
			>
				<Icon icon={language.icon} width="20px" height="20px" />
				<div>
					{language.name}
				</div>
			</button>
		{/each}
	</div>
	<div class="mt-2 overflow-hidden rounded-md">
		{#if getLanguageCode(selectedLanguage)?.code}
			<CodeBlock
				language={getLanguageCode(selectedLanguage)?.lang}
				code={getLanguageCode(selectedLanguage)?.code}
				lineNumbers
			/>
		{/if}
	</div>
</section>
