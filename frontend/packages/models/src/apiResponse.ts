export default interface APIResponse {
  status: string
  message: string
  data?: any
  serverTime: bigint
}