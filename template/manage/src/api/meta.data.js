import request from '@/plugin/axios'
/**
 * @description 分类
 */
export function GetMetaDataCatalogList (data) {
  return request({
    url: '/metadata/catalog/list',
    method: 'post',
    data
  })
}
export function AddMetaDataCatalog (data) {
  return request({
    url: '/metadata/catalog/add',
    method: 'post',
    data
  })
}
export function DeleteMetaDataCatalog (data) {
  return request({
    url: '/metadata/catalog/del',
    method: 'post',
    data
  })
}
/**
 * @description 元数据
 */
export function GetMetaDataList (data) {
  return request({
    url: '/metadata/list',
    method: 'post',
    data
  })
}
export function DeleteMetaData (data) {
  return request({
    url: '/metadata/del',
    method: 'post',
    data
  })
}
export function AddMetaDataGeneral (data) {
  return request({
    url: '/metadata/general/add',
    method: 'post',
    data
  })
}
export function AddMetaDataField (data) {
  return request({
    url: '/metadata/field/add',
    method: 'post',
    data
  })
}
export function AddMetaDataIndex (data) {
  return request({
    url: '/metadata/index/add',
    method: 'post',
    data
  })
}
export function AddMetaDataAudit (data) {
  return request({
    url: '/metadata/audit/add',
    method: 'post',
    data
  })
}
export function GetMetaDataFieldRuleList (data) {
  return request({
    url: '/metadata/field/rule/list',
    method: 'post',
    data
  })
}

export function AddMetaData (data) {
  return request({
    url: '/metadata/add',
    method: 'post',
    data
  })
}
export function GetMetaDataById (data) {
  return request({
    url: '/metadata/get',
    method: 'post',
    data
  })
}
/**
 * @description 发布版元数据
 */
export function GetPublishMetaDataList (data) {
  return request({
    url: '/metadata/publish/list',
    method: 'post',
    data
  })
}
export function GetPublishMetaData (data) {
  return request({
    url: '/metadata/publish/get',
    method: 'post',
    data
  })
}
export function GetPublishMetaDataFieldList (data) {
  return request({
    url: '/metadata/publish/field/list',
    method: 'post',
    data
  })
}
export function GetPublishMetaDataFieldRuleList (data) {
  return request({
    url: '/metadata/publish/field/rule/list',
    method: 'post',
    data
  })
}
/**
 * @description 审核相关的结构
 */
export function GetMetaDataAuditList (data) {
  return request({
    url: '/metadata/audit/list',
    method: 'post',
    data
  })
}
export function GetMetaDataAuditInfo (data) {
  return request({
    url: '/metadata/audit/get',
    method: 'post',
    data
  })
}
export function SetMetaDataAuditStatus (data) {
  return request({
    url: '/metadata/audit/set',
    method: 'post',
    data
  })
}
export function GetMetaDataAuditHistory (data) {
  return request({
    url: '/metadata/audit/history',
    method: 'post',
    data
  })
}
