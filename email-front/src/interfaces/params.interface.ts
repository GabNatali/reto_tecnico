export interface IParams {
  fromEmail: string
  to: string
  subject: string
  from: string
  size: string
  start_time: string
  end_time: string
  stream_log: string
}


export type TimeUnit = 'minute' | 'hour' | 'day' | 'week';
