<template>
  <header class="fixed top-0 left-0 right-0 bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700 z-50">
    <nav class="container mx-auto px-4 h-20 flex items-center justify-between">
      <NuxtLink to="/" class="flex items-center">
        <img src="/images/logo-transparent.svg" alt="Humanize AI" class="h-12 md:h-16 w-auto" />
      </NuxtLink>
      
      <!-- 移动端菜单按钮 -->
      <button 
        class="md:hidden p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
        @click="isMobileMenuOpen = !isMobileMenuOpen"
      >
        <svg 
          xmlns="http://www.w3.org/2000/svg" 
          class="h-6 w-6" 
          fill="none" 
          viewBox="0 0 24 24" 
          stroke="currentColor"
        >
          <path 
            v-if="!isMobileMenuOpen" 
            stroke-linecap="round" 
            stroke-linejoin="round" 
            stroke-width="2" 
            d="M4 6h16M4 12h16M4 18h16"
          />
          <path 
            v-else 
            stroke-linecap="round" 
            stroke-linejoin="round" 
            stroke-width="2" 
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
      
      <!-- 桌面端菜单 -->
      <div class="hidden md:flex items-center space-x-6">
        <template v-for="(link, index) in navigationLinks" :key="index">
          <NuxtLink 
            :to="link.to" 
            class="relative py-1"
            :class="[
              route.path === link.to
                ? 'text-cyan-500'
                : 'text-gray-300 dark:text-gray-300 text-gray-600 hover:text-gray-900 dark:hover:text-white'
            ]"
          >
            {{ link.text }}
            <div 
              v-if="route.path === link.to"
              class="absolute bottom-0 left-0 right-0 h-0.5 bg-cyan-500"
            ></div>
          </NuxtLink>
        </template>
        
        <!-- 用户头像和下拉菜单 -->
        <Menu v-if="authStore.isAuthenticated" as="div" class="relative">
          <MenuButton class="flex items-center space-x-2">
            <div class="w-8 h-8">
              <img :src="authStore.user?.picture" :alt="authStore.user?.name" class="w-full h-full rounded-full object-cover" />
            </div>
          </MenuButton>

          <transition
            enter-active-class="transition duration-100 ease-out"
            enter-from-class="transform scale-95 opacity-0"
            enter-to-class="transform scale-100 opacity-100"
            leave-active-class="transition duration-75 ease-in"
            leave-from-class="transform scale-100 opacity-100"
            leave-to-class="transform scale-95 opacity-0"
          >
            <MenuItems class="absolute right-0 mt-2 w-48 origin-top-right divide-y divide-gray-100 rounded-md bg-white dark:bg-gray-800 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
              <div class="px-1 py-1">
                <MenuItem v-slot="{ active }">
                  <button
                    :class="[
                      active ? 'bg-gray-100 dark:bg-gray-700' : '',
                      'group flex w-full items-center rounded-md px-2 py-2 text-sm text-gray-900 dark:text-white'
                    ]"
                    @click="handleLogout"
                  >
                    Sign Out
                  </button>
                </MenuItem>
              </div>
            </MenuItems>
          </transition>
        </Menu>

        <button v-else class="flex items-center" @click="openLoginModal">
          <div class="w-8 h-8 rounded-full flex items-center justify-center bg-gray-200 dark:bg-gray-700">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 text-gray-500 dark:text-gray-400">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
            </svg>
          </div>
        </button>

        <button
          class="p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
          @click="toggleColorMode"
        >
          <ClientOnly>
            <div class="w-5 h-5">
              <svg v-if="colorMode.value === 'dark'" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="text-gray-300">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="text-gray-600">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z" />
              </svg>
            </div>
          </ClientOnly>
        </button>
      </div>
    </nav>

    <!-- 移动端菜单 -->
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform scale-y-0 opacity-0"
      enter-to-class="transform scale-y-100 opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="transform scale-y-100 opacity-100"
      leave-to-class="transform scale-y-0 opacity-0"
    >
      <div 
        v-if="isMobileMenuOpen" 
        class="md:hidden absolute top-[80px] left-0 right-0 bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700 origin-top"
      >
        <div class="px-4 py-2 space-y-2">
          <template v-for="(link, index) in navigationLinks" :key="index">
            <NuxtLink 
              :to="link.to" 
              class="block px-4 py-2 rounded-lg"
              :class="[
                route.path === link.to
                  ? 'bg-cyan-50 dark:bg-cyan-900/20 text-cyan-500'
                  : 'text-gray-600 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-800'
              ]"
              @click="isMobileMenuOpen = false"
            >
              {{ link.text }}
            </NuxtLink>
          </template>
        </div>
      </div>
    </Transition>

    <!-- 登录弹窗 -->
    <LoginModal :is-open="isLoginModalOpen" @close="closeLoginModal" />
  </header>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useColorMode, useRoute, navigateTo } from '#imports'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue'
import { useAuthStore } from '../stores/auth'
import LoginModal from './LoginModal.vue'

const colorMode = useColorMode()
const route = useRoute()
const authStore = useAuthStore()
const isLoginModalOpen = ref(false)
const isMobileMenuOpen = ref(false)

const navigationLinks = [
  { to: '/', text: 'Generator' }
]

const openLoginModal = () => {
  isLoginModalOpen.value = true
}

const closeLoginModal = () => {
  isLoginModalOpen.value = false
}

onMounted(() => {
  authStore.initGoogleOneTap()
})

// 监听路由变化
watch(() => route.path, () => {
  authStore.initGoogleOneTap()
})

const toggleColorMode = () => {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}

const handleLogout = () => {
  authStore.logout()
  localStorage.removeItem('accessToken')
}
</script> 