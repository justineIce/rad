import request from '@/plugin/axios'

// 管理平台 - 获取顶级菜单
export function GetTopMenuList (data) {
  return request({
    url: '/platformManage/resource/menu/top/list',
    method: 'post',
    data
  })
}

// 管理平台 - 添加/修改顶级菜单
export function AddTopLevelMenu (data) {
  return request({
    url: '/platformManage/resource/menu/add',
    method: 'post',
    data
  })
}

// 管理平台 - 删除顶级菜单
export function DelTopLevelMenu (data) {
  return request({
    url: '/platformManage/resource/menu/del',
    method: 'post',
    data
  })
}

// 管理平台 - 获取所有菜单
export function GetAllMenuList (data) {
  return request({
    url: '/platformManage/resource/menu/all/list',
    method: 'post',
    data
  })
}

// 管理平台 - 根据菜单id获取按钮操作权限
export function GetBtnMenuList (data) {
  return request({
    url: '/platformManage/resource/menu/btn/list',
    method: 'post',
    data
  })
}

// 获取所有的菜单列表
export function GetMenuList (data) {
  return request({
    url: '/platformManage/resource/menu/list',
    method: 'post',
    data
  })
}

/* -----------角色管理------------- */

// 获取角色列表
export function GetRoleList (data) {
  return request({
    url: '/platformManage/role/list',
    method: 'post',
    data
  })
}

// 添加角色
export function AddRole (data) {
  return request({
    url: '/platformManage/role/add',
    method: 'post',
    data
  })
}

// 删除角色
export function DelRole (data) {
  return request({
    url: '/platformManage/role/del',
    method: 'post',
    data
  })
}

// 删除角色
export function GetCheckedMenuBtnLimit (data) {
  return request({
    url: '/platformManage/role/resource/v/list',
    method: 'post',
    data
  })
}

// 操作角色资源关系表
export function AddRoleResource (data) {
  return request({
    url: '/platformManage/role/resource/add',
    method: 'post',
    data
  })
}

/* -----------------用户管理--------------------- */

// 添加用户
export function AddSystemManage (data) {
  return request({
    url: '/platformManage/system/manage/add',
    method: 'post',
    data
  })
}

// 获取用户列表
export function GetAuthorityList (data) {
  return request({
    url: '/platformManage/system/manage/list',
    method: 'post',
    data
  })
}

// 获取用户列表
export function DelSystemManage (data) {
  return request({
    url: '/platformManage/system/manage/del',
    method: 'post',
    data
  })
}

// 重置密码
export function ResetPassword (data) {
  return request({
    url: '/platformManage/system/manage/password/reset',
    method: 'post',
    data
  })
}

// 根据多个菜单id获取其子按钮资源
export function GetBtnMenuListByIds (data) {
  return request({
    url: '/platformManage/resource/menu/btn/list/by/ids',
    method: 'post',
    data
  })
}

/* ------------------白名单管理----------------- */

// 获取白名单列表
export function GetIpList (data) {
  return request({
    url: '/platformManage/ip/list',
    method: 'post',
    data
  })
}

// 添加白名单
export function AddIp (data) {
  return request({
    url: '/platformManage/ip/add',
    method: 'post',
    data
  })
}

// 删除白名单
export function DelIp (data) {
  return request({
    url: '/platformManage/ip/del',
    method: 'post',
    data
  })
}

/* ------------操作日志---------------- */

// 获取操作列表
export function GetManageActionLog (data) {
  return request({
    url: '/platformManage/system/manage/action/log/list',
    method: 'post',
    data
  })
}
