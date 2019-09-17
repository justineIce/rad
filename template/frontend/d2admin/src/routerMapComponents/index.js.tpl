`const files = require.context('./', true, /\.js$/);
 import layoutHeaderAside from '@/layout/header-aside'

 let componentMaps = {
     "layoutHeaderAside": layoutHeaderAside,
     "index": () => import('@/pages/index'),
     "platformManage_resourceManage": () => import('@/pages/platformManage/resourceManage'),
     "platformManage_userManage": () => import('@/pages/platformManage/userManage'),
     "platformManage_roleManage": () => import('@/pages/platformManage/roleManage'),
     "platformManage_logRecord": () => import('@/pages/platformManage/logRecord'),
     {{ range .tableList }}"Business_{{.SingName}}Index": () => import('@/pages/{{.SingName}}/index'),
     {{end}}
 }
 files.keys().forEach((key) => {
     if (key === './index.js') return
     Object.assign(componentMaps, files(key).default)
 })
 export default componentMaps`
