import { defineNuxtPlugin, useHead, useRuntimeConfig, useRoute } from '#app'
import { computed } from 'vue'

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const route = useRoute()
  
  useHead({
    link: [
      {
        rel: 'canonical',
        href: computed(() => {
          const baseUrl = (config.public.siteUrl as string) || 'https://randomanimalgenerator.art'
          // 移除结尾的斜杠
          const cleanBaseUrl = baseUrl.replace(/\/$/, '')
          // 确保路径以斜杠开头且移除结尾斜杠
          const cleanPath = route.path === '/' ? '' : route.path.replace(/\/$/, '')
          // 组合完整的 URL
          return `${cleanBaseUrl}${cleanPath}`
        })
      }
    ]
  })
}) 