<script setup lang="ts">
import { ref } from 'vue'

const filterSearch = ref('')
let debounceTimeout: ReturnType<typeof setTimeout> | null = null

const emit = defineEmits(['search'])

function emitFilter() {
  if (debounceTimeout) {
    clearTimeout(debounceTimeout)
  }

  debounceTimeout = setTimeout(() => {
    emit('search', filterSearch.value)
  }, 500)
}
</script>

<template>
  <div class="flex">
    <div class="w-9/10 relative bg-white lg:w-[450px]">
      <div class="relative">
        <input
          class="bg-white w-full pr-11 h-10 pl-2 py-2 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md"
          placeholder="Search for Subject..."
          tyque="text"
          v-model="filterSearch"
          @input="emitFilter"
        />
        <button
          class="absolute h-8 w-8 right-1 top-1 my-auto px-2 flex items-center bg-white rounded"
          type="button"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="3"
            stroke="currentColor"
            class="w-8 h-8 text-slate-600"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"
            />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>
