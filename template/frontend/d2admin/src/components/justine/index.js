import Vue from 'vue'
Vue.component('JusLayout', () => import('./packages/JusLayout'))
Vue.component('JusForm', () => import('./packages/JusForm'))
Vue.component('JusTable', () => import('./packages/JusTable'))
Vue.component('JusRemind', () => import('./packages/JusRemind'))
Vue.component('JusTree', () => import('./packages/JusTree'))
Vue.component('JusTreeNode', () => import('./packages/JusTree/src/tree-node'))
Vue.component('JusDragList', () => import('./packages/drag-list'))
Vue.component('CountTo', () => import('./packages/count-to'))
Vue.component('CountCard', () => import('./packages/count-card'))

Vue.component('JustineBlock', () => import('./packages/JusBlock'))

Vue.component('JusAlert', () => import('./packages/JusAlter/index'))
Vue.component('JusCard', () => import('./packages/JusCard/index'))

Vue.component('el-icon-popover', () => import('./packages/el-icon-popover/index'))
Vue.component('demo-page-footer', () => import('./packages/demo-page-footer/index'))
Vue.component('el-select-label', () => import('./packages/el-select-label/index'))
