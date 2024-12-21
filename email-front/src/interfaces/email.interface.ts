export interface IEmailResponse {
  count: number
  results: IEmail[]
}

export interface IEmail {
  _timestamp: number
  body: string
  date: string
  from: string
  message_id: string
  subject: string
  to: string
}
