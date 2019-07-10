<!-- 白名单管理 -->
<template>
    <d2-container class="meta-public">
        <JusLayout title="白名单管理">
            <JusTable ref="table"
                      :columns="columns"
                      :data="data"
                      indexRow
                      :row-handle="rowHandle"
                      :pagination="pagination"
                      @row-remove="rowRemove"
                      @query-data="handleQueryData"
                      @rowEdit="rowEdit">
                <template slot="header">
                    <el-button @click="addBlacklist">新增</el-button>
                </template>
            </JusTable>

            <el-dialog title="添加到白名单" :visible.sync="blacklistDialogVisible">
                <el-form :model="form" ref="form" :rules="formRules" :label-width="labelWidth">
                    <el-form-item label="IP：" prop="ip">
                        <el-input v-model="form.ip" placeholder="请输入加入到白名单的IP"></el-input>
                    </el-form-item>
                    <el-form-item label="备注：">
                        <el-input type="textarea" v-model="form.remarks" :rows="3"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submit">保存</el-button>
                        <el-button @click="blacklistDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </JusLayout>
    </d2-container>
</template>
<script>
import { GetIpList, AddIp, DelIp } from '@/api/platform.manage'
export default {
  name: 'whitelist',
  data () {
    return {
      columns: [
        { title: 'IP地址', key: 'Ip', isQuery: true },
        { title: '创建时间',
          key: 'CreatedAt',
          formatter: (row, column, value, index) => {
            return this.$formatDate(value)
          }
        },
        { title: '备注', key: 'Remarks' }
      ],
      data: [],
      rowHandle: {
        remove: {},
        custom: [
          { text: '编辑', emit: 'rowEdit' }
        ]
      },

      blacklistDialogVisible: false,
      labelWidth: '120px',
      form: {
        ip: '',
        remarks: ''
      },
      formRules: {
        ip: [{ required: true, message: '请输入IP', trigger: 'blur' }]
      },
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 0
      },
      searchForm: {
        pageIndex: '',
        pageSize: '',
        ip: ''
      }
    }
  },
  methods: {
    handleQueryData (data, currentPage) {
      console.log(data, currentPage)
      this.searchForm = {
        pageIndex: currentPage.pageIndex,
        pageSize: currentPage.pageSize,
        ip: data.Ip
      }
      this.getDataList()
    },
    // 删除
    rowRemove ({ index, row }) {
      DelIp({ id: row.ID }).then((res) => {
        if (res.ret === 200) {
          this.$message.success('删除成功')
          this.getDataList()
        }
      })
    },
    rowEdit ({ index, row }) {
      this.blacklistDialogVisible = true
      this.form = {
        id: row.ID,
        ip: row.Ip,
        remarks: row.Remarks
      }
    },

    addBlacklist () {
      this.blacklistDialogVisible = true
      this.form = {
        ip: '',
        remarks: ''
      }
    },

    submit () {
      this.$refs.form.validate((valid) => {
        if (valid) {
          AddIp(this.form).then((res) => {
            if (res.ret === 200) {
              this.$message.success('保存成功')
              this.blacklistDialogVisible = false
              this.getDataList()
            }
          })
        }
      })
    },

    // 获取数据
    getDataList () {
      GetIpList(this.searchForm).then((res) => {
        if (res.ret === 200) {
          var data = res.data
          this.data = data.data
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
<style scoped></style>
