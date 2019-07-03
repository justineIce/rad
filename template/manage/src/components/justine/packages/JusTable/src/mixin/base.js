export default {
  props: {
    /**
    * @description 表格配置
    */
    options: {
      type: Object,
      default: null
    },
    /**
    * @description 表头数据
    */
    columns: {
      type: Array,
      request: true
    },
    /**
     * @description 表格数据
     */
    data: {
      type: Array,
      request: true
    },
    /**
     * @description 默认排序
     */
    rowHandle: {
      type: Object,
      default: null
    },
    /**
       * @description 多选
       */
    selectionRow: {
      default: null
    },
    pagination: {
      type: Object,
      default: null
    },
    /**
     * @description 查询表单数据
     */
    queryFormData: {
      type: Object,
      default: null
    },
    /**
       * @description 查询表单是否显示
       */
    queryForm: {
      type: Boolean,
      default: true
    },
    /**
     * @description 是否显示翻页控件
     */
    showPagination: {
      type: Boolean,
      default: true
    },
    /**
     * @description 扩展
     */
    expand: {
      type: Boolean,
      default: false
    },
    /**
     * @description 表格处于编辑状态
     */
    isEdit: {
      type: Boolean,
      default: false
    },
    /**
     * @indexRow 排序
     */
    indexRow: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      queryData: {},
      selectionData: []
    }
  },
  methods: {
    /**
       * @description 勾选数据时触发的事件
       */
    handleSelect (selection, row) {
      this.$emit('select', selection, row)
    },
    /**
       * @description 勾选全选时触发的事件
       */
    handleSelectAll (selection) {
      this.$emit('select-all', selection)
    },
    /**
       * @description 复选框选择项发生变化时触发的事件
       */
    handleSelectionChange (selection) {
      this.selectionData = selection
      this.$emit('selection-change', selection)
    },
    /**
       * @description 单元格 hover 进入时触发的事件
       */
    handleCellMouseEnter (row, column, cell, event) {
      this.$emit('cell-mouse-enter', row, column, cell, event)
    },
    /**
       * @description 单元格 hover 退出时触发的事件
       */
    handleCellMouseLeave (row, column, cell, event) {
      this.$emit('cell-mouse-leave', row, column, cell, event)
    },
    /**
       * @description 单元格点击时触发的事件
       */
    handleCellClick (row, column, cell, event) {
      this.$emit('cell-click', row, column, cell, event)
    },
    /**
       * @description 单元格双击时触发的事件
       */
    handleCellDblclick (row, column, cell, event) {
      this.$emit('cell-dblclick', row, column, cell, event)
    },
    /**
       * @description 行点击时触发的事件
       */
    handleRowClick (row, event, column) {
      this.$emit('row-click', row, event, column)
    },
    /**
       * @description 行右键点击时触发的事件
       */
    handleRowContextmenu (row, event) {
      this.$emit('row-contextmenu', row, event)
    },
    /**
       * @description 行双击时触发的事件
       */
    handleRowDblclick (row, event) {
      this.$emit('row-dblclick', row, event)
    },
    /**
       * @description 表头点击时触发的事件
       */
    handleHeaderClick (column, event) {
      this.$emit('header-click', column, event)
    },
    /**
       * @description 表头右键点击时触发的事件
       */
    handleHeaderContextmenu (column, event) {
      this.$emit('header-contextmenu', column, event)
    },
    /**
       * @description 行选中状态
       */
    handleCurrentChange (currentRow, oldCurrentRow) {
      this.$emit('current-change', currentRow, oldCurrentRow)
    },
    handleSortChange (sort) {
      this.$emit('sort-change', sort)
    },
    handleFilterChange (filters) {
      this.$emit('filter-change', filters)
    },
    /**
    * @description 获取选中的数据
    */
    getSelectsData () {
      return this.selectionData
    },
    /**
     * @description 翻页发生变化时触发的事件
     */
    handleCurrentPage (page) {
      this.$emit('query-data', this.queryData, {
        pageIndex: page,
        pageSize: this.pagination && this.pagination.pageSize ? this.pagination.pageSize : 10
      })
    },
    /**
    * @description 查询按钮触发的事件
    */
    handleQueryMethod (data) {
      this.queryData = data
      this.$emit('query-data', this.queryData, {
        pageIndex: 1,
        pageSize: this.pagination && this.pagination.pageSize ? this.pagination.pageSize : 10
      })
    },
    /**
     * @description 新增行数据
     * @param {Object} row 新增的表格行数据
     */
    handleAddRow (row) {
      this.$set(this.data, this.data.length, row)
    },
    handleRemoveDone (index) {
      this.$delete(this.data, index)
    },
    /**
       * @description 更新行数据
       * @param {Number} index 表格数据索引
       * @param {Object} row 更新的表格行数据
       */
    handleUpdateRow (index, row) {
      this.$set(this.data, index, row)
    },
    /**
    * @description 删除行
    */
    handleRemove (index, row) {
      if ((!('confirm' in this.rowHandle.remove)) || this.rowHandle.remove.confirm === true) {
        this.$confirm(this.handleAttribute(this.rowHandle.remove.confirmText, '确定删除吗？'), this.handleAttribute(this.rowHandle.remove.confirmTitle, '删除'), {
          confirmButtonText: this.handleAttribute(this.rowHandle.remove.confirmButtonText, '确定'),
          cancelButtonText: this.handleAttribute(this.rowHandle.remove.cancelButtonText, '取消'),
          type: this.handleAttribute(this.rowHandle.remove.type, 'error')
        }).then(() => {
          this.$emit('row-remove', { index, row: row }, () => {
            this.handleRemoveDone(index)
          })
        }).catch(() => {})
      } else {
        this.$emit('row-remove', { index, row: row }, () => {
          this.handleRemoveDone(index)
        })
      }
    }
  }
}
