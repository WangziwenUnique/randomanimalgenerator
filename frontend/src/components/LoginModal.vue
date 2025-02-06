<template>
  <TransitionRoot appear :show="isOpen" as="template">
    <Dialog as="div" @close="closeModal" class="relative z-50">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="duration-200 ease-in"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4 text-center">
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <DialogPanel class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white dark:bg-gray-900 p-6 text-left align-middle shadow-xl transition-all relative">
              <button
                @click="closeModal"
                class="absolute top-4 right-4 text-gray-400 hover:text-gray-500 dark:hover:text-gray-300"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
              <DialogTitle as="h3" class="text-center text-2xl font-bold text-gray-900 dark:text-white mb-4">
                Sign in to your account
              </DialogTitle>

              <div class="flex justify-center pt-2">
                <div id="g_id_signin_modal"></div>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script setup lang="ts">
import { onMounted, defineProps, defineEmits, onUnmounted } from 'vue'
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { useAuthStore } from '../stores/auth'
import { useRuntimeConfig } from '#app'

const config = useRuntimeConfig()
const authStore = useAuthStore()

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const closeModal = () => {
  emit('close')
}

declare global {
  interface Window {
    handleGoogleCredentialResponseModal: (response: any) => void
  }
}

onMounted(() => {
  window.handleGoogleCredentialResponseModal = (response) => {
    console.log("Encoded JWT ID token: " + response.credential)
    authStore.handleGoogleCredential(response)
    closeModal()
  }

  // 初始化 Google Sign-In
  if (window.google?.accounts?.id) {
    window.google.accounts.id.initialize({
      client_id: config.public.googleClientId,
      callback: window.handleGoogleCredentialResponseModal,
      auto_select: false,
      cancel_on_tap_outside: true
    })
  }

  // 添加登录成功事件监听器
  window.addEventListener('login-success', closeModal)
})

// 在组件卸载时移除事件监听器
onUnmounted(() => {
  window.removeEventListener('login-success', closeModal)
})

// 监听 props.isOpen 的变化，当弹窗打开时渲染登录按钮
watch(() => props.isOpen, (newValue) => {
  if (newValue && window.google?.accounts?.id) {
    // 使用 nextTick 确保 DOM 已更新
    nextTick(() => {
      const googleAccounts = window.google?.accounts
      if (googleAccounts?.id) {
        googleAccounts.id.renderButton(
          document.getElementById("g_id_signin_modal"),
          {
            theme: "outline",
            size: "large",
            type: "standard",
            shape: "rectangular",
            text: "signin_with",
            logo_alignment: "left"
          }
        )
      }
    })
  }
})
</script> 