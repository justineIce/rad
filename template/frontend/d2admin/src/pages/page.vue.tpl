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
import { Get{{.modelInfo.StructName}}Page, Save{{.modelInfo.StructName}}, Del{{.modelInfo.StructName}} } from '@/api/{{.modelInfo.TableName}}'
export default {
  name: '{{.modelInfo.SingName}}',
  data () {
    return {
      column: [ {{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
        { title: '{{.ColumnCNName}}', key: '{{.StructName}}' {{if eq .DataTypeLower "datetime"}} , formatter: this.formatter {{end}} }, {{end}}{{end}}
      ],
      data: [],
      addTemplate: {    {{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
        {{if eq .DataTypeLower "date"}}{{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-date-picker' }},
        {{else if eq .DataTypeLower "datetime"}}{{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-time-picker' }},
        {{else if eq .DataTypeLower "time"}}{{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-time-select' }},
        {{else if and (gt .CharacterMaximumLength "300") (or (eq .DataTypeLower "char") (eq .DataTypeLower "varchar") (eq .DataTypeLower "longtext") (eq .DataTypeLower "mediumtext") (eq .DataTypeLower "text") (eq .DataTypeLower "tinytext"))}}
            {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-input', type: 'textarea', rows: 5 }},
        {{else if eq .DataTypeLower "enum"}}
          {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-radio',options:[{{ range .ColumnItemValue }}'{{.}}',{{end}}] }},
        {{else if or (eq .DataTypeLower "decimal") (eq .DataTypeLower "double") (eq .DataTypeLower "float") }}
          {{.StructName}}: { title: '{{.ColumnCNName}}', component: {name: 'el-input', type: 'number', min: 0 }},
        {{else}}
          {{.StructName}}: { title: '{{.ColumnCNName}}' },
        {{end}}
        {{end}}{{end}}
      },
      rowHandle: {
        edit: {},
        remove: {}
      },
      rules: {  {{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
         {{.StructName}}: [
           { {{if or (eq .DataTypeLower "date") (eq .DataTypeLower "datetime") (eq .DataTypeLower "time") (eq .DataTypeLower "timestamp")}}type: 'date',{{end}}{{if not (eq .IsNullable "YES")}}required: true,{{end}}message: '请输入{{.ColumnCNName}}', trigger: 'blur' },
           {{ if not (eq .CharacterMaximumLength "")}}{ max: {{.CharacterMaximumLength}}, message: '长度不能超过{{.CharacterMaximumLength}}个字符', trigger: 'blur' },{{end}}
           {{ range .Valid }}{{if eq . "Email"}}{ type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },{{end}}
           {{end}}
         ],{{end}}{{end}}
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
           if (result.ret === 200) {
                this.data = result.data.data
           }
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
      Save{{.modelInfo.StructName}}({
        ID: data.ID,{{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
        {{.StructName}}: data.row.{{.StructName}},{{end}}{{end}}
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
      Save{{.modelInfo.StructName}}({
        ID: data.row.ID,{{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
        {{.StructName}}: data.row.{{.StructName}},{{end}}{{end}}
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
