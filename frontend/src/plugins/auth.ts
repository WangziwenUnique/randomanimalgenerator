import { useAuthStore } from '../stores/auth'

export default defineNuxtPlugin(async () => {
  // 仅在客户端执行
  if (process.client) {
    const authStore = useAuthStore()
    await authStore.initializeAuth()
  }
}) 