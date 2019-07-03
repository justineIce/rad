<!-- 日志审计 -->
<template>
    <d2-container class="meta-public">
        <JusLayout title="日志审计">
            <div>
                <el-row :gutter="20">
                    <el-form :model="searchForm" :label-width="labelWidth" label-position="left">
                        <el-col :span="4">
                            <!--<el-form-item label="模块类型：">-->
                                <!--<el-select v-model="searchForm.resourceId">-->
                                    <!--<el-option v-for="item in modelCategoryOption" :key="item.id" :value="item.id" :label="item.name"></el-option>-->
                                <!--</el-select>-->
                            <!--</el-form-item>-->
                        </el-col>
                        <el-col :span="4">
                            <el-form-item label="操作人：">
                                <el-select v-model="searchForm.systemManageId">
                                    <el-option v-for="item in userOption" :key="item.id" :value="item.id" :label="item.FullName"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="4">
                            <el-form-item label="结果类型：">
                                <el-select v-model="searchForm.resultStatus">
                                    <el-option v-for="item in resOption" :key="item.value" :value="item.value" :label="item.label"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="6">
                            <el-form-item label="操作时间：" label-width="80px">
                                <div class="block">
                                    <el-date-picker
                                            v-model="searchForm.operaTime"
                                            size="mini"
                                            type="daterange"
                                            @change="getValue"
                                            range-separator="至"
                                            start-placeholder="开始日期"
                                            end-placeholder="结束日期">
                                    </el-date-picker>
                                </div>
                            </el-form-item>
                        </el-col>
                        <el-col :span="4">
                            <el-button @click="searchLog()">查询</el-button>
                            <el-button @click="reSearchLog">重置</el-button>
                        </el-col>
                    </el-form>
                </el-row>
            </div>
            <div>
                <!--数据列表-->
                <JusTable  class="mT20"
                           :options="{'row-key':'id'}"
                           :columns="column"
                           :data="data"
                           :pagination="pagination"
                           :showPagination="true"
                           @query-data="queryData">
                    <template slot="ResultStatus" slot-scope="scope">
                        <span v-if="scope.ResultStatus=='1'" class="el-icon-success success"> 成功</span>
                        <span v-if="scope.ResultStatus=='0'" class="el-icon-error info"> 失败</span>
                    </template>
                </JusTable>
            </div>
        </JusLayout>
    </d2-container>
</template>
<script>
import { GetTopMenuList, GetAuthorityList, GetManageActionLog } from '@/api/platform.manage'
import util from '@/libs/util'
export default {
  name: 'logRecord',
  data () {
    return {
      searchForm: {
        resourceId: '',
        systemManageId: '',
        resultStatus: '',
        operaTime: [],
        time1: '',
        time2: '',
        pageIndex: 1,
        pageSize: 10
      },
      labelWidth: '80px',
      modelCategoryOption: [],
      userOption: [
        { id: '1', name: 'manger' },
        { id: '2', name: 'dev' },
        { id: '3', name: 'admin' },
        { id: '4', name: 'test' }
      ],
      resOption: [
        { value: '0', label: '失败' },
        { value: '1', label: '成功' }
      ],

      column: [
        { title: '主题', key: 'ActionType' },
        { title: '操作时间',
          key: 'CreatedAt',
          formatter: (row, column, value, index) => {
            return this.$formatDate(value)
          } },
        { title: '操作详情', key: 'FuncName' },
        { title: '操作人', key: 'ManageFullName' },
        { title: '操作结果', key: 'ResultStatus' }
      ],
      data: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 0
      }
    }
  },
  methods: {

    // 获取数据
    searchLog () {
      // 需将时间格式化
      // 时间范围
      this.getDataList()
    },
    // 重置
    reSearchLog () {
      this.searchForm = {
        resourceId: '',
        systemManageId: '',
        resultStatus: '',
        operaTime: [],
        time1: '',
        time2: '',
        pageIndex: 1,
        pageSize: 10
      }
      this.getDataList()
    },

    // 操作时间
    getValue (val) {
      if (this.searchForm.operaTime && this.searchForm.operaTime.length > 0) {
        var date1 = this.searchForm.operaTime[0]
        var date2 = this.searchForm.operaTime[1]
        this.searchForm.time1 = util.formatDate(date1, 'yyyy-MM-dd')
        this.searchForm.time2 = util.formatDate(date2, 'yyyy-MM-dd')
      } else {
        this.searchForm.time1 = ''
        this.searchForm.time2 = ''
      }
    },

    queryData (data, currentPage) {
      this.searchForm.pageIndex = currentPage.pageIndex
      this.searchForm.pageSize = currentPage.pageSize
      this.getDataList()
    },

    // 获取数据
    getDataList () {
      GetManageActionLog(this.searchForm).then((res) => {
        if (res.ret === 200) {
          var data = res.data
          this.data = data.Data
          this.pagination = {
            currentPage: data.PageIndex,
            pageSize: data.PageSize,
            total: data.Count
          }
        }
      })
    }
  },
  mounted () {
    this.getDataList()
  }
}
</script>
<style scoped>
    .success{
        color: #67C23A;
    }
    .info {
        color: #909399;
    }
    .el-form .el-input__inner {
        width: 100%;
    }
</style>
