import { api } from '@/api/api'
import type { IEmailResponse, IParams } from '@/interfaces'

export default {
  getAllEmails(params: IParams) {
    return api.get<IEmailResponse>('/email', { params: params })
  },
}
