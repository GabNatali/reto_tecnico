<script setup lang="ts">
import { computed, defineProps, defineEmits, ref, watch } from 'vue'

interface Props {
  totalItems: number
  currentPage: number
  pageSize: number
  pageSizeOptions: number[]
}

const props = defineProps<Props>()

const localPageSize = ref(props.pageSize)
const totalPages = computed(() => Math.ceil(props.totalItems / props.pageSize))
const startItem = computed(() => (props.currentPage - 1) * props.pageSize + 1)
const endItem = computed(() => Math.min(props.currentPage * props.pageSize, props.totalItems))

const emit = defineEmits(['page-change', 'page-size-change'])
const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    emit('page-change', page)
  }
}

const changePageSize = (newSize: number ) => {
  localPageSize.value = newSize
  emit('page-size-change', newSize)
}

watch(
  () => props.pageSize,
  (newSize) => {
    changePageSize(newSize);
  }
);
</script>

<template>
  <div class="flex items-center justify-between p-4">
    <div class="text-gray-900 text-sm flex items-center gap-4">
      Items por pagina

      <select
        v-model="localPageSize"
        @change="(event) => changePageSize(Number((event.target as HTMLSelectElement).value))"
        class="border border-gray-300 rounded px-2 py-1"
      >
        <option v-for="size in pageSizeOptions" :key="size" :value="size">
          {{ size }}
        </option>
      </select>
      {{ startItem }} - {{ endItem }} de {{ totalItems }}

      <div class="flex items-center space-x-2">
        <button
          @click="changePage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="px-3 py-1 rounded disabled:opacity-50"
        >
          <svg
            class="w-3 h-3 rtl:rotate-180"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 1 1 5l4 4"
            />
          </svg>
        </button>
        <button
          @click="changePage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="px-3 py-1 rounded disabled:opacity-50"
        >
          <svg
            class="w-3 h-3 rtl:rotate-180"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="m1 9 4-4-4-4"
            />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>
