import request from '@/plugin/axios'

/**
 * @description 存储数据库的相关接口
 * @param data
 * @constructor
 */
export function GetDatabaseList (data) {
  return request({
    url: '/storage/database/list',
    method: 'post',
    data
  })
}
export function GetStorageResourceList (data) {
  return request({
    url: '/storage/resource/list',
    method: 'post',
    data
  })
}
