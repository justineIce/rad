<template>
    <div class="drag-list-wrapper">
        <draggable element="div" class="drag-list"
                   :options="options"
                   :move="onMove"
                   :v-model="list"
                   @start="handleStart"
                   @add="handleAdd"
                   @remove="handleRemove"
                   @choose="handleChoose"
                   @end="handleEnd"
                   @sort="handleSort"
                   @filter="handleFilter"
                   @clone="handleClone">
            <div class="drag-list-item" v-for=" (item,index) in list" :key="`drag_list_item${index}`">
                <slot v-bind="item"></slot>
            </div>
        </draggable>
    </div>
</template>

<script>
import draggable from 'vuedraggable'
export default {
  name: 'DragList',
  components: { draggable },
  props: {
    list: {
      type: Array,
      required: true
    },
    options: {
      type: Object
    }
  },
  computed: {
    dragOptions () {
      return Object.assign(this.defaultOption, this.options)
    }
  },
  data () {
    return {
      defaultOption: {
        sort: true,
        animation: 200,
        disabled: false,
        preventOnFilter: true
      }
    }
  },
  methods: {
    onMove (evt, draggedContext, relatedContext) {
      console.log('move------')
      console.log('evt')
      console.log('draggedContext')
      console.log('relatedContext')
    },
    // 开始拖动时的回调函数
    handleStart (event) {
      this.$emit('drag-start', event, this)
    },
    // 拖动结束时的回调函数
    handleEnd (event) {
      this.$emit('drag-end', event, this)
    },
    // 添加单元时的回调函数
    handleAdd (event) {
      this.$emit('drag-add', event, this)
    },
    // 单元被移动到另一个列表时的回调函数
    handleRemove (event) {
      this.$emit('drag-remove', event, this)
    },
    // 选择单元时的回调函数
    handleChoose (event) {
      this.$emit('drag-choose', event, this)
    },
    handleSort (event) {
      this.$emit('drag-sort', event, this)
    },
    // 尝试选择一个被filter过滤的单元的回调函数
    handleFilter (event) {
      this.$emit('drag-filter', event, this)
    },
    // clone时的回调函数
    handleClone (event) {
      this.$emit('drag-clone', event, this)
    }

  }
}
</script>

<style lang="scss" scoped>
.drag-list-wrapper{
    position: relative;
    width: 100%;
    height: 100%;
    border: 1px solid red;
    .drag-list{
        width: 100%;
        height: 100%;
        position: relative;
    }
}
</style>
