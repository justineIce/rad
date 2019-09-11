import Vue from 'vue'
import VueRouter from 'vue-router'
// 进度条
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

import store from '@/store/index'

import util from '@/libs/util.js'
// 路由数据
import routes from './routes'
// 固定菜单与路由
import menuHeader from '@/menu/header'
import menuAside from '@/menu/aside'
import {frameInRoutes} from '@/router/routes'
//路由与组件映射关系
import routerMapComponents from '@/routerMapComponents'
//代码生成器生成的菜单与路由
import autoGenerateMenusAndRouters from '@/development'
import {GetLoginMenu} from '@/api/platform.manage';

Vue.use(VueRouter)

// 导出路由 在 main.js 里使用
const router = new VueRouter({
    routes
})

let permissionMenu,permissionTopMenu, permissionRouter = [], allMenuAside, allMenuHeader

let permission = {
    functions: [],
    roles: [],
    isAdmin: false
}

//标记是否已经拉取权限信息
let isFetchPermissionInfo = false

let fetchPermissionInfo = async () => {
    //处理动态添加的路由
    const formatRoutes = function (routes) {
        routes.forEach(route => {
            route.component = routerMapComponents[route.component]
            if (route.children) {
                formatRoutes(route.children)
            }
        })
    }
    let userPermissionInfo
    /*let userPermissionInfo = localStorage.getItem('dynamic_router')
    if (!userPermissionInfo) {*/
    try {
        let result = await GetLoginMenu()
        if (result.ret === 200) {
            userPermissionInfo = result.data.loginInfo
            localStorage.setItem('dynamic_router', JSON.stringify(userPermissionInfo))
        }
    } catch (ex) {
        console.log(ex)
    }
    /*}else{
        userPermissionInfo=JSON.parse(userPermissionInfo)
    }*/
    if (userPermissionInfo) {
        permissionMenu = getMenuThree(userPermissionInfo.SysMenu, '0')
        permissionTopMenu = getMenuThree(userPermissionInfo.SysMenu.filter(obj=>obj.ParentID==='0'), '0')
        permissionRouter = getRouterThree(userPermissionInfo.SysMenu, '0')
        permission.functions = getMenuBtn(userPermissionInfo.SysMenuBtns)
        permission.roles = userPermissionInfo.SysRoles
        permission.interfaces = util.formatInterfaces(getInterfaces(userPermissionInfo.SysMenuBtns)) // util.formatInterfaces(userPermissionInfo.accessInterfaces)
        permission.isAdmin = userPermissionInfo.IsSuperAdmin

        //组合代码生成器生成的菜单和路由
        permissionMenu = permissionMenu == null ? null : [...permissionMenu, ...autoGenerateMenusAndRouters.menus]
        permissionTopMenu = permissionTopMenu == null ? [...menuAside] : [...menuAside, ...permissionTopMenu, ...autoGenerateMenusAndRouters.menus]
        permissionRouter = permissionMenu == null ? [] : [...permissionRouter, ...autoGenerateMenusAndRouters.routers]
        allMenuAside = permissionMenu == null ? [...menuAside] : [...menuAside, ...permissionMenu]
        allMenuHeader = permissionMenu == null ? [...menuHeader] : [...menuHeader, ...permissionMenu]
        formatRoutes(permissionRouter)
    }
    permissionRouter.push(
        // 404
        {
            path: '*',
            name: '404',
            component: () => import('@/pages/error-page-404')
        }
    )
    //动态添加路由
    router.addRoutes(permissionRouter);
    if (permissionRouter != null) {
        // 处理路由 得到每一级的路由设置
        store.commit('d2admin/page/init', [...frameInRoutes, ...permissionRouter])
    } else {
        // 处理路由 得到每一级的路由设置
        store.commit('d2admin/page/init', [...frameInRoutes])
    }
    // 设置顶栏菜单
    store.commit('d2admin/menu/headerSet', permissionTopMenu)
    // 设置侧边栏菜单
    store.commit('d2admin/menu/fullAsideSet', allMenuAside)
    // 初始化菜单搜索功能
    store.commit('d2admin/search/init', allMenuHeader)
    // 设置权限信息
    store.commit('d2admin/permission/set', permission)
    // 加载上次退出时的多页列表
    store.dispatch('d2admin/page/openedLoad')
    await Promise.resolve()
}

//组装菜单
function getMenuThree(list, parentId) {
    if (list == null) return null
    let newList = list.filter(item => item.ParentID === parentId)
    if (newList.length > 0) {
        return newList.map(item => {
            var children = getMenuThree(list, item.ID)
            var m = {
                title: item.Name,
                path: item.Href,
                icon: item.Icon !== null ? item.Icon.replace('fa fa-', '') : null
            }
            if (children) {
                m.children = children
            }
            return m
        })
    }
    return null
}

//组装路由
function getRouterThree(list, parentId) {
    if (list == null) return null
    let newList = list.filter(item => item.ParentID === parentId)
    if (newList.length > 0) {
        return newList.map(item => {
            var children = getRouterThree(list, item.ID)
            var r = {
                name: item.RouteName,
                path: item.Href,
                component: item.Component,
                componentPath: item.Target,
                meta: {
                    'title': item.Name,
                    'cache': true,
                    'auth': true
                }
            }
            if (children) {
                r.children = children
                r.redirect = children[0].path
            }
            return r
        })
    }
    return null
}

//组装方法
function getMenuBtn(list) {
    if (list == null) return null
    return list.map(item => item.Permission)
}

//组装接口
function getInterfaces(list) {
    if (list == null) return null
    let arr = new Array()
    for (let i = 0; i < list.length; i++) {
        if (list[i].Path !== '') {
            let str = list[i].Path.split(',')
            for (let j = 0; j < str.length; j++) {
                arr.push({path: str[j], method: list[i].Method.toUpperCase()})
            }
        }
    }
    return arr
}


//免校验token白名单
let whiteList = ['/login']

/**
 * 路由拦截
 * 权限验证
 */
router.beforeEach(async (to, from, next) => {
    // 进度条
    NProgress.start()
    // 关闭搜索面板
    store.commit('d2admin/search/set', false)
    // 验证当前路由所有的匹配中是否需要有登录验证的
    if (whiteList.indexOf(to.path) === -1) {
        // 这里暂时将cookie里是否存有token作为验证是否登录的条件
        // 请根据自身业务需要修改
        const token = util.cookies.get('token')
        if (token && token !== 'undefined') {
            //拉取权限信息
            if (!isFetchPermissionInfo) {
                await fetchPermissionInfo();
                isFetchPermissionInfo = true;
                next({...to, replace: true})
            } else {
                next()
            }
        } else {
            // 没有登录的时候跳转到登录界面
            // 携带上登陆成功之后需要跳转的页面完整路径
            next({
                name: 'login',
                query: {
                    redirect: to.fullPath
                }
            })
            // https://github.com/d2-projects/d2-admin/issues/138
            NProgress.done()
        }
    } else {
        // 不需要身份校验 直接通过
        next()
    }
})

router.afterEach(to => {
    // 进度条
    NProgress.done()
    // 多页控制 打开新的页面
    // store.dispatch('d2admin/page/open', to)
    // 更改标题
    util.title(to.meta.title)
})

export default router
