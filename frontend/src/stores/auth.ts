import { defineStore } from 'pinia'
import { useRuntimeConfig, navigateTo, useRoute } from '#app'

interface UserInfo {
  id: string
  email: string
  name: string
  picture: string
  google_id: string
  created_at: string
  last_login_at: string
}

interface GoogleCredentialResponse {
  credential: string
}

interface AuthResponse {
  success: boolean
  data: {
    user: UserInfo
    token: string
  }
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    user: null as UserInfo | null,
    isLoading: false
  }),
  
  actions: {
    setUser(user: UserInfo | null) {
      this.user = user
      this.isAuthenticated = !!user
    },
    
    async initializeAuth() {
      const token = localStorage.getItem('accessToken')
      if (!token) return

      try {
        this.isLoading = true
        const config = useRuntimeConfig()
        
        // 使用 refresh_token 接口验证并刷新 token
        const result = await $fetch<AuthResponse>(`${config.public.apiBaseUrl}/api/auth/refresh_token`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`
          }
        })

        if (result.success && result.data?.user) {
          this.setUser(result.data.user)
          // 更新 token
          if (result.data.token) {
            localStorage.setItem('accessToken', result.data.token)
          }
        } else {
          this.logout()
        }
      } catch (error) {
        console.error('Failed to initialize auth:', error)
        this.logout()
      } finally {
        this.isLoading = false
      }
    },
    
    async handleGoogleCredential(response: GoogleCredentialResponse) {
      try {
        console.log('handleGoogleCredential response:', response)
        this.isLoading = true
        const token = response.credential
        const config = useRuntimeConfig()
        
        // 调用后端 API 进行验证
        const result = await $fetch<AuthResponse>(`${config.public.apiBaseUrl}/api/auth/google-one-tap`, {
          method: 'POST',
          body: { 
            credential: token,
          }
        })
        
        console.log('API response:', result)
        
        // 如果后端验证成功，设置用户信息
        if (result.success && result.data?.user) {
          console.log('Setting user:', result.data.user)
          this.setUser(result.data.user)
          
          // 存储访问令牌
          if (result.data.token) {
            localStorage.setItem('accessToken', result.data.token)
            // 发送登录成功事件
            window.dispatchEvent(new CustomEvent('login-success'))
          }
        } else {
          console.error('Authentication failed:', result)
          throw new Error('Authentication failed')
        }
        
      } catch (error) {
        console.error('Google login error:', error)
        this.logout()
      } finally {
        this.isLoading = false
      }
    },
    
    initGoogleOneTap() {
      const config = useRuntimeConfig()
      console.log('initGoogleOneTap', config)
      if (typeof window !== 'undefined' && !this.isAuthenticated) {
        console.log('initGoogleOneTap', config.public.googleClientId)
        window.google?.accounts.id.initialize({
          client_id: config.public.googleClientId,
          callback: this.handleGoogleCredential,
          auto_select: true, // 启用自动登录
          cancel_on_tap_outside: false
        })
        
        window.google?.accounts.id.prompt()
      }
    },
    
    logout() {
      if (typeof window !== 'undefined') {
        window.google?.accounts.id.disableAutoSelect()
      }
      this.user = null
      this.isAuthenticated = false
      localStorage.removeItem('accessToken')
      // 触发登出事件
      window.dispatchEvent(new CustomEvent('logout'))
    }
  },
}) 