// Vue
import Vue from 'vue'
import App from './App'
// store
import store from '@/store/index'
// 多国语
import i18n from './i18n'
// 核心插件
import d2Admin from '@/plugin/d2admin'

// 可选插件
// 右键菜单
import contentmenu from 'v-contextmenu'
import 'v-contextmenu/dist/index.css'

// 菜单和路由设置
import router from './router'
import menuHeader from '@/menu/header'
import menuAside from '@/menu/aside'
import { frameInRoutes } from '@/router/routes'

// 自定义菜单组件
import navMenu from './components/navmenu/index.js'
import elTreeSelect from './components/el-tree-select/index.js'

// [可选插件组件]v-charts
import VCharts from 'v-charts'

// 代码编辑器
import { codemirror } from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'

// 核心插件
Vue.use(d2Admin)

// 可选插件
Vue.use(contentmenu)

Vue.use(VCharts)

// 自定义菜单组件
Vue.use(navMenu)
// 文件选择组件，元素:el-tree-select
Vue.use(elTreeSelect)

Vue.use(codemirror)

export default new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
  created () {
    // 处理路由 得到每一级的路由设置
    this.$store.commit('d2admin/page/init', frameInRoutes)
    // 设置顶栏菜单
    this.$store.commit('d2admin/menu/headerSet', menuHeader)
    // 设置侧边栏菜单
    this.$store.commit('d2admin/menu/asideSet', menuAside)
  },
  mounted () {
    // 展示系统信息
    this.$store.commit('d2admin/releases/versionShow')
    // 用户登录后从数据库加载一系列的设置
    this.$store.dispatch('d2admin/account/load')
    // 获取并记录用户 UA
    this.$store.commit('d2admin/ua/get')
    // 初始化全屏监听
    this.$store.dispatch('d2admin/fullscreen/listen')
  }
}).$mount('#app')
