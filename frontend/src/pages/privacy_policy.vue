<template>
  <div class="container mx-auto px-4 py-8">
    <div class="text-center mb-12 mt-16">
      <h1 class="text-4xl font-bold mb-4">Privacy Policy</h1>
      <p class="text-gray-600 dark:text-gray-400">Last updated: January 2024</p>
    </div>

    <div class="prose dark:prose-invert max-w-none">
      <ContentDoc path="/privacy_policy">
        <template #not-found>
          <div class="text-center py-8">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-emerald-500 border-t-transparent mb-4"></div>
            <h2 class="text-2xl font-semibold mb-4">Loading...</h2>
            <p class="text-gray-600 dark:text-gray-400">Please wait while the content is being loaded. If it takes too long, try refreshing the page.</p>
          </div>
        </template>
      </ContentDoc>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  title: 'Privacy Policy - Best AI Humanizer',
  description: 'Privacy Policy for Best AI Humanizer - Your AI Content Humanization Tool'
})
</script>

<style lang="postcss">
.prose {
  @apply mx-auto;
}
.prose h1 {
  @apply text-4xl font-bold mb-8;
}
.prose h2 {
  @apply text-2xl font-semibold mt-8 mb-4;
}
.prose p {
  @apply mb-4;
}
.prose a {
  @apply text-emerald-600 hover:text-emerald-500;
}
.dark .prose a {
  @apply text-emerald-500 hover:text-emerald-400;
}
</style> 