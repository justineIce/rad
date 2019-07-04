import util from '@/libs/util'

export default {
  install (Vue, options) {
    Vue.prototype.$open = util.open
    Vue.prototype.$objectAssign = util.objectAssign
    Vue.prototype.$formatDate = util.formatDate
    Vue.prototype.$dataType = function () {
      return util.dataType()
    }
  }
}
