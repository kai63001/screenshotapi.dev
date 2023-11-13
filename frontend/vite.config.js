import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { imagetools } from 'vite-imagetools'

export default defineConfig({
	plugins: [sveltekit(), imagetools()],
	server: {
		fs: {
		  // Allow serving files from one level up to the project root
		  // posts, copy
		  allow: ['..'],
		},
	  },
});
