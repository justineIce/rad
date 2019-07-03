<template>
    <el-pagination
            :current-page="pagination.pageIndex"
            :page-size="pagination.pageSize"
            :total="pagination.total"
            layout="total, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange">
    </el-pagination>
</template>

<script>
export default {
  name: 'demo-page-footer',
  props: {
    pagination: {
      pageIndex: 1,
      pageSize: 10,
      total: 0
    }
  },
  methods: {
    handleSizeChange (val) {
      this.$emit('change', {
        pageIndex: this.pagination.pageIndex,
        pageSize: val,
        total: this.pagination.total
      })
    },
    handleCurrentChange (val) {
      this.$emit('change', {
        pageIndex: val,
        pageSize: this.pagination.pageSize,
        total: this.pagination.total
      })
    }
  }
}
</script>
