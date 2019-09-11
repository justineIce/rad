`const files = require.context('./', true, /\.js$/);
 import layoutHeaderAside from '@/layout/header-aside'

 let componentMaps = {
     "layoutHeaderAside": layoutHeaderAside,
     "index": () => import('@/pages/index'),
     "platformManage_resourceManage": () => import(/* webpackChunkName: "menu" */'@/pages/platformManage/resourceManage'),
     "platformManage_userManage": () => import(/* webpackChunkName: "menu" */'@/pages/platformManage/userManage'),
     "platformManage_roleManage": () => import(/* webpackChunkName: "menu" */'@/pages/platformManage/roleManage'),
     "platformManage_logRecord": () => import(/* webpackChunkName: "menu" */'@/pages/platformManage/logRecord'),
     {{ range .tableList }}"{{.SingName}}_index": () => import('@/pages/{{.SingName}}/index'),
     {{end}}
 }
 files.keys().forEach((key) => {
     if (key === './index.js') return
     Object.assign(componentMaps, files(key).default)
 })
 export default componentMaps`
