import { defineNuxtPlugin, useRuntimeConfig, useHead } from '#app'

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()

  if (process.env.NODE_ENV !== 'production') {
    return
  }

  useHead({
    script: [
      {
        src: 'https://www.googletagmanager.com/gtag/js?id=G-0XTDKNEBD8',
        async: true,
      },
      {
        innerHTML: `window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());
          gtag('config', 'G-0XTDKNEBD8');`
      }
    ]
  })
}) 