<template>
  <d2-container>
    <el-row :gutter="24">
      <!--统计总览-->
      <el-col :span="24">
        <el-row :gutter="24">
            <el-col :span="6">
              <count-card circle :end="23" backgroundColor="#FC9C32" icon="el-icon-star-off" title="元数据总数" >
                <span>今日新增：<span>0</span></span>
              </count-card>
            </el-col>
            <el-col :span="6">
              <count-card circle :end="45" backgroundColor="#FC9C32" icon="el-icon-star-off" title="元数据审核总数" >
                <span>今日新增：<span>0</span></span>
              </count-card>
            </el-col>
          <el-col :span="6">
            <count-card circle :end="23" backgroundColor="#80AE42" icon="el-icon-star-off" title="标准数据项总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
          <el-col :span="6">
            <count-card circle :end="45" backgroundColor="#80AE42" icon="el-icon-star-off" title="编码标准总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
        </el-row>
        <el-row :gutter="20" class="mT20">
          <el-col :span="12">
            <count-card circle :end="45" backgroundColor="#58BEB8" icon="el-icon-star-off" title="检核规则总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
          <el-col :span="12">
            <count-card circle :end="45" backgroundColor="#58BEB8" icon="el-icon-star-off" title="中心库总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
        </el-row>
        <el-row :gutter="20" class="mT20">
          <el-col :span="12">
            <count-card circle :end="45" backgroundColor="#5686B8" icon="el-icon-star-off" title="检核规则总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
          <el-col :span="12">
            <count-card circle :end="45" backgroundColor="#5686B8" icon="el-icon-star-off" title="中心库总数" >
              <span>今日新增：<span>0</span></span>
            </count-card>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </d2-container>
</template>

<script>
export default {
  components: {

  },
  data () {
    return {

    }
  }
}
</script>

<style lang="scss" scoped>

</style>
