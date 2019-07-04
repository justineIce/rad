import Form from './src/form'

Form.install = (Vue) => {
  Vue.component(Form.name, Form)
}
// if (typeof window !== 'undefined' && window.Vue) {
//   button.install(window.Vue)
// }
export default Form
