export default {
  methods: {
    /**
         * @description 组件属性默认值
         */
    handleAttribute (attribute, defaultValue) {
      if (attribute === false || attribute === 0 || attribute === '') {
        return attribute
      }
      return attribute || defaultValue
    },
    handleAttributeDisabled (disabled = false, data) {
      if (typeof disabled === 'boolean') {
        return disabled
      } else if (typeof disabled === 'function') {
        return disabled(data)
      }
      return Boolean(disabled)
    }
  }
}
