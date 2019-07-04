<template>
    <div class="justine-block">
        <div class="justine-block-container" :class="selected? 'block-active': ''" :style="{'flex-direction': (position==='right' ? 'row-reverse' : (position==='top' ? 'column' : 'row'))}">
            <!--图片-->
            <div v-if="icon" class="justine-block-container-aside">
                <div class="aside-logo">
                    <img v-if="img" class="icon"  :src="icon">
                    <i v-if="!img" :class="[icon]" aria-hidden="true"></i>
                </div>
            </div>
            <!--内容-->
            <div class="justine-block-container-body" :style="{'padding-left': position==='right' ? '0' :'15px','padding-right': position==='right' ? '10px' :'0' }">
                <!--标题-->
                <div class="justine-block-container-body-title">
                    <p class="title">{{title}}</p>
                    <div class="desc">
                        <slot v-bind="data">
                            {{desc}}
                        </slot>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'JustineBlock',
  props: {
    /**
     * @description 是否img图片
     */
    img: {
      type: Boolean,
      default: false
    },
    /**
     * @description 图片地址或是icon
     */
    icon: {
      type: String
    },
    /**
     * @description 图片位置 left right top
     */
    position: {
      type: String,
      default: 'left'
    },
    /**
     * @description 标题
     */
    title: String,
    /**
     * @description 数据
     */
    data: [Object, Number, String],
    /**
     * @description 描述
     */
    desc: String,
    /**
     *@description 是否选中
     */
    selected: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style lang="scss" scoped>
$hover-color: #00c4e2;
$theme-color: #e5e5e5;
.justine-block{
    .block-active{
        border: 1px solid $hover-color;
        box-shadow: 0 0 6px $hover-color;
    }
    &-container{
        padding: 20px 15px;
        display: flex;
        border: 1px solid $theme-color;
        box-shadow: 0 0 6px $theme-color;
        &:hover{
            border: 1px solid $hover-color;
            box-shadow: 0 0 6px $hover-color;
            transform: scale(1.1,1.1);
        }
        &-aside{
            max-width: 20%;
            .aside-logo{
                overflow: hidden;
                width: 100%;
                text-align: center;
                .icon{
                    max-width: 100%;
                    display: block;
                    font-size: 32px;
                    margin-bottom: 15px;
                    color: #606266;
                    transition: color .15s linear;
                }
            }
        }
        &-body{
            padding-left: 15px;
            flex:1;
            &-title{
                .title{
                    overflow: hidden;
                    text-overflow: ellipsis;
                    -webkit-line-clamp: 2;
                    display: -webkit-box;
                    -webkit-box-orient: vertical;
                    margin: 0 0 10px 0;
                    font-size: 15px;
                    font-weight: 700;
                }
                .desc{
                    font-size: 12px;
                    color: #999;
                    line-height: 20px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-line-clamp: 3;
                    -webkit-box-orient: vertical;
                }
            }
        }
    }
}
</style>
