import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
import { viteStaticCopy } from 'vite-plugin-static-copy'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    svelte(),
    viteStaticCopy({
      targets: [
        {
          src: 'src/vditor/js', // Path to your module
          dest: '', // Destination directory in 'dist'
        },
        {
          src: 'src/vditor/css', // Path to your module
          dest: '', // Destination directory in 'dist'
        },
        {
          src: 'src/vditor/images', // Path to your module
          dest: '', // Destination directory in 'dist'
        },
      ],
    }),
  ]
})
