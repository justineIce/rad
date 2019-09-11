import request from '@/plugin/axios'

export function AccountLogin (data) {
    return request({
        url: '/login',
        method: 'post',
        data
    })
}
export function AccountLogOut () {
    return request({
        url: '/logout',
        method: 'post'
    })
}
