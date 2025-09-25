import { defineConfig } from 'vite'
import { tanstackStart } from '@tanstack/react-start/plugin/vite'
import viteReact from '@vitejs/plugin-react'
import viteTsConfigPaths from 'vite-tsconfig-paths'
import tailwindcss from '@tailwindcss/vite'
import ViteFonts from 'vite-plugin-fonts'

const config = defineConfig({
  plugins: [
    // this is the plugin that enables path aliases
    tailwindcss(),
    tanstackStart(),
    viteReact(),
    viteTsConfigPaths({
      projects: ['./tsconfig.json'],
    }),
    ViteFonts({
      google: {
        families: [
          {
            name: 'Fredoka',
            styles: 'wght@300..700',
          },
                    {
            name: 'Quicksand',
            styles: 'wght@300..700',
          },
        ],
      },
    }),
  ],
})

export default config
