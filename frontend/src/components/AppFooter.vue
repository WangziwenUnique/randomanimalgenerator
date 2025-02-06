<template>
  <footer class="bg-white dark:bg-gray-900 border-t border-gray-200 dark:border-gray-800 py-6 md:py-8">
    <div class="container mx-auto px-4">
      <div class="flex flex-col items-center mb-6 md:mb-8">
        <NuxtLink to="/" class="flex items-center space-x-2 mb-4 md:mb-6">
          <span class="font-meddon text-gray-900 dark:text-white text-xl md:text-2xl">Best AI Humanize</span>
        </NuxtLink>
        
        <div class="flex flex-wrap justify-center gap-4 md:gap-6 text-sm md:text-base text-gray-600 dark:text-gray-300">
          <NuxtLink to="/" class="hover:text-gray-900 dark:hover:text-white">Generator</NuxtLink>
          <NuxtLink to="/tos" class="hover:text-gray-900 dark:hover:text-white">Terms of Service</NuxtLink>
          <NuxtLink to="/privacy_policy" class="hover:text-gray-900 dark:hover:text-white">Privacy Policy</NuxtLink>
        </div>
      </div>
      
      <div class="text-center text-gray-500 dark:text-gray-400 text-xs md:text-sm">
        <p class="mb-2">© 2025 randomanimalgenerator.art. All rights reserved.</p>
        <p class="text-xs">
          This site is protected by reCAPTCHA and the Google 
          <a href="https://policies.google.com/privacy" class="text-emerald-600 dark:text-emerald-500 hover:text-emerald-500 dark:hover:text-emerald-400">Privacy Policy</a>
          and
          <a href="https://policies.google.com/terms" class="text-emerald-600 dark:text-emerald-500 hover:text-emerald-500 dark:hover:text-emerald-400">Terms of Service</a>
          apply.
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
</script> 