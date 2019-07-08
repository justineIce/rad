`import request from '@/plugin/axios'

// 获取单条
export function Get{{.modelInfo.StructName}} (data) {
    return request({
        url: '/auth/{{.modelInfo.SingName}}/get',
        method: 'post',
        data
    })
}
// 获取全部
export function Get{{.modelInfo.StructName}}All (data) {
    return request({
        url: '/auth/{{.modelInfo.SingName}}/get',
        method: 'post',
        data
    })
}
// 获取分页
export function Get{{.modelInfo.StructName}}Page (data) {
    return request({
        url: '/auth/{{.modelInfo.SingName}}/get',
        method: 'post',
        data
    })
}
// 保存
export function Save{{.modelInfo.StructName}} (data) {
    return request({
        url: '/auth/{{.modelInfo.SingName}}/get',
        method: 'post',
        data
    })
}
// 删除
export function Del{{.modelInfo.StructName}} (data) {
    return request({
        url: '/auth/{{.modelInfo.SingName}}/get',
        method: 'post',
        data
    })
}`
