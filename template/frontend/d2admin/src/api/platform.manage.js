import request from '@/plugin/axios'

// 管理平台 - 添加/修改顶级菜单
export function SaveMenu (data) {
    return request({
        url: '/auth/sys_menu/save',
        method: 'post',
        data
    })
}

// 管理平台 - 删除顶级菜单
export function DelMenu (data) {
    return request({
        url: '/auth/sys_menu/del',
        method: 'post',
        data
    })
}


// 管理平台 - 获取菜单按钮
export function GetBtnMenuList (data) {
    return request({
        url: '/auth/sys_menu_btn/all',
        method: 'post',
        data
    })
}

// 管理平台 - 保存菜单按钮
export function SaveMenuBtn (data) {
    return request({
        url: '/auth/sys_menu_btn/save',
        method: 'post',
        data
    })
}

// 管理平台 - 删除菜单按钮
export function DelMenuBtn (data) {
    return request({
        url: '/auth/sys_menu_btn/del',
        method: 'post',
        data
    })
}

// 获取所有的菜单列表
export function GetMenuList (data) {
    return request({
        url: '/auth/sys_menu/all',
        method: 'post',
        data
    })
}

/* -----------角色管理------------- */

// 获取角色列表
export function GetRoleList (data) {
    return request({
        url: '/auth/sys_role/all',
        method: 'post',
        data
    })
}

// 添加角色
export function AddRole (data) {
    return request({
        url: '/auth/sys_role/save',
        method: 'post',
        data
    })
}

// 删除角色
export function DelRole (data) {
    return request({
        url: '/auth/sys_role/del',
        method: 'post',
        data
    })
}

// 根据角色获取菜单资源
export function GetCheckedMenuLimit (data) {
    return request({
        url: '/auth/sys_role_menu/all',
        method: 'post',
        data
    })
}
// 根据角色获取菜单资源按钮
export function GetCheckedMenuBtnLimit (data) {
    return request({
        url: '/auth/sys_role_menu_btn/all',
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
        url: '/auth/sys_user/save',
        method: 'post',
        data
    })
}

// 获取用户列表
export function GetAuthorityList (data) {
    return request({
        url: '/auth/sys_user/page',
        method: 'post',
        data
    })
}

// 获取用户列表
export function DelSystemManage (data) {
    return request({
        url: '/auth/sys_user/del',
        method: 'post',
        data
    })
}

// 重置密码
export function ResetPassword (data) {
    return request({
        url: '/auth/sys_user/pwd/reset',
        method: 'post',
        data
    })
}

// 添加菜单权限
export function UpdateRoleMenuByIds (data) {
    return request({
        url: '/auth/sys_role_menu/save',
        method: 'post',
        data
    })
}
// 添加菜单按钮权限
export function UpdateRoleMenuBtnById (data) {
    return request({
        url: '/auth/sys_role_menu_btn/save',
        method: 'post',
        data
    })
}

/* ------------------部门管理----------------- */

export function SaveSysOffice (data) {
    return request({
        url: '/auth/sys_office/save',
        method: 'post',
        data
    })
}

export function GetSysOfficeAll (data) {
    return request({
        url: '/auth/sys_office/all',
        method: 'post',
        data
    })
}

export function DelSysOffice (data) {
    return request({
        url: '/auth/sys_office/del',
        method: 'post',
        data
    })
}

/* ------------操作日志---------------- */

// 获取操作列表
export function GetManageActionLog (data) {
    return request({
        url: '/auth/sys_log/page',
        method: 'post',
        data
    })
}

/* ------------登陆者权限---------------- */

// 登录的菜单资源
export function GetLoginMenu (data) {
    return request({
        url: '/auth/login/menu',
        method: 'post',
        data
    })
}
