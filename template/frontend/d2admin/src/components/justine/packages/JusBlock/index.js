import JustineBlock from './JustineBlock'

JustineBlock.install = (Vue) => {
  Vue.component(JustineBlock.name, JustineBlock)
}
if (typeof window !== 'undefined' && window.Vue) {
  JustineBlock.install(window.Vue)
}
export default JustineBlock
