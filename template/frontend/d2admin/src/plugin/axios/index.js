import vm from '@/main'
import store from '@/store'
import axios from 'axios'
import qs from 'qs'
import { Message } from 'element-ui'
import util from '@/libs/util'

// 创建一个错误
function errorCreate (msg) {
  const err = new Error(msg)
  errorLog(err)
  throw err
}

// 记录和显示错误
function errorLog (err) {
  // 添加到日志
  store.dispatch('d2admin/log/add', {
    type: 'error',
    err,
    info: '数据请求异常'
  })
  // 打印到控制台
  if (process.env.NODE_ENV === 'development') {
    util.log.danger('>>>>>> Error >>>>>>')
    console.log(err)
  }
  // 显示提示
  Message({
    message: err.message,
    type: 'error',
    duration: 5 * 1000
  })
}

// 创建一个 axios 实例
const service = axios.create({
  baseURL: process.env.VUE_APP_API,
  timeout: 5000, // 请求超时时间
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
    'Access-Token': util.cookies.get('token'),
    'Content-Type': 'application/x-www-form-urlencoded;application/json; charset=UTF-8;'
  },
  responseType: 'json',
  transformRequest: [function (data) {
    data = util.toLowerCase(data)
    return qs.stringify(data)
  }]
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 在请求发送之前做一些处理
    const token = util.cookies.get('token')
    // 让每个请求携带token-- ['X-Token']为自定义key 请根据实际情况自行修改
    config.headers['Access-Token'] = token
    return config
  },
  error => {
    // 发送失败
    console.log(error)
    Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    // dataAxios 是 axios 返回数据中的 data
    const dataAxios = response.data
    // 这个状态码是和后端约定的
    const { ret, msg } = dataAxios
    if (ret === undefined) {
      return dataAxios
    } else {
      switch (ret) {
        case 200:
          return dataAxios
        case 401:
          Message({
            message: '请重新登陆',
            type: 'error',
            duration: 2000,
            onClose: () => {
              if (vm.$router.history.current.name !== 'login') {
                vm.$router.replace({ name: 'login' })
              }
            }
          })
          return null
        default:
          Message({
            message: msg,
            type: 'error',
            duration: 2000
          })
          // 不是正确的 code
          // errorCreate(`${dataAxios.msg}: ${response.config.url}`)
          errorCreate(`${dataAxios.msg}`)
          break
      }
    }
  },
  error => {
    if (error && error.response) {
      switch (error.response.status) {
        case 400: error.message = '请求错误'; break
        case 401: error.message = '未授权，请登录'; break
        case 403: error.message = '拒绝访问'; break
        case 404: error.message = `请求地址出错: ${error.response.config.url}`; break
        case 408: error.message = '请求超时'; break
        case 500: error.message = '服务器内部错误'; break
        case 501: error.message = '服务未实现'; break
        case 502: error.message = '网关错误'; break
        case 503: error.message = '服务不可用'; break
        case 504: error.message = '网关超时'; break
        case 505: error.message = 'HTTP版本不受支持'; break
        default: break
      }
    }
    errorLog(error)
    return Promise.reject(error)
  }
)

export default service
