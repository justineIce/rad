import _forEach from 'lodash.foreach'
import _set from 'lodash.set'
export default {
  props: {
    formOptions: Object,
    /**
     * @description add edit 添加修改
     */
    addMode: {
      type: Boolean,
      default: false
    },
    addTitle: {
      type: String
    },
    addTemplate: {
      type: Object,
      default: null
    },
    editTitle: {
      type: String,
      default: '编辑'
    },
    editTemplate: {
      type: Object,
      default: null
    },
    addRules: {
      type: Object,
      default: null
    },
    editRules: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      /**
       * @description dialog显示与隐藏
       */
      dialogVisible: false,
      formMode: 'edit',
      editIndex: 0,
      formData: {},
      formTemplate: {},
      formRules: {}
    }
  },
  methods: {
    /**
     * @description 关闭模态框
     */
    handleCloseDialog () {
      this.dialogVisible = false
    },
    /**
       * @description 编辑行数据
       * @param {Number} index 行所在索引
       * @param {Object} row 行数据
       */
    handleEdit (index, row, templage = null) {
      this.formMode = 'edit'
      this.dialogVisible = true
      this.editIndex = index
      if (!this.editTemplate) throw new Error('[table] editTemplate is required when edit')
      this.formTemplate = this.editTemplate
      this.formRules = this.editRules
      var obj = {}
      for (let prop in row) {
        obj[prop] = row[prop]
      }
      this.formData = obj
    },
    handleAdd () {
      this.formMode = 'add'
      if (!this.addTemplate) throw new Error('[table] addTemplate is required when edit')
      this.dialogVisible = true
      this.formTemplate = this.addTemplate
      this.formRules = this.addRules
      this.formData = {}
    },
    handleDialogSaveDone (rowData) {
      if (this.formMode === 'edit') {
        this.handleUpdateRow(this.editIndex, rowData)
      } else if (this.formMode === 'add') {
        this.handleAddRow(rowData)
      }
      this.handleCloseDialog()
    },
    /**
     * @description 表单提交数据
     * @param data
     */
    formSubmit (data) {
      let rowData = {}
      if (this.formMode === 'edit') {
        _forEach(data, (value, key) => {
          _set(rowData, key, value)
        })
        this.$emit('row-edit', {
          index: this.editIndex,
          row: rowData
        }, (param = null) => {
          if (param === false) {
            this.handleCloseDialog()
            return
          }
          this.handleDialogSaveDone({
            ...rowData,
            ...param
          })
        })
      } else if (this.formMode === 'add') {
        _forEach(data, (value, key) => {
          _set(rowData, key, value)
        })
        this.$emit('row-add', data, (param = null) => {
          if (param === false) {
            this.handleCloseDialog()
            return
          }
          this.handleDialogSaveDone({
            ...data,
            ...param
          })
        })
      }
    }
  }
}
