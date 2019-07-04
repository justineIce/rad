import Layout from './src/index'

Layout.install = function (Vue) {
  Vue.component(Layout.name, Layout)
}
export default Layout
