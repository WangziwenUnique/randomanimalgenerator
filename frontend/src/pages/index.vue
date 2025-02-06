<template>
  <div class="flex flex-col items-center justify-center min-h-screen p-4">
    <h1 class="text-4xl font-bold mb-4">Random Animal Generator</h1>
    <p class="text-lg text-gray-600 mb-8">Generate random cute animal images</p>
    
    <div class="relative mb-4">
      <button 
        @click="showTypeSelector = !showTypeSelector"
        class="bg-blue-100 hover:bg-blue-200 text-blue-800 font-semibold py-2 px-4 rounded-lg flex items-center gap-2"
      >
        <span>{{ selectedType === 'all' ? 'All Animals' : selectedType.charAt(0).toUpperCase() + selectedType.slice(1) }}</span>
        <span class="transform transition-transform" :class="{ 'rotate-180': showTypeSelector }">▼</span>
      </button>
      
      <div 
        v-if="showTypeSelector"
        class="absolute top-full mt-1 w-48 bg-white rounded-lg shadow-lg py-2 z-10"
      >
        <button 
          @click="selectedType = 'all'; showTypeSelector = false; generateRandomAnimals()"
          class="w-full text-left px-4 py-2 hover:bg-blue-50"
        >
          All Animals
        </button>
        <button 
          v-for="(_, type) in animalTypes" 
          :key="type"
          @click="selectedType = type; showTypeSelector = false; generateRandomAnimals()"
          class="w-full text-left px-4 py-2 hover:bg-blue-50 capitalize"
        >
          {{ type }}
        </button>
      </div>
    </div>

    <button 
      @click="generateRandomAnimals" 
      class="mb-8 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-6 rounded-full transition-all duration-300 transform hover:scale-105"
    >
      Generate New Animals
    </button>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-6xl mb-12">
      <div 
        v-for="(animal, index) in selectedAnimals" 
        :key="index"
        class="bg-white rounded-lg shadow-lg overflow-hidden transform transition-all duration-300 hover:scale-105"
      >
        <img 
          :src="'/images/' + animal"
          :alt="animal.replace('.webp', '')"
          class="w-full h-64 object-cover"
          loading="lazy"
        />
        <div class="p-4">
          <h3 class="text-lg font-semibold capitalize">
            {{ animal.replace('.webp', '').replace(/-/g, ' ') }}
          </h3>
        </div>
      </div>
    </div>

    <div class="max-w-4xl mx-auto px-4 mb-16">
      <h2 class="text-3xl font-bold mb-6">About Our Random Animal Generator</h2>
      <p class="text-gray-700 mb-8">
        Discover the fascinating world of animals with our Random Animal Generator tool. Each click brings you a unique selection of 6 different animal images, perfect for education, inspiration, or just pure entertainment.
      </p>

      <h3 class="text-2xl font-semibold mb-4">How Our Random Animal Generator Works</h3>
      <p class="text-gray-700 mb-8">
        Our Random Animal Generator uses a sophisticated algorithm to select 6 unique animals from our extensive collection. Each generated set is completely random, ensuring a fresh and exciting experience every time you click the generate button.
      </p>

      <h3 class="text-2xl font-semibold mb-4">Features of the Random Animal Generator</h3>
      <ul class="list-disc list-inside text-gray-700 mb-8 space-y-2">
        <li>Instant generation of 6 unique animal images</li>
        <li>High-quality animal photographs</li>
        <li>Educational animal names and descriptions</li>
        <li>Mobile-friendly responsive design</li>
        <li>Smooth animations and transitions</li>
      </ul>

      <h3 class="text-2xl font-semibold mb-4">Uses for the Random Animal Generator</h3>
      <p class="text-gray-700 mb-8">
        Our Random Animal Generator is perfect for teachers, students, artists, and animal enthusiasts. Use it for educational purposes, creative inspiration, or simply to learn more about different animal species from around the world.
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// All animal images
const animalImages = [
  'zebu.webp', 'yak.webp', 'zebra.webp', 'wolverine.webp', 'wombat.webp',
  'woodchuck.webp', 'whale.webp', 'wildcat.webp', 'wolf.webp', 'warthog.webp',
  'waterbuck.webp', 'weasel.webp', 'vicuna.webp', 'walrus.webp', 'tiger.webp',
  'toad.webp', 'turtle.webp', 'thorny devil.webp', 'starfish.webp', 'steer.webp',
  'tapir.webp', 'squirrel.webp', 'stallion.webp', 'snake.webp', 'snowy owl.webp',
  'springbok.webp', 'silver fox.webp', 'skunk.webp', 'sloth.webp', 'sheep.webp',
  'shrew.webp', 'salamander.webp', 'seal.webp', 'rhinoceros.webp', 'rooster.webp',
  'ram.webp', 'rat.webp', 'reindeer.webp', 'rabbit.webp', 'raccoon.webp',
  'puma.webp', 'puppy.webp', 'quagga.webp', 'prairie dog.webp', 'pronghorn.webp',
  'pony.webp', 'porcupine.webp', 'porpoise.webp', 'peccary.webp', 'pig.webp'
]

const animalTypes = {
  'mammals': ['zebu', 'yak', 'zebra', 'wolverine', 'wombat', 'woodchuck', 'whale', 'wildcat', 'wolf', 'warthog', 'waterbuck', 'weasel', 'vicuna', 'walrus', 'tiger', 'squirrel', 'stallion', 'silver fox', 'skunk', 'sloth', 'sheep', 'shrew', 'seal', 'rhinoceros', 'ram', 'rat', 'reindeer', 'rabbit', 'raccoon', 'puma', 'puppy', 'pony', 'porcupine', 'porpoise', 'peccary', 'pig'],
  'birds': ['snowy owl', 'rooster'],
  'reptiles': ['snake', 'turtle', 'thorny devil'],
  'amphibians': ['toad', 'salamander'],
  'fish': ['starfish']
}

const selectedAnimals = ref<string[]>([])
const selectedType = ref<string>('all')
const showTypeSelector = ref(false)

const getAnimalsByType = (type: string) => {
  if (type === 'all') {
    return animalImages
  }
  return animalImages.filter(animal => 
    animalTypes[type]?.some(typeAnimal => 
      animal.toLowerCase().includes(typeAnimal.toLowerCase())
    )
  )
}

// Function to generate 6 random unique animals
const generateRandomAnimals = () => {
  const filteredAnimals = getAnimalsByType(selectedType.value)
  const shuffled = [...filteredAnimals].sort(() => Math.random() - 0.5)
  selectedAnimals.value = shuffled.slice(0, Math.min(6, shuffled.length))
}

// Generate initial set of animals
onMounted(() => {
  generateRandomAnimals()
})
</script>

<style scoped>
.grid {
  opacity: 1;
  transition: opacity 0.3s ease-in-out;
}
</style> 