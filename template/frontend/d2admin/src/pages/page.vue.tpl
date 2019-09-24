`<template>
    <d2-container class="meta-public">
        <JusLayout title="{{.modelInfo.TableRemark}}">
            <JusTable :columns="column"
                      :data="data.data"
                      :pagination="{total:data.totalCount}"
                      {{if .modelInfo.TableHandle.select}}@query-data="handleQueryData"{{end}}
                     {{if or .modelInfo.TableHandle.insert .modelInfo.TableHandle.update}}addMode{{end}}
                      {{if .modelInfo.TableHandle.insert}}addTitle="添加{{.modelInfo.TableRemark}}"
                      :addTemplate="addTemplate"
                      :addRules="rules"{{end}}
                      {{if .modelInfo.TableHandle.update}}editTitle="编辑{{.modelInfo.TableRemark}}"
                      :editTemplate="addTemplate"
                      :editRules="rules"{{end}}
                      {{if or .modelInfo.TableHandle.update .modelInfo.TableHandle.delete }}:rowHandle="rowHandle"{{end}}
                      {{if .modelInfo.TableHandle.insert}}@row-add="handleRowSave"{{end}}
                      {{if .modelInfo.TableHandle.delete}}@row-remove="handleRowRemove"{{end}}
                      {{if .modelInfo.TableHandle.update}}@row-edit="handleRowSave"{{end}}
                      />
        </JusLayout>

    </d2-container>
</template>
<script>
import util from '@/libs/util'
import { Get{{.modelInfo.StructName}}Page{{if or .modelInfo.TableHandle.insert .modelInfo.TableHandle.update}}, Save{{.modelInfo.StructName}}{{end}}{{if .modelInfo.TableHandle.delete}}, Del{{.modelInfo.StructName}}{{end}} } from '@/api/{{.modelInfo.SingName}}'
export default {
  name: '{{.modelInfo.SingName}}Index',
  data () {
    return {
      column: [
      {{- if not (eq .modelInfo.TableView "") -}}
        {{range .ViewInfo.Columns}}
        {{- if not .PageTable.none -}}
        {{- if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at")) -}}
        { title: '{{if eq .ColumnCNName ""}}{{.StructName}}{{else}}{{.ColumnCNName}}{{end}}', key: '{{.StructName}}'{{if .PageTable.search}}, isQuery: {{.PageTable.search}}{{end}} {{if eq .DataTypeLower "datetime"}} , formatter: this.formatter {{end}} },
        {{end}}
        {{- end -}}
        {{- end -}}
      {{else}}
        {{range .modelInfo.Columns}}
        {{- if not .PageTable.none -}}
        {{- if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at")) -}}
        { title: '{{if eq .ColumnCNName ""}}{{.StructName}}{{else}}{{.ColumnCNName}}{{end}}', key: '{{.StructName}}'{{if .PageTable.search}}, isQuery: {{.PageTable.search}}{{end}} {{if eq .DataTypeLower "datetime"}} , formatter: this.formatter {{end}} },
        {{end}}
        {{- end -}}
        {{- end -}}
      {{- end -}}
      ],
      data: [],
      {{if or .modelInfo.TableHandle.insert .modelInfo.TableHandle.update}}
      addTemplate: {
        {{range .modelInfo.Columns}}
        {{- if not .PageForm.none -}}
        {{- if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at")) -}}
        {{- if eq .DataTypeLower "date" -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-date-picker',type:'date' }},
        {{- else if eq .DataTypeLower "datetime" -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-date-picker',type:'datetime' }},
        {{- else if eq .DataTypeLower "time" -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-time-select' }},
        {{- else if and (gt .CharacterMaximumLength 300) (or (eq .DataTypeLower "char") (eq .DataTypeLower "varchar") (eq .DataTypeLower "longtext") (eq .DataTypeLower "mediumtext") (eq .DataTypeLower "text") (eq .DataTypeLower "tinytext")) -}}
        {{.StructName}}: { title: '{{.ColumnCNName}} {{(gt .CharacterMaximumLength 300)}} {{(.CharacterMaximumLength)}}', component: { name: 'el-input', type: 'textarea', rows: 5 }},
        {{- else if eq .DataTypeLower "enum" -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}', component: { name: 'el-radio',options:[{{ range .ColumnItemValue }}{value:'{{.}}',label:'{{.}}'},{{end}}] }},
        {{- else if or (eq .DataTypeLower "decimal") (eq .DataTypeLower "double") (eq .DataTypeLower "float") -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}', component: {name: 'el-input', type: 'number', min: 0 }},
        {{- else -}}
        {{.StructName}}: { title: '{{.ColumnCNName}}' },{{end}}
        {{end}}
        {{- end -}}{{- end -}}
      },{{end}}
      rowHandle: {
        edit: {},
        remove: {}
      },
      rules: {
      {{- range .modelInfo.Columns -}}
      {{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
         {{.StructName}}: [
           {{if not (eq .IsNullable "YES")}}
           { required: true,message: '请输入{{.ColumnCNName}}', trigger: 'blur' },
           {{end}}
           {{- if or (eq .DataTypeLower "date") (eq .DataTypeLower "datetime") (eq .DataTypeLower "time") (eq .DataTypeLower "timestamp") -}}
           { type: 'date',message: '请正确选择的时间', trigger: ['blur', 'change'] }
           {{- else if not (eq .CharacterMaximumLength 0) -}}
           { max: {{.CharacterMaximumLength}}, message: '长度不能超过{{.CharacterMaximumLength}}个字符', trigger: 'blur' },
           {{end}}
           {{- range .Valid -}}
           {{- if eq . "Email" -}}
           { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },
           {{- end -}}
           {{- end -}}
         ],
         {{- end -}}
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
      // 加载loading
      util.loading = this.$loading({
        lock: true,
        text: '保存中...',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      Get{{.modelInfo.StructName}}Page(data).then((result) => {
           if (result.ret === 200) {
                this.data = result.data
           }
      })
    },
    /**
       * 条件查询及翻页
       * @param data
       * @param pagination
       */
    handleQueryData (data, pagination) {
      // 加载loading
      util.loading = this.$loading({
        lock: true,
        text: '保存中...',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.getListData(this.$objectAssign(data, pagination))
    },
    handleRowRemove (data, done) {
      // 加载loading
      util.loading = this.$loading({
        lock: true,
        text: '保存中...',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      Del{{.modelInfo.StructName}}({ id: data.row.ID }).then((result) => {
        if (result.ret === 200) {
            this.$message({ message: '删除成功', type: 'success' })
            done(false)
            this.getListData()
        }
      })
    },
    handleRowSave (data, done) {
      // 加载loading
      util.loading = this.$loading({
        lock: true,
        text: '保存中...',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var isAdd=true
      if (data.row){
        data=data.row
        isAdd=false
      }
      Save{{.modelInfo.StructName}}({
        ID: data.ID,{{ range .modelInfo.Columns }}{{if and (not (eq .ColumnName "id")) (not (eq .ColumnName "created_at")) (not (eq .ColumnName "created_by")) (not (eq .ColumnName "updated_by")) (not (eq .ColumnName "updated_at")) (not (eq .ColumnName "deleted_at"))}}
        {{.StructName}}: data.{{.StructName}},{{end}}{{end}}
      }).then((result) => {
        if (result.ret === 200) {
            this.$message({ message: isAdd?'添加成功':'修改成功', type: 'success' })
            done(false)
            this.getListData()
        }
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
