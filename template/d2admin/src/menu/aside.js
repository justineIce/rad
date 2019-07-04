// 菜单 侧边栏
export default [
    {path: '/index', title: '首页', icon: 'home'},
    {
        path: '/platform-manage',
        title: '平台管理',
        icon: 'window-restore',
        children: (pre => [
            {path: `${pre}userManage`, title: '用户管理'},
            {path: `${pre}roleManage`, title: '角色管理'},
            {path: `${pre}resourceManage`, title: '菜单管理'},
            {path: `${pre}logRecord`, title: '日志审计'},
            {path: `${pre}whitelist`, title: '白名单管理'}
        ])('/platform-manage/')
    }
]
