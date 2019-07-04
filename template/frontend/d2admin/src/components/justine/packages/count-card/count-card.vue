<template>
    <div class="count-card">
        <div class="count-card__wrapper">
            <div class="count-card__left" :class="{'circle':circle}" >
                <div class="count-card-icon"  :style="{background: backgroundColor}">
                    <i class="iconfont" :class="[icon]"></i>
                </div>
            </div>
            <div class="count-card__right">
                <div class="count-card__number">
                    <count-to :end="end" :startVal="start" usegroup></count-to>
                </div>
                <p>{{title}}</p>
            </div>
        </div>
        <div class="count-card__foot" v-if="$slots.default">
            <slot></slot>
        </div>
    </div>
</template>

<script>
export default {
  name: 'CountCard',
  props: {
    title: String,
    start: {
      type: Number,
      default: 0
    },
    end: Number,
    icon: String,
    backgroundColor: {
      type: String,
      default: '#5BC3EB'
    },
    circle: Boolean
  }
}
</script>
<style lang="scss" scoped>
    .count-card{
        width: 100%;
        height: 100%;
        box-sizing: border-box;
        box-shadow: 0 0 5px rgba(0,0,0,0.15);
        border-radius: 4px;
        position: relative;
        .count-card__wrapper{
            display: flex;
            .count-card__left{
                /*background-color: #0086b3;*/
                border-top-left-radius: 4px;
                border-bottom-left-radius: 4px;
                width: 60px;
                .count-card-icon{
                    width: 100%;
                    height: 100%;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    color: #fff;
                    /*padding: 0 10px;*/
                    &>i{
                        font-size: 30px;
                    }
                }
                &.circle{
                    width: 80px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    .count-card-icon{
                        border-radius: 50%;
                        width: 50px;
                        height: 50px;
                        color: #fff;
                    }
                }
            }
            .count-card__right{
                width: 65%;
                padding: 10px;
                display: flex;
                flex-direction: column;
                text-align: center;
                .count-card__number{
                    font-size: 30px;
                    flex: 1;
                }
                &>p{
                    margin: 5px 0;
                    color: #B7B7B7;
                }
            }
        }
        .count-card__foot{
            border-top: 1px solid rgba(0,0,0,0.1);
            padding: 10px;
            text-align: left;
            color: #B7B7B7;
        }
    }
</style>
