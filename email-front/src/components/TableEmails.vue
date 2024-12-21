<script setup lang="ts">
import type { IEmail } from '@/interfaces'
import { ref } from 'vue'

interface Props {
  data: IEmail[]
  currentPage: number
  pageSize: number
  isSelected: boolean
}
const selected = ref<IEmail | null>(null)
const props = defineProps<Props>()
const columns: (keyof IEmail)[] = ['subject', 'to', 'from']
const displayColumn = ['subject', 'to', 'from']

const emit = defineEmits(['select'])
const selectedRow = (email: IEmail) => {
  selected.value = email
  emit('select', email)
}
</script>
<template>
  <div class="flex flex-col w-full gap-4 flex-auto">
    <div
      class="relative overflow-x-auto overflow-auto mx-h-[calc(100vh-20rem)] md:h-[calc(100vh-17rem)] scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-white"
    >
      <table class="w-full text-sm text-left rtl:text-right">
        <thead class="text-xs text-gray-700 border-b-2 border-b-gray-300 uppercase">
          <tr>
            <th scope="col" class="pl-2 py-3 sticky top-0 border-b bg-white w-1">Nro</th>
            <th
              v-for="(column, index) in displayColumn"
              :key="index"
              scope="col"
              class="pl-2 py-3 sticky top-0 border-b bg-white"
            >
              {{ column }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(email, index) in props.data"
            :key="index"
            class="border-b border-b-gray-200 pointer"
            @click="selectedRow(email)"
            :class="{
              'bg-blue-100': isSelected && selected && selected.message_id === email.message_id,
              'bg-white': !selected || selected.message_id !== email.message_id,
            }"
          >
            <td class="pl-4 py-4 font-medium text-gray-900 whitespace-nowrap">
              {{ pageSize * currentPage + index + 1 }}
            </td>
            <td
              v-for="(column, index) in columns"
              :key="index"
              scope="row"
              class="px-2 py-4 font-medium text-gray-900 whitespace-nowrap truncate w-auto max-w-28"
            >
              {{ email[column] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
