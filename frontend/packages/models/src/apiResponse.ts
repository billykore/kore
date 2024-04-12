export enum Status {
  OK = "OK"
}

export interface APIResponse {
  status: string
  message: string
  data?: any
  serverTime: bigint
}
