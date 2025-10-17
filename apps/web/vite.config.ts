import { defineConfig } from 'vite'
import { tanstackStart } from '@tanstack/react-start/plugin/vite'
import viteReact from '@vitejs/plugin-react'
import viteTsConfigPaths from 'vite-tsconfig-paths'
import tailwindcss from '@tailwindcss/vite'
import { VitePluginFonts } from 'vite-plugin-fonts'
import netlify from '@netlify/vite-plugin-tanstack-start'

export default defineConfig({
  plugins: [
    tailwindcss(),
    tanstackStart(),
    netlify(),
    viteReact(),
    viteTsConfigPaths({
      projects: ['./tsconfig.json'],
    }),
    VitePluginFonts({
      google: {
        families: [
          { name: 'Fredoka', styles: 'wght@300..700' },
          { name: 'Quicksand', styles: 'wght@300..700' },
          { name: 'Inter', styles: 'wght@100..900' }, 
        ],
      },
      // @ts-ignore
       order: 'pre',
    }),
  ],
})
