import layoutHeaderAside from '@/layout/header-aside'

const meta = { auth: true }

export default {
  path: '/platform-manage',
  name: 'platform-manage',
  meta,
  redirect: { name: 'platform-manage' },
  component: layoutHeaderAside,
  children: (pre => [
    { path: 'userManage', name: `${pre}userManage`, component: () => import('@/pages/platform-manage/userManage'), meta: { ...meta, title: '用户管理' } },
    { path: 'roleManage', name: `${pre}roleManage`, component: () => import('@/pages/platform-manage/roleManage'), meta: { ...meta, title: '角色管理' } },
    { path: 'logRecord', name: `${pre}logRecord`, component: () => import('@/pages/platform-manage/logRecord'), meta: { ...meta, title: '日志审计' } },
    { path: 'resourceManage', name: `${pre}resourceManage`, component: () => import('@/pages/platform-manage/resourceManage'), meta: { ...meta, title: '菜单管理' } },
    { path: 'whitelist', name: `${pre}whitelist`, component: () => import('@/pages/platform-manage/whitelist'), meta: { ...meta, title: '白名单管理' } },
  ])('platform-manage-')
}
