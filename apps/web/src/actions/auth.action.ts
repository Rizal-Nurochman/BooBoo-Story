import { AuthService } from '@/services/auth.service'
import { createServerFn } from '@tanstack/react-start'
import { getCookie } from '@tanstack/react-start/server'

export const fetchUserLogin=createServerFn({ method:'GET' }).handler(async()=>{
  const cookie = getCookie('access_token') 
  const { data } =await AuthService.me(cookie)

  return{
    user:data
  }
})
