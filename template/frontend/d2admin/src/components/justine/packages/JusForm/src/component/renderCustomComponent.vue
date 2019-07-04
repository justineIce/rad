<script>
export default {
  name: 'render-custom-component',
  props: {
    /**
       * @description v-model双向绑定的值
       */
    value: {
      type: String,
      required: true
    },
    /**
       * @description 传入的组件名
       */
    componentName: {
      required: true
    },
    /**
       * @description 传入的自定义参数
       */
    props: {
      default: null
    },
    /**
       * @description 传入的行数据
       */
    scope: {
      default: null
    }
  },
  render (h) {
    let self = this
    console.log(self.value)
    return h(self.componentName, {
      props: {
        value: self.value,
        scope: self.scope,
        ...self.props
      },
      on: {
        input: function (event) {
          self.$emit('input', event)
        },
        change: function (event) {
          self.$emit('change', event)
        }
      }
    })
  }
}
</script>

<style scoped>

</style>
