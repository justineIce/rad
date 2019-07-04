import Table from './src/table'

Table.install = (Vue) => {
  Vue.component(Table.name, Table)
}
// if (typeof window !== 'undefined' && window.Vue) {
//   button.install(window.Vue)
// }
export default Table
