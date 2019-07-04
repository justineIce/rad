<template>
  <i class="iconfont" :class="`${name}`" :style="{color:color} " aria-hidden="true"></i>
</template>

<script>
import './iconfont/iconfont.css'
export default {
  name: 'icon-icon',
  props: {
    name: {
      type: String,
      required: false,
      default: 'icon-tianjia'
    },
    color: String
  }
}
</script>
