// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxtjs/tailwindcss',
    'nuxt-icon',
    '@nuxt/image',
    '@nuxt/content',
    'nuxt-site-config',
    '@nuxtjs/sitemap',
    '@nuxtjs/color-mode',
    '@pinia/nuxt'
  ],

  runtimeConfig: {
    public: {
      googleClientId: process.env.GOOGLE_CLIENT_ID,
      apiBaseUrl: process.env.API_BASE_URL,
      siteUrl: process.env.SITE_URL || 'https://randomanimalgenerator.art'
    }
  },

  srcDir: 'src/',
  ssr: true,

  app: {
    head: {
      titleTemplate: '%s',
      title: 'Random Animal Generator - Discover Amazing Animals',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Generate and discover random animals from around the world. Learn about unique and fascinating creatures with our free random animal generator tool!' },
        { name: 'format-detection', content: 'telephone=no' },
        // Open Graph
        { property: 'og:title', content: 'Random Animal Generator - Discover Amazing Animals' },
        { property: 'og:description', content: 'Generate and discover random animals from around the world. Learn about unique and fascinating creatures with our free random animal generator tool!' },
        { property: 'og:type', content: 'website' },
        { property: 'og:url', content: 'https://randomanimalgenerator.art' },
        { property: 'og:image', content: 'https://randomanimalgenerator.art/favicon.png' },
        // Twitter Card
        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:title', content: 'Random Animal Generator - Discover Amazing Animals' },
        { name: 'twitter:description', content: 'Generate and discover random animals from around the world. Learn about unique and fascinating creatures with our free random animal generator tool!' },
        { name: 'twitter:image', content: 'https://randomanimalgenerator.art/favicon.png' }
      ],
      link: [
        { rel: 'icon', type: 'image/png', href: '/favicon.png' },
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Meddon&display=swap' }
      ],
      script: [
        {
          src: 'https://accounts.google.com/gsi/client',
          async: true,
          defer: true
        },
        {
          type: 'application/ld+json',
          children: JSON.stringify({
            '@context': 'https://schema.org',
            '@type': 'Organization',
            name: 'Random Animal Generator',
            url: 'https://randomanimalgenerator.art',
            logo: 'https://randomanimalgenerator.art/favicon.png',
            sameAs: [],
            description: 'Generate and discover random animals from around the world. Learn about unique and fascinating creatures.'
          })
        },
        {
          type: 'application/ld+json',
          children: JSON.stringify({
            '@context': 'https://schema.org',
            '@type': 'WebPage',
            name: 'Random Animal Generator - Discover Amazing Animals',
            description: 'Generate and discover random animals from around the world. Learn about unique and fascinating creatures with our free random animal generator tool!',
            url: 'https://randomanimalgenerator.art',
            mainEntity: {
              '@type': 'WebApplication',
              name: 'Random Animal Generator',
              applicationCategory: 'Educational Tool',
              operatingSystem: 'All',
              aggregateRating: {
                '@type': 'AggregateRating',
                ratingValue: '4.8',
                ratingCount: '1250',
                bestRating: '5',
                worstRating: '1'
              },
              offers: {
                '@type': 'Offer',
                price: '0',
                priceCurrency: 'USD'
              },
              featureList: [
                'Random Animal Generation',
                'Animal Information',
                'Educational Content',
                'Free to Use'
              ]
            }
          })
        }
      ]
    }
  },

  nitro: {
    prerender: {
      failOnError: false,
      crawlLinks: true,
      routes: [
        '/',
        '/tos',
        '/privacy_policy',
      ]
    },
    publicAssets: [
      {
        dir: 'public',
        baseURL: '/public'
      }
    ],
    routeRules: {
      '/**': {
        headers: {
          'Cross-Origin-Opener-Policy': 'same-origin-allow-popups'
        }
      },
      '/api/_content/**': {
        cache: {
          maxAge: 60 * 60
        }
      }
    },
    storage: {
      'content': {
        driver: 'fs',
        base: './.content/cache'
      }
    }
  },

  experimental: {
    viewTransition: true,
    renderJsonPayloads: true
  },

  image: {
    dir: 'public',
    domains: ['localhost'],
    format: ['webp', 'jpg', 'png'],
    provider: 'ipx',
  },

  compatibilityDate: '2025-01-08',

  site: {
    url: 'https://randomanimalgenerator.art'
  },

  sitemap: {
    exclude: [
      '/404',
      '/payment-success'
    ]
  },

  content: {
    markdown: {
      toc: {
        depth: 3,
        searchDepth: 3
      },
      anchorLinks: true,
      remarkPlugins: [],
      rehypePlugins: []
    },
    documentDriven: true,
    experimental: {
      clientDB: false,
      stripQueryParameters: true
    },
    api: {
      baseURL: '/api/_content'
    }
  },

  tailwindcss: {
    cssPath: '~/assets/css/tailwind.css',
    configPath: 'tailwind.config.js',
    exposeConfig: true,
    config: {
      content: [
        './src/components/**/*.{vue,js,ts}',
        './src/layouts/**/*.vue',
        './src/pages/**/*.vue',
        './src/plugins/**/*.{js,ts}',
        './nuxt.config.{js,ts}'
      ],
    },
    viewer: false
  },

  colorMode: {
    classSuffix: '',
    preference: 'system',
    fallback: 'dark',
    globalName: '__NUXT_COLOR_MODE__',
    componentName: 'ColorScheme',
    classPrefix: '',
    storageKey: 'nuxt-color-mode'
  }
})