export default {
  methods: {
    /**
     * @description 组件属性默认值
     */
    handleAttribute (attribute, defaultValue) {
      if (attribute === false || attribute === 0 || attribute === '') {
        return attribute
      }
      return attribute || defaultValue
    },
    /**
     * @description 控制操作列 show 的方法
     */
    handleRowHandleButtonShow (show = true, index, row) {
      if (typeof show === 'boolean') {
        return show
      } else if (typeof show === 'function') {
        return show(index, row)
      }
      return Boolean(show)
    },
    /**
     * @description 控制操作列 disabled 的方法
     */
    handleRowHandleButtonDisabled (disabled = false, index, row) {
      if (typeof disabled === 'boolean') {
        return disabled
      } else if (typeof disabled === 'function') {
        return disabled(index, row)
      }
      return Boolean(disabled)
    },
    /**
     * @description 字段间的关联关系
    */
    handleRowHandleRelation (relation, index, row) {
      if (relation !== undefined) {
        if (relation.field && relation.fieldValue) {
          if (row[relation.field] !== relation.fieldValue) {
            return false
          }
        }
      }
      return true
    },
    /**
 * @description row处于编辑状态
 */
    handleRowHandleEditStatus (editStatus = false, index, row) {
      if (typeof editStatus === 'boolean') {
        return editStatus
      } else if (typeof editStatus === 'function') {
        return editStatus(index, row)
      }
      return Boolean(editStatus)
    }
  }
}
