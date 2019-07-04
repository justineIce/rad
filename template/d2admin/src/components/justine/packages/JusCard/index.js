import Card from './src/card'

Card.install = (Vue) => {
  Vue.component(Card.name, Card)
}
if (typeof window !== 'undefined' && window.Vue) {
  Card.install(window.Vue)
}
export default Card
