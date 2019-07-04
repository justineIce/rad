<template>
    <!--提醒-->
    <div class="jus-remind">
        <el-tooltip popper-class="b-popover"  effect="dark" :content="text" placement="bottom-start">
           <slot>
               <i class="el-icon-question" style="font-size: 16px;" :style="{color:color}"></i>
           </slot>
        </el-tooltip>
        <!--<el-popover popper-class="b-popover"-->
                <!--placement="top-start"-->
                <!--width="200"-->
                <!--trigger="hover"-->
                <!--:content="text">-->
            <!--<i class="el-icon-question" style="font-size: 16px;" slot="reference"></i>-->
        <!--</el-popover>-->
    </div>
</template>

<script>
export default {
  name: 'JusRemind',
  props: {
    text: String,
    color: {
      type: String,
      default: '#333333'
    }
  }
}
</script>

<style lang="scss">
.b-popover{
    width: 200px;
}
.jus-remind{
    position: relative;
    /*width: 100%;*/
    &>span{
        display: inline-block;
    }
}
</style>
