`<template>
    <d2-container class="meta-public">
        <JusLayout title="{{.modelInfo.TableRemark}}">
            <JusTable :columns="column"
                      :data="data"
                      @query-data="handleQueryData"
                      addMode
                      addTitle="添加{{.modelInfo.TableRemark}}"
                      :addTemplate="addTemplate"
                      :addRules="rules"
                      editTitle="编辑{{.modelInfo.TableRemark}}"
                      :editTemplate="addTemplate"
                      :editRules="rules"
                      :rowHandle="rowHandle"
                      @row-add="handleRowAdd"
                      @row-remove="handleRowRemove"
                      @row-edit="handleRowEdit"/>
        </JusLayout>
    </d2-container>
</template>
<script>

import { Get{{.modelInfo.StructName}}Page, Add{{.modelInfo.StructName}}, Del{{.modelInfo.StructName}} } from '@/api/{{.modelInfo.TableName}}'
export default {
  name: '{{.modelInfo.SingName}}',
  data () {
    return {
      column: [
        {{ range .modelInfo.Columns }}
        { title: '{{.ColumnCNName}}', key: '{{.StructName}}' {{if eq .DataTypeLower "datetime"}} , formatter: this.formatter {{end}} }, {{end}}
      ],
      data: [],
      addTemplate: {
        {{ range .modelInfo.Columns }}{{.StructName}}: { title: '类型名称', component: { name: 'el-input', type: 'textarea', rows: 5 }},
        {{end}}
      },
      rowHandle: {
        edit: {},
        remove: {}
      },
      rules: {

        {{ range .modelInfo.Columns }}
         Name: [
                  { required: true, message: '请输入名称', trigger: 'blur' },
                  { min: 1, max: 50, message: '长度不能超过50个字符', trigger: 'blur' }
                ],
        {{end}}

      }
    }
  },
  methods: {
    formatter (row, column, value, index) {
      return this.$formatDate(value)
    },
    /**
       * 获取列表数据
       * @param data
       */
    getListData (data) {
      Get{{.modelInfo.StructName}}Page(data).then((result) => {
        this.data = result.data
      })
    },
    /**
       * 条件查询及翻页
       * @param data
       * @param pagination
       */
    handleQueryData (data, pagination) {
      this.getListData(this.$objectAssign(data, pagination))
    },
    /**
       * 添加数据
       * @param data
       * @param done
       */
    handleRowAdd (data, done) {
      Add{{.modelInfo.StructName}}({
        id: data.ID,
        name: data.Name,
        remark: data.Remark
      }).then((result) => {
        this.$message({ message: '添加成功', type: 'success' })
        done(false)
        this.getListData()
      })
    },
    handleRowRemove (data, done) {
      Del{{.modelInfo.StructName}}({ id: data.row.ID }).then((result) => {
        this.$message({ message: '删除成功', type: 'success' })
        done(false)
        this.getListData()
      })
    },
    handleRowEdit (data, done) {
      Add{{.modelInfo.StructName}}({
        id: data.row.ID,
        name: data.row.Name,
        remark: data.row.Remark
      }).then((result) => {
        this.$message({ message: '修改成功', type: 'success' })
        done(false)
        this.getListData()
      })
    }
  },
  mounted () {
    this.getListData()
  }
}
</script>

<style scoped>

</style>
`
