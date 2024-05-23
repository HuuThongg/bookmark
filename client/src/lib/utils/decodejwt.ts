import { Buffer } from 'buffer'

export function decodeJwt(token: string) {
  const base64payload = token.split('.')[1]
  const payloadBuffer = Buffer.from(base64payload, 'base64')
  return JSON.parse(payloadBuffer.toString())
}
