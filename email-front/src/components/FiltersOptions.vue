<script setup lang="ts">
import { onMounted, ref } from 'vue';
import {  initDropdowns } from 'flowbite'
import type {TimeUnit } from '@/interfaces';
import { parseTime } from '@/actions/partseTime';

const filters = ref({
  from:"",
  to:"",
  timestampH:50,
  timestampM:'day' as TimeUnit
})


const emit = defineEmits(['send-filters'])


const getTime = () => {

  const {endTime , startTime } = parseTime(filters.value.timestampH,filters.value.timestampM)

  const dataEmit = {
    start_time: startTime.toString(),
    end_time: endTime.toString(),
  }

  return dataEmit
}

const sendFilters = () => {

  const dataEmit = {
    ...getTime(),
    to: filters.value.to,
    fromEmail: filters.value.from
  }

  emit('send-filters', dataEmit)

}

const clearFilters = () => {
  filters.value = {
    from:"",
    to:"",
    timestampH:50,
    timestampM:"day"
  }
  emit('send-filters', getTime())

}

onMounted(() => {
  initDropdowns();
  emit('send-filters', getTime())
})
</script>

<template>

  <button id="dropdownDefaultButton" data-dropdown-toggle="dropdown" class="text-blue-700 bg-blue-200 hover:bg-blue-500 hover:text-white  focus:outline-none  font-medium rounded-md text-sm px-5 py-2.5 text-center inline-flex items-center " type="button">
    Filtros
  </button>

  <!-- Dropdown menu -->
  <div id="dropdown" class="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-68">
      <ul class="py-2 text-sm text-gray-700 p-2" aria-labelledby="dropdownDefaultButton">
        <li class="mt-3">
          <p class="block px-4 py-2">Timestamp</p>
          <div class="flex px-2 gap-2 ">
            <input
              v-model.trim="filters.timestampH"
              type="number" class="border border-gray-300 rounded px-2 py-2 w-24">
            <select
                v-model="filters.timestampM"
                id="time" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-sm focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
              <option value=hour selected >Hours</option>
              <option value=minute>Minutes</option>
              <option value=day>days</option>
              <option value=week>Weeks</option>
            </select>
          </div>

        </li>
        <li class="mt-3">
          <p class="block px-4 p-2 ">From</p>
          <div class="flex px-2 gap-2 ">
            <input
            v-model.trim="filters.from"
              type="text" class="border border-gray-300 rounded px-2 w-full">
          </div>
        </li>

        <li class="mt-3">
          <p class="block px-4 p-2 ">To</p>
          <div class="flex px-2 gap-2 ">
            <input
              v-model.trim="filters.to"
              type="text" class="border border-gray-300 rounded px-2 w-full">
          </div>
        </li>

        <li>
          <div class="flex gap-2 p-2 my-3">
            <button
              @click="clearFilters" class="w-full bg-white text-blue-600 py-3 rounded-md font-bold hover:border border-blue-600"
            >Limpiar</button>
            <button
              @click="sendFilters"
              class="w-full bg-blue-600 text-white py-3 rounded-md font-bold hover:bg-blue-800">Buscar</button>
          </div>
        </li>
      </ul>

  </div>

</template>

