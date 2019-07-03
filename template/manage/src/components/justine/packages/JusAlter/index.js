import Alert from './src/alert'

Alert.install = (Vue) => {
  Vue.component(Alert.name, Alert)
}
if (typeof window !== 'undefined' && window.Vue) {
  Alert.install(window.Vue)
}
export default Alert
