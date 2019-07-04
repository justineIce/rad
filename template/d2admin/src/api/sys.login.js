import request from '@/plugin/axios'

export function AccountLogin (data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}
export function AccountExit () {
  return request({
    url: '/exit',
    method: 'get'
  })
}
