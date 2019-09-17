import cookies from './util.cookies'
import db from './util.db'
import log from './util.log'
import { parse } from 'matchit'

const util = {
    cookies,
    db,
    log
}

const loading = null

/**
 * @description 更新标题
 * @param {String} title 标题
 */
util.title = function (titleText) {
    const processTitle = process.env.VUE_APP_TITLE || 'govern'
    window.document.title = `${processTitle}${titleText ? ` | ${titleText}` : ''}`
}

/**
 * @description 打开新页面
 * @param {String} url 地址
 */
util.open = function (url) {
    var a = document.createElement('a')
    a.setAttribute('href', url)
    a.setAttribute('target', '_blank')
    a.setAttribute('id', 'd2admin-link-temp')
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(document.getElementById('d2admin-link-temp'))
}
/**
 * @description json 首页字母小写
 * @param {object} data
 */
util.toLowerCase = function (data) {
    let obj = {}
    if (data instanceof Object) {
        for (var key in data) {
            if (key !== 'ID') {
                var _key = key.substring(0, 1).toLowerCase() + key.substring(1)
                _key = _key.replace(/([A-Z])/g, '_$1').toLowerCase()
                if (data[key] instanceof Array) {
                    var arr = []
                    data[key].forEach((item) => {
                        if (data instanceof Object) {
                            arr.push(util.toLowerCase(item))
                        } else {
                            arr.push(item)
                        }
                    })
                    obj[_key] = JSON.stringify(arr)
                } else {
                    obj[_key.toLowerCase()] = data[key]
                }
            } else {
                obj[key.toLowerCase()] = data[key]
            }
        }
        return obj
    }
    return data
}
/**
 * objectAssign
 * @param target
 * @return {*}
 */
util.objectAssign = (target, json) => {
    var jsonObj = {}
    for (var key in target) {
        jsonObj[key] = target[key]
    }
    return Object.assign(jsonObj, json)
}
/**
 格式化时间
 */
util.formatDate = (date, fmt) => {
    if (date === '' || date === null || date === undefined) {
        return null
    }
    if (fmt === '' || fmt == null || fmt === undefined) {
        fmt = 'yyyy-MM-dd hh:mm:ss'
    }
    date = new Date(date)
    var o = {
        'M+': date.getMonth() + 1, // 月份
        'd+': date.getDate(), // 日
        'h+': date.getHours(), // 小时
        'm+': date.getMinutes(), // 分
        's+': date.getSeconds(), // 秒
        'q+': Math.floor((date.getMonth() + 3) / 3), // 季度
        'S': date.getMilliseconds() // 毫秒
    }
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
    for (var k in o) {
        if (new RegExp('(' + k + ')').test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length)))
    }
    return fmt
}
/**
 数据类型
 */
util.dataType = () => {
    return [
        {label: '字符型', value: 'varchar'},
        {label: '数字型', value: 'number'},
        {label: '币值型', value: 'money'},
        {label: '二进制类型', value: 'binary'},
        {label: '文本型', value: 'text'},
        {label: '时间型', value: 'datetime'}
    ]
}

/**
 *  @description 获取url参数
 *  @variable {String} 参数名称
 *  @return {String} 参数值/‘’
 */
util.getQueryVariable = function (name) {
    var reg = new RegExp('(^|\\?|&)' + name + '=([^&]*)(\\s|&|$)', 'i')
    if (reg.test(location.href)) {
        return decodeURIComponent(RegExp.$2.replace(/\+/g, ' '))
    }
    return ''
}
/**
 * @description 处理接口
 * @param {Array} interfaces 接口
 */
util.formatInterfaces = function (interfaces) {
    let i = {}
    if (interfaces == null || interfaces.length <= 0) return i
    i["GET"] = interfaces.filter(s => s.method === "GET").map(s => parse(s.path))
    i["POST"] = interfaces.filter(s => s.method === "POST").map(s => parse(s.path))
    i["PUT"] = interfaces.filter(s => s.method === "PUT").map(s => parse(s.path))
    i["DELETE"] = interfaces.filter(s => s.method === "DELETE").map(s => parse(s.path))
    return i
}
export default util
