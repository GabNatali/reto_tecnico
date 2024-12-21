<script setup lang="ts">
import type { IEmail } from '@/interfaces'

interface Props {
  email: IEmail
}

defineProps<Props>()
</script>

<template>
  <div class="px-2 flex flex-col gap-1 justify-center">
    <div class="flex gap-1 justify-start flex-wrap">
      <p class="text-md font-medium">{{ email.subject }}</p>
      <p class="text-md">&lt;{{ email.from }}&gt;</p>
    </div>

    <p>{{ email.date }}</p>

    <p class="whitespace-pre-wrap pt-2">{{ email.body }}</p>
  </div>
</template>
