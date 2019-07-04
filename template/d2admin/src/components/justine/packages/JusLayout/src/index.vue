<template>
    <!--左右结构框架 默认左侧内容是树形菜单-->
    <div class="justine-layout">
        <div class="justine-layout-container">
            <div class="justine-layout-container__header" v-if="title!==''">
                <slot name="header">
                    <span v-if="!$slots.header && title!==''">{{title}}</span>
                    <i class="icon-back jus-icon-fanhui" @click="handleGoback"></i>
                </slot>
            </div>
            <div class="justine-layout-container__body">
                <!--展开按钮-->
                <div class="justine-layout-container-collapse" v-if="showCollapse && aside && position === 'left'"  @click="collapse = !collapse" :style="{left: (collapse ? (asideWidth - 16) : 0) + 'px' ,transform: 'rotate(' + (collapse ? 0 : 180 ) + 'deg)' }">
                    <span class="collapse-icon"><i class="el-icon-d-arrow-left"></i></span>
                </div>
                <!--左边内容区域-->
                <div flex-box="0"  class="justine-layout-container-aside" v-if="aside"
                     :style="{
                 'background-color':backgroundColor,
                 width: asideWidth + 'px',
                 display: collapse ? 'block' : 'none',
                 left: position === 'left' ? '0' : 'unset' ,
                 right: position === 'right' ? '0' : 'unset' ,
                 'border-right': position === 'left' ? '1px solid #dddddd': 'unset' ,
                'border-left': position === 'right' ? '1px solid #dddddd': 'unset' } ">
                    <!--标题-->
                    <div class="justine-layout-container-aside-header" v-if="asideHeader">
                        <div>
                            <i class="fa el-icon-menu icon"></i>
                            <span class="aside-header-title">{{asideTitle}}</span>
                        </div>
                        <div style="display: flex;align-items: center; justify-content: space-between;">
                            <!--<i class="el-icon-plus"></i>-->
                            <slot name="aside-header"/>
                        </div>
                    </div>
                    <slot name="aside"/>
                </div>
                <!--右边内容区域-->
                <div class="justine-layout-container-main" flex flex-box="1"
                     :style="{
                 left: !aside ? 0 : (position === 'left'? (((collapse ? asideWidth : 0))  + 'px') : '0'),
                 right: !aside ? 0 : (position === 'right' ? ((collapse ? asideWidth : 0))  + 'px' : '0' ) }">
                    <!--标题-->
                    <!--<div class="justine-layout-container-main-header" v-if="!$slots.header && title!==''">-->
                    <!--<span >{{title}}</span>-->
                    <!--</div>-->
                    <!--<slot name="header"></slot>-->
                    <!--内容-->
                    <div class="justine-layout-container-main-body">
                        <slot></slot>
                    </div>
                    <!--footer-->
                    <div class="justine-layout-container-main-footer">
                        <slot name="footer" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'JusLayout',
  props: {
    /**
     *@description 标题
     */
    title: {
      type: String,
      default: ''
    },
    /**
    /**
     * @description 左边侧边栏
     */
    aside: Boolean,
    asideTitle: String,
    asideHeader: Boolean,
    asideWidth: {
      type: Number,
      default: 200
    },
    position: {
      type: String,
      default: 'left'
    },
    backgroundColor: {
      type: String,
      default: '#EAEDF1;'
    },
    showCollapse: {
      type: Boolean,
      default: true
    }
  },
  data () {
    return {
      collapse: true
      // [侧边栏宽度]
    }
  },
  methods: {
    handleGoback () {
      window.history.length > 1
        ? this.$router.go(-1)
        : this.$router.push('/')
    }
  },
  mounted () {

  }
}
</script>

<style lang="scss" scoped>
.justine-layout{
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    right: 0;
    overflow: auto;
    &-container{
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        &__header{
            width: 100%;
            box-sizing: border-box;
            height: 50px;
            line-height: 50px;
            border-bottom: 1px solid #dedede;
            padding: 0 16px;
            color: #373d41;
            background: #fff;
            font-size: 16px;
            .icon-back{
                color: $theme-color;
                margin:  0 10px;
                font-weight: 700;
                cursor: pointer;
            }
            & > span {
                font-size: 16px;
                font-weight: 500;
                position: relative;
                border: none;
                padding-left: 11px;
                &:before {
                    content: "";
                    position: absolute;
                    top: 50%;
                    left: 0;
                    margin-top: -9px;
                    width: 3px;
                    height: 18px;
                    background: -webkit-gradient(linear, left bottom, left top, from($theme-linear-gradient-color), to($theme-color));
                    background: -webkit-linear-gradient(bottom, $theme-linear-gradient-color, $theme-color);
                    background: linear-gradient(0deg, $theme-linear-gradient-color, $theme-color);
                    -webkit-border-radius: 4px;
                    -moz-border-radius: 4px;
                    border-radius: 4px;
                }
            }
        }
        &__body{
            display: flex;
            flex:1;
            position: relative;
            padding: 0 20px;
        }
        &::-webkit-scrollbar {
            display: none;
        }
        &-collapse{
            width: 15px;
            height: 50px;
            left: 254px;
            position: absolute;
            z-index: 234;
            top: 50%;
            background-color: #F7F7F7;
            cursor: pointer;
            text-align: center;
            -webkit-transition: all 0.1s ease, 0.1s ease;
            clip-path: polygon(100% 0, 100% 100%, 0 80%, 0 20%);
            -webkit-clip-path: polygon(100% 0, 100% 100%, 0 80%, 0 20%);
            .collapse-icon{
                line-height: 50px;
                color: #546478;
            }
        }
        &-aside{
            background: #eaedf1;
            width: 270px;
            border-right: 1px solid #dddddd;
            position: absolute;
            top: 0;
            left: 0;
            bottom: 0;
            overflow: auto;
            box-sizing: border-box;
            /*padding: 10px;*/
            &-header{
                background-color: #f8f8f8;
                height: 35px;
                line-height: 35px;
                padding: 0 16px;
                border-bottom: 1px solid #ebebeb;
                display: flex;
                justify-content: space-between;
                align-items: center;
                color: $theme-color;
                font-size: 15px;
                .icon{
                    height: 35px;
                    line-height: 35px;
                    margin: 0;
                    color: $theme-color;
                    font-size: 15px;
                    -webkit-user-select: none;
                    -moz-user-select: none;
                    -ms-user-select: none;
                    user-select: none;
                }
                .aside-header-title{
                    height: 35px;
                    margin: 0 10px;
                    font-size: 13px;
                    display: inline-block;
                    vertical-align: top;
                    color: $theme-color;
                }
            }
        }
        &-main{
            display: flex;
            flex-direction: column;
            position: absolute;
            top: 0;
            right: 0;
            bottom: 0;
            left: 270px;
            overflow: auto;
            &::-webkit-scrollbar {
                display: none;
            }
            &-header {
                height: 50px;
                line-height: 50px;
                /*border-bottom: 1px solid #dedede;*/
                padding: 0 16px;
                color: #373d41;
                background: #fff;
                font-size: 16px;
                box-sizing: border-box;
                & > span {
                    font-size: 16px;
                    font-weight: 500;
                    position: relative;
                    border: none;
                    padding-left: 11px;
                    &:before {
                        content: "";
                        position: absolute;
                        top: 50%;
                        left: 0;
                        margin-top: -9px;
                        width: 3px;
                        height: 18px;
                        background: -webkit-gradient(linear, left bottom, left top, from($theme-linear-gradient-color), to($theme-color));
                        background: -webkit-linear-gradient(bottom, $theme-linear-gradient-color, $theme-color);
                        background: linear-gradient(0deg, $theme-linear-gradient-color, $theme-color);
                        -webkit-border-radius: 4px;
                        -moz-border-radius: 4px;
                        border-radius: 4px;
                    }
                }
            }
            &-body{
                /*-webkit-box-flex: 1;*/
                /*-ms-flex-positive: 1;*/
                /*flex-grow: 1;*/
                padding: 16px;
                flex: 1;
                overflow: auto;
                position: relative;
                background:#ffffff;
                box-sizing: border-box;
                &::-webkit-scrollbar {
                    display: none;
                }
            }
            &-footer{
                /*padding: 20px;*/
                border-top: 1px solid #cfd7e5;
                background: #FFF;
            }
        }
    }
}
</style>
