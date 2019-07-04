<template>
<!--卡片-->
    <div class="jus-card" >
        <div class="jus-card__header" v-if="$slots.header || header">
            <slot name="header">{{ header }}</slot>
        </div>
        <div class="jus-card__body" :style="bodyStyle">
            <slot></slot>
        </div>
    </div>
</template>

<script>
export default {
  name: 'JusCard',
  props: {
    header: {},
    bodyStyle: {}
  }
}
</script>

<style lang="scss" scoped>
    .jus-card{
        box-sizing: border-box;
        margin: 0 0 16px 0;
        //box-shadow: 0 1px 1px rgba(0,0,0,0.15);
        box-shadow: 0 0 3px rgba(0,0,0,0.15);
        position: relative;
        &__header{
            box-sizing:border-box;
            height: 45px;
            line-height: 45px;
            padding-left: 16px;
            //background: #F5F5F6;
            border-bottom: 1px solid #F2F2F2;
            font-size: 14px;
            color: #666666;
        }
        &__body{
            background: #fff;
            padding: 16px;
            color: #8b8c8c;
            position:relative;
        }
    }
</style>
