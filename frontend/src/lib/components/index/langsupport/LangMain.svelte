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
	storeHighlightJs.set(hljs);

	let languages = [
		{
			name: 'Node.js',
			value: 'nodejs',
			lang: 'javascript',
			icon: 'carbon:nodejs',
			code: `const axios = require('axios');`
		},
		{
			name: 'PHP',
			value: 'php',
			icon: 'vscode-icons:file-type-php',
			code: `<?php`,
			lang: 'php'
		},
		{
			name: 'Go',
			value: 'go',
			icon: 'vscode-icons:file-type-go',
			code: `import "net/http"`,
			lang: 'go'
		},
		{
			name: 'Java',
			value: 'java',
			icon: 'vscode-icons:file-type-java',
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
        }
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
				class="hover:bg-red-600 hover:text-white duration-300 font-bold py-2 px-4 rounded"
				class:bg-red-600={selectedLanguage == language.value}
				class:text-white={selectedLanguage == language.value}
				on:click={() => selectLanguage(language.value)}
			>
				{language.name}
			</button>
		{/each}
	</div>
	<div class="mt-2 overflow-hidden rounded-md">
        {#if getLanguageCode(selectedLanguage)?.code}
            <CodeBlock language={getLanguageCode(selectedLanguage)?.lang} code={getLanguageCode(selectedLanguage)?.code} lineNumbers />
        {/if}
	</div>
</section>
