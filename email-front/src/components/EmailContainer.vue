<script setup lang="ts">
import EmailApi from '../actions/EmailApi'
import SearchApp from '../components/SearchApp.vue'
import Paginator from '../components/PaginatorTable.vue'
import TableEmails from '../components/TableEmails.vue'
import type { IEmail, IEmailResponse, IParams } from '../interfaces'
import { ref, watch } from 'vue'
import ContentEmail from './ContentEmail.vue'
import { useToast } from 'vue-toastification'
import FiltersOptions from './FiltersOptions.vue'

const isLoading = ref(false)
const toast = useToast()

// paginator
const currentPage = ref(1)
const isSelected = ref(false)
const rowSelected = ref<IEmail>({
  _timestamp: 0,
  body: '',
  date: '',
  from: '',
  message_id: '',
  subject: '',
  to: '',
})

const filters = ref<IParams>({
  fromEmail: '',
  to: '',
  subject: '',
  from: '1',
  size: '500',
  start_time:'',
  end_time: '',
  stream_log: 'email_l3'
})


const pageChange = async (page: number) => {
  console.log('pageChange', page)
  isSelected.value = false
  filters.value = {
    ...filters.value,
    from: ((page - 1) * Number(filters.value.size)).toString()
  }

  currentPage.value = page

}

const pageSizeChange = async (pageSize: number) => {

  filters.value = {
    ...filters.value,
    size: pageSize.toString(),
    from: '1'
  }
  currentPage.value = 1
  isSelected.value = false
}

const getFilters = (otherFilters: IParams) => {
  filters.value = {
    ...filters.value,
    ...otherFilters,
    from: '1',
    size: '500'
  }

  isSelected.value = false
}

const getFilterSearch = (search: string) => {
  filters.value = {
    ...filters.value,
    subject: search,
    from: '1',
    size: '500'
  }
  isSelected.value = false
}

const emailSelected = (email: IEmail) => {
  isSelected.value = true
  rowSelected.value = email
}

const response = ref<IEmailResponse>({
  count: 0,
  results: [],
})
const getEmails = async (params: IParams) => {
  isLoading.value = true
  try {
    const { data } = await EmailApi.getAllEmails(params)
    response.value = data
    isLoading.value = false
  } catch{
    toast.error("Error server try again")
  }
}

watch(
  filters,
  (newFilters) => {
    getEmails(newFilters)
  },
  { deep: true }
)


</script>

<template>
  <div class="flex flex-col w-full gap-4 flex-auto">
    <div class="w-full flex items-center justify-start flex-wrap gap-4">
      <SearchApp @search="getFilterSearch"></SearchApp>
      <FiltersOptions @send-filters="getFilters"></FiltersOptions>
    </div>
    <div class="flex flex-col w-full gap-4 flex-auto">
      <div  class="grid grid-cols-1 lg:grid-cols-2 gap-4 h-full">
        <div class="flex flex-col bg-white p-4 rounded-md">
          <TableEmails
            @select="emailSelected"
            :data="response.results"
            :current-page="currentPage - 1"
            :page-size="Number(filters.size)"
            :is-selected="isSelected"
          ></TableEmails>
          <Paginator
            :total-items="response.count"
            :current-page="currentPage"
            :page-size="Number(filters.size)"
            :page-size-options="[100, 250, 500]"
            @page-change="pageChange"
            @page-size-change="pageSizeChange"
          />
        </div>
        <div
          class="flex flex-col bg-white p-4 rounded-md overflow-auto h-[calc(100vh-11rem)] scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-white"
        >
          <ContentEmail v-if="isSelected" :email="rowSelected"></ContentEmail>
          <div v-else class="flex items-center justify-center h-full">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              class="w-28 h-28 text-gray-300"
            >
              <path
                d="M1.5 4.5h21a1.5 1.5 0 011.5 1.5v12a1.5 1.5 0 01-1.5 1.5h-21A1.5 1.5 0 010 18V6a1.5 1.5 0 011.5-1.5zM21 6l-9 6.75L3 6v12h18V6zm-9 7.5L21 6H3l9 7.5z"
              />
            </svg>
          </div>
        </div>
      </div>
      <div v-if="isLoading" class="absolute inset-0 flex items-center justify-center bg-white/50">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          style="margin: auto; background: none; display: block;"
          width="100px"
          height="100px"
          viewBox="0 0 100 100"
          preserveAspectRatio="xMidYMid"
        >
          <circle
            cx="50"
            cy="50"
            fill="none"
            stroke="#4a90e2"
            stroke-width="10"
            r="35"
            stroke-dasharray="164.93361431346415 56.97787143782138"
          >
            <animateTransform
              attributeName="transform"
              type="rotate"
              repeatCount="indefinite"
              dur="1s"
              values="0 50 50;360 50 50"
              keyTimes="0;1"
            ></animateTransform>
          </circle>
        </svg>

      </div>
    </div>
  </div>
</template>
