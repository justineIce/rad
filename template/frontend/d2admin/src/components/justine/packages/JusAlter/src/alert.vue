<!--alert提示框-->
<template>
    <transition name="fade">
        <div class="jus-alert" role="alert" v-show="visible" :class="[alertClass]">
            <!--图标-->
            <i class="jus-alert__icon" :class="[iconClass,isBigIcon]" v-if="showIcon"></i>
            <div class="jus-alert__content">
                <!--标题-->
                <span class="jus-alert__title" v-if="title || $slots.title">
                    <slot name="title">{{title}}</slot>
                </span>
                <!--描述-->
                <slot>
                    <p class="jus-alert__description" v-if="description">{{description}}</p>
                </slot>
                <!--按钮-->
                <i class="jus-alert__closebtn" v-show="closable" @click="close()">{{closeText ? closeText : 'x' }}</i>
            </div>
        </div>
    </transition>
</template>

<script>
const TYPE_CLASSES_MAP = {
  'success': 'jus-icon-success',
  'warning': 'jus-icon-warning',
  'danger': 'jus-icon-danger',
  'info': 'jus-icon-info'
}

export default {
  name: 'JusAlert',
  props: {
    title: {
      type: String,
      default: ''
    },
    description: {
      type: String,
      default: ''
    },
    closable: {
      type: Boolean,
      default: false
    },
    closeText: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: 'info'
    },
    showIcon: Boolean
  },
  data () {
    return {
      visible: true
    }
  },
  computed: {
    isBigIcon () {
      return this.description || this.$slots.default ? 'is-big' : ''
    },
    iconClass () {
      return TYPE_CLASSES_MAP[this.type] || 'jus-icon-info'
    },
    alertClass () {
      return 'jus-alert-' + this.type
    }
  },
  methods: {
    close () {
      this.visible = false
      this.$emit('close')
    }
  }
}
</script>

<style lang="scss" scoped>
    $--color-success: #35B34A !default;
    $--color-warning: #8a6d3b !default;
    $--color-danger: #f56c6c !default;
    $--color-info: #909399 !default;

    $--border-success: #35B34A !default;
    $--border-warning: #faebcc !default;
    $--border-danger: #f56c6c !default;
    $--border-info: #909399 !default;

    $--background-success: #f4fbf5 !default;
    $--background-warning: #fcf8e3 !default;
    $--background-danger: #FAF2F2 !default;
    $--background-info: #F4F4F5 !default;
    .jus-alert{
        width: 100%;
        color: #373d41;
        border-radius: 0;
        padding: 8px;
        box-shadow: none;
        font-size: 12px;
        position: relative;
        min-height: 28px;
        box-sizing: border-box;
        display: flex;
        align-items: center;
        border: 1px solid transparent;
        font-family: PingFangSC,helvetica neue,hiragino sans gb,arial,microsoft yahei ui,microsoft yahei,simsun,sans-serif;
        &-success{
            border: 1px solid $--border-success;
            color: $--color-success;
            background-color: $--background-success;
        }
        &-warning{
            border: 1px solid  $--border-warning ;
            color: $--color-warning;
            background-color: $--background-warning;
        }
        &-danger{
            border: 1px solid  $--border-danger ;
            color: $--color-danger;
            background-color: $--background-danger;
        }
        &-info{
            border: 1px solid  $--border-info ;
            color: $--color-info;
            background-color: $--background-info;
        }
        &__icon{
            font-size: 16px;
            width: 16px;
            speak: none;
            position: relative;
            display: inline-block;
            font-style: normal;
            font-weight: 400;
            font-variant: normal;
            text-transform: none;
            -webkit-font-smoothing: antialiased;
            border: none !important;
        }
        .is-big{
            font-size: 28px;
            width: 28px;
        }
        &__content{
            padding: 0 8px;
            font-size: 12px;
            font-weight: 400;
        }
        &__title{
            line-height: 16px;
            font-size: 12px;
            word-break: break-word;
        }
        &__description{
            font-size: 11px;
            margin: 5px 0 0;
        }
        &__closebtn{
            font-size: 13px;
            color: #c0c4cc;
            opacity: 1;
            position: absolute;
            top: 8px;
            right: 15px;
            cursor: pointer;
            line-height: 16px;
            font-style: normal;
        }
    }
</style>
