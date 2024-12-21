<script setup lang="ts">
import EmailApi from '../actions/EmailApi'
import SearchApp from '../components/SearchApp.vue'
import Paginator from '../components/PaginatorTable.vue'
import TableEmails from '../components/TableEmails.vue'
import type { IEmail, IEmailResponse, IParams } from '../interfaces'
import { computed, onMounted, ref } from 'vue'
import ContentEmail from './ContentEmail.vue'
import { useToast } from 'vue-toastification'


const toast = useToast()
// paginator
const currentPage = ref(1)
const currentpageSize = ref(500)
const rowSelected = ref<IEmail>({
  _timestamp: 0,
  body: '',
  date: '',
  from: '',
  message_id: '',
  subject: '',
  to: '',
})

const pageChange = async (page: number) => {
  currentPage.value = page
  params.value.from = ((page - 1) * currentpageSize.value).toString()
  getEmails(params.value)
}

const pageSizeChange = async (pageSize: number) => {
  currentpageSize.value = pageSize
  currentPage.value = 1

  getEmails(params.value)
}

const params = computed<IParams>(() => {
  return {
    fromEmail: '',
    subject: '',
    to: '',
    from: (currentPage.value - 1).toString(),
    size: currentpageSize.value.toString(),
    start_time: '1734561039155000',
    end_time: '1734647439155000',
    stream_log: 'email_l3',
  }
})

const response = ref<IEmailResponse>({
  count: 0,
  results: [],
})

const getFilterSearch = (search: string) => {
  params.value.subject = search
  currentpageSize.value = 500
  currentPage.value = 1
  getEmails(params.value)
}

const emailSelected = (email: IEmail) => {
  rowSelected.value = email
}
const getEmails = async (params: IParams) => {
  try {
    const { data } = await EmailApi.getAllEmails(params)
    response.value = data
  } catch{
    toast.error("Error server try again")
  }
}

onMounted(() => {
  getEmails(params.value)
})
</script>

<template>
  <div class="flex flex-col w-full gap-4 flex-auto">
    <div class="w-full flex items-center">
      <SearchApp @search="getFilterSearch"></SearchApp>
    </div>
    <div class="flex flex-col w-full gap-4 flex-auto">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 h-full">
        <div class="flex flex-col bg-white p-4 rounded-md">
          <TableEmails
            @select="emailSelected"
            :data="response.results"
            :current-page="currentPage - 1"
            :page-size="currentpageSize"
          ></TableEmails>
          <Paginator
            :total-items="response.count"
            :current-page="currentPage"
            :page-size="currentpageSize"
            :page-size-options="[100, 250, 500]"
            @page-change="pageChange"
            @page-size-change="pageSizeChange"
          />
        </div>
        <div
          class="flex flex-col bg-white p-4 rounded-md overflow-auto h-[calc(100vh-11rem)] scrollbar-thin scrollbar-thumb-gray-400 scrollbar-track-white"
        >
          <ContentEmail v-if="rowSelected.message_id" :email="rowSelected"></ContentEmail>
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
    </div>
  </div>
</template>
