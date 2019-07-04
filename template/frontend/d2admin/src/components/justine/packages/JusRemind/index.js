import Remind from './src/index'
Remind.install = (Vue) => {
  Vue.component(Remind.name, Remind)
}
export default Remind
