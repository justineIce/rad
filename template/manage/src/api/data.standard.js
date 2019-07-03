import request from '@/plugin/axios'

// 数据标准类型
export function GetDataStandardTypeList (data) {
  return request({
    url: '/standard/type/list',
    method: 'post',
    data
  })
}
export function AddDataStandardType (data) {
  return request({
    url: '/standard/type/add',
    method: 'post',
    data
  })
}
export function DelDataStandardType (data) {
  return request({
    url: '/standard/type/del',
    method: 'post',
    data
  })
}

// 检核规则
export function GetDataStandardRuleList (data) {
  return request({
    url: '/standard/rule/list',
    method: 'post',
    data
  })
}
export function AddDataStandardRule (data) {
  return request({
    url: '/standard/rule/add',
    method: 'post',
    data
  })
}
export function DelDataStandardRule (data) {
  return request({
    url: '/standard/rule/del',
    method: 'post',
    data
  })
}

// 标准管理
export function AddDataStandard (data) {
  return request({
    url: '/standard/add',
    method: 'post',
    data
  })
}
export function DelDataStandard (data) {
  return request({
    url: '/standard/del',
    method: 'post',
    data
  })
}
export function GetDataStandardList (data) {
  return request({
    url: '/standard/list',
    method: 'post',
    data
  })
}
export function AddDataStandardVersion (data) {
  return request({
    url: '/standard/version/add',
    method: 'post',
    data
  })
}
export function DeleteDataStandardVersion (data) {
  return request({
    url: '/standard/version/del',
    method: 'post',
    data
  })
}
export function SetDataStandardVersionUseStatus (data) {
  return request({
    url: '/standard/version/set',
    method: 'post',
    data
  })
}
export function GetDataStandardCatalog (data) {
  return request({
    url: '/standard/catalog',
    method: 'post',
    data
  })
}

// 标准管理数据项
export function GetDataStandardFieldList (data) {
  return request({
    url: '/standard/field/list',
    method: 'post',
    data
  })
}
export function AddDataStandardField (data) {
  return request({
    url: '/standard/field/add',
    method: 'post',
    data
  })
}
export function DeleteDataStandardField (data) {
  return request({
    url: '/standard/field/del',
    method: 'post',
    data
  })
}

// 审核相关接口
export function SetDataStandardFieldAudit (data) {
  return request({
    url: '/standard/field/audit/set',
    method: 'post',
    data
  })
}
export function GetDataStandardAudit (data) {
  return request({
    url: '/standard/audit/get',
    method: 'post',
    data
  })
}
export function GetDataStandardAuditList (data) {
  return request({
    url: '/standard/audit/list',
    method: 'post',
    data
  })
}
export function SetDataStandardAudit (data) {
  return request({
    url: '/standard/audit/set',
    method: 'post',
    data
  })
}
export function GetDataStandardAuditHistory (data) {
  return request({
    url: '/standard/audit/history',
    method: 'post',
    data
  })
}

// 编码标准
export function AddDataStandardCoding (data) {
  return request({
    url: '/standard/coding/add',
    method: 'post',
    data
  })
}
export function GetDataStandardCodingList (data) {
  return request({
    url: '/standard/coding/list',
    method: 'post',
    data
  })
}
export function SetDataStandardCodingAudit (data) {
  return request({
    url: '/standard/coding/audit/set',
    method: 'post',
    data
  })
}
export function GetDataStandardCodingItems (data) {
  return request({
    url: '/standard/coding/item',
    method: 'post',
    data
  })
}

// 发布版
export function GetPublishStandardFieldList (data) {
  return request({
    url: '/standard/publish/field/list',
    method: 'post',
    data
  })
}
export function GetStandardPublishCodeList (data) {
  return request({
    url: '/standard/publish/code/list',
    method: 'post',
    data
  })
}
export function GetStandardPublishCodeItemList (data) {
  return request({
    url: '/standard/publish/code/item/list',
    method: 'post',
    data
  })
}
export function SetDataStandardFieldCode (data) {
  return request({
    url: '/standard/field/code/set',
    method: 'post',
    data
  })
}
