import axios from 'axios'

export const eDNAAxios = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://localhost:8080',
})
