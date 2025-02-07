import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
import { viteStaticCopy } from 'vite-plugin-static-copy'
import { visualizer } from 'rollup-plugin-visualizer'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    svelte(),
    viteStaticCopy({
      targets: [
        {
          src: 'src/vditor/js', // Path to your module
          dest: 'dist', // Destination directory in 'dist'
        },
        {
          src: 'src/vditor/css', // Path to your module
          dest: 'dist', // Destination directory in 'dist'
        },
        {
          src: 'src/vditor/images', // Path to your module
          dest: 'dist', // Destination directory in 'dist'
        },
      ],
    }),
    // visualizer(),
  ]
})
