import request from '@/plugin/axios'

/**
 * @description 项目管理的相关接口
 * @param data
 * @constructor
 */
export function AddProject (data) {
  return request({
    url: '/govern/project/add',
    method: 'post',
    data
  })
}
export function GetProjectList (data) {
  return request({
    url: '/govern/project/list',
    method: 'post',
    data
  })
}
export function DeleteProject (data) {
  return request({
    url: '/govern/project/del',
    method: 'post',
    data
  })
}
export function GetProject (data) {
  return request({
    url: '/govern/project/get',
    method: 'post',
    data
  })
}
export function GetMetadata (data) {
  return request({
    url: '/govern/metadata/list',
    method: 'post',
    data
  })
}

/**
 * @description 任务相关接口
 */
export function GetTaskList (data) {
  return request({
    url: '/govern/task/list',
    method: 'post',
    data
  })
}
export function AddTask (data) {
  return request({
    url: '/govern/task/add',
    method: 'post',
    data
  })
}
export function DeleteTask (data) {
  return request({
    url: '/govern/task/del',
    method: 'post',
    data
  })
}
/**
 * @description 模块的相关接口
 * @param data
 * @constructor
 */
export function AddModule (data) {
  return request({
    url: '/govern/module/add',
    method: 'post',
    data
  })
}
export function GetModuleList (data) {
  return request({
    url: '/govern/module/list',
    method: 'post',
    data
  })
}
export function SaveModule (data) {
  return request({
    url: '/govern/module/save',
    method: 'post',
    data
  })
}
export function GetModule (data) {
  return request({
    url: '/govern/module/get',
    method: 'post',
    data
  })
}
export function SaveModuleInput (data) {
  return request({
    url: '/govern/module/input/save',
    method: 'post',
    data
  })
}
export function GetModuleInput (data) {
  return request({
    url: '/govern/module/input/get',
    method: 'post',
    data
  })
}
export function GetModuleStandardize (data) {
  return request({
    url: '/govern/module/standardize/get',
    method: 'post',
    data
  })
}
export function SaveModuleStandardize (data) {
  return request({
    url: '/govern/module/standardize/save',
    method: 'post',
    data
  })
}
export function GetModuleDatabaseList (data) {
  return request({
    url: '/govern/module/database/list',
    method: 'post',
    data
  })
}
export function GetModuleDatabaseTableList (data) {
  return request({
    url: '/govern/module/database/table/list',
    method: 'post',
    data
  })
}
