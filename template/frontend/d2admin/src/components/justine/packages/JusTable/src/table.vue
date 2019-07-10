<!-- 表格-->
<template>
    <div class="jus-table">
        <div class="jus-table-container">
            <!--条件查询-->
            <div class="jus-table-container-query" v-if="queryForm || addMode">
                <JusForm v-if="queryFormTemplate || addMode"
                         :inline="true"
                         :template="queryFormTemplate"
                         :data="queryFormData"
                         :options="{saveButtonText:'查询'}"
                         @form-submit="handleQueryMethod">
                    <!--slot 自定义按钮-->
                    <template slot="custom-button">
                        <el-button size="small" v-show="addMode" @click="handleAdd">+新增</el-button>
                        <slot name="custom-button"></slot>
                    </template>
                </JusForm>
            </div>
            <!--列表内容-->
            <div class="jus-table-container-table">
                <slot name="table">
                    <div style="padding: 10px;border: 1px solid #eee;" v-if="$slots.header">
                        <slot name="header"></slot>
                    </div>
                    <el-table role="table"
                              ref="table"
                              v-bind="options"
                              :data="data"
                              :header-cell-style="{padding: '10px 0', background: '#f3f3f3'}"
                              @select="handleSelect"
                              @select-all="handleSelectAll"
                              @selection-change="handleSelectionChange"
                              @cell-mouse-enter="handleCellMouseEnter"
                              @cell-mouse-leave="handleCellMouseLeave"
                              @cell-click="handleCellClick"
                              @cell-dblclick="handleCellDblclick"
                              @row-click="handleRowClick"
                              @row-contextmenu="handleRowContextmenu"
                              @row-dblclick="handleRowDblclick"
                              @header-click="handleHeaderClick"
                              @header-contextmenu="handleHeaderContextmenu"
                              @sort-change="handleSortChange"
                              @filter-change="handleFilterChange"
                              @current-change="handleCurrentChange">
                        <!--列表信息-->
                        <!--全选-->
                        <el-table-column v-if="indexRow" type="index" width="45"></el-table-column>
                        <el-table-column v-if="expand" type="expand" width="45">
                            <template slot-scope="props">
                                <slot name="expand"  v-bind="props.row"></slot>
                            </template>
                        </el-table-column>
                        <el-table-column
                                v-if="selectionRow || selectionRow === ''"
                                type="selection" width="30"
                                :label="handleAttribute(selectionRow.title, '')"
                                v-bind="selectionRow">
                        </el-table-column>
                        <el-table-column v-for="(item, index) in columns"
                                         v-if="item.onlyQuery ? !item.onlyQuery : true "
                                         v-bind="item"
                                         :prop="item.key"
                                         :label="item.title"
                                         :key="index"
                                         :filters="item.filters ? item.filters : null"
                                         :filter-method="item.filterMethod ? item.filterMethod : null">
                            <template slot-scope="scope">
                                <div v-if="isEdit || (item.component && item.component.isEdit && handleRowHandleButtonShow(item.component.isEdit,scope.$index, scope.row)) ">
                                    <el-input
                                            v-if="item.component && item.component.name === 'el-input' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-input>
                                    <el-input-number
                                            v-else-if="item.component && item.component.name === 'el-input-number' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-input-number>
                                    <el-radio-group
                                            v-else-if="item.component && item.component.name === 'el-radio' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                        <template v-if="item.component.buttonMode">
                                            <el-radio-button
                                                    v-for="option in item.component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-radio-button>
                                        </template>
                                        <template v-else>
                                            <el-radio
                                                    v-for="option in item.component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-radio>
                                        </template>
                                    </el-radio-group>
                                    <el-checkbox-group
                                            v-else-if="item.component && item.component.name === 'el-checkbox' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                        <template v-if="item.component.buttonMode">
                                            <el-checkbox-button
                                                    v-for="option in item.component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-checkbox-button>
                                        </template>
                                        <template v-else>
                                            <el-checkbox
                                                    v-for="option in item.component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-checkbox>
                                        </template>
                                    </el-checkbox-group>
                                    <el-select
                                            v-else-if="item.component && item.component.name === 'el-select' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                        <el-option
                                                v-for="option in item.component.options"
                                                :key="option.value"
                                                v-bind="option"
                                        >
                                        </el-option>
                                    </el-select>
                                    <el-cascader
                                            v-else-if="item.component && item.component.name === 'el-cascader' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-cascader>
                                    <el-switch
                                            v-else-if="item.component && item.component.name === 'el-switch' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-switch>
                                    <el-slider
                                            v-else-if="item.component && item.component.name === 'el-slider' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-slider>
                                    <el-time-select
                                            v-else-if="item.component && item.component.name === 'el-time-select' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-time-select>
                                    <el-time-picker
                                            v-else-if="item.component && item.component.name === 'el-time-picker' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-time-picker>
                                    <el-date-picker
                                            v-else-if="item.component && item.component.name === 'el-date-picker' && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            v-bind="item.component"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})"
                                    >
                                    </el-date-picker>
                                    <render-custom-component
                                            v-else-if="item.component  && item.component.name && handleRowHandleRelation(item.component.relation, scope.$index, scope.row)"
                                            v-model="scope.row[item.key]"
                                            :component-name="item.component.name"
                                            :props="item.component.props ? item.component.props : null"
                                            :scope="scope"
                                            @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})">
                                    </render-custom-component>
                                    <slot :name="item.key" v-bind="scope.row" v-else>{{scope.row[item.key]}}</slot>
                                </div>
                                <render-custom-component
                                        v-else-if="!isEdit && item.component && item.component.name && item.component.show"
                                        v-model="scope.row[item.key]"
                                        :component-name="item.component.name"
                                        :props="item.component.props ? item.component.props : null"
                                        :pProps="{isEdit:isEdit}"
                                        :scope="scope"
                                        @change="$emit('row-data-change', {rowIndex: scope.$index, key: item.key, value: scope.row[item.key], row: scope.row})">
                                </render-custom-component>
                                <slot :name="item.key" v-bind="scope.row" v-else>{{item.formatter ? item.formatter(scope.row, scope.column, scope.row[item.key], scope.$index) : scope.row[item.key]}}</slot>
                            </template>
                        </el-table-column>
                        <!--操作栏-->
                        <el-table-column v-if="rowHandle" v-bind="rowHandle" :label="handleAttribute(rowHandle.columnHeader, '操作')">
                            <template slot-scope="scope">
                                <div flex>
                                    <div v-for="(item, index) in handleAttribute(rowHandle.custom, [])" :key="index">
                                        <el-button  size="mini"
                                                    v-bind="item"
                                                    v-if="handleRowHandleButtonShow(item.show, scope.$index, scope.row)"
                                                    :disabled="handleRowHandleButtonDisabled(item.disabled, scope.$index, scope.row)"
                                                    @click="$emit(item.emit, {index: scope.$index, row: scope.row})">{{item.text}}</el-button>
                                    </div>
                                    <el-button  v-if="rowHandle.edit && handleRowHandleButtonShow(rowHandle.edit.show, scope.$index, scope.row)"
                                                :disabled="handleRowHandleButtonDisabled(rowHandle.edit.disabled, scope.$index, scope.row)"
                                                v-bind="rowHandle.edit"
                                                size="mini"
                                                @click="handleEdit(scope.$index, scope.row)">{{handleAttribute(rowHandle.edit.text, '编辑')}}</el-button>

                                    <el-button  v-if="rowHandle.remove && handleRowHandleButtonShow(rowHandle.remove.show, scope.$index, scope.row)"
                                                :disabled="handleRowHandleButtonDisabled(rowHandle.remove.disabled, scope.$index, scope.row)"
                                                v-bind="rowHandle.remove"
                                                size="mini"
                                                @click="handleRemove(scope.$index, scope.row)">{{handleAttribute(rowHandle.remove.text, '删除')}}</el-button>
                                </div>
                            </template>
                        </el-table-column>
                    </el-table>
                </slot>
            </div>
            <!--翻页-->
            <div style="padding: 10px 5px;" v-if="$slots.foot">
                <slot name="foot"></slot>
            </div>
            <div flex="main:justify"  class="jus-table-container-pagination" v-if="showPagination">
                <div><slot name="extend-handle"></slot></div>
                <el-pagination v-if="showPagination"
                               layout="total, prev, pager, next,jumper"
                               v-bind="pagination"
                               :total="pagination ? pagination.total : this.data.length"
                               @current-change="handleCurrentPage"/>
            </div>
        </div>
        <!--添加/修改表单 弹出层-->
        <el-dialog :title="formMode === 'edit' ? editTitle : addTitle" width="40%"
                   :visible.sync="dialogVisible"
                   @close="dialogVisible = false"
                   append-to-body>
            <JusForm
                    :options="formOptions ? formOptions : {labelWidth: '100px'}"
                    :template="formTemplate"
                    :data="formData"
                    :rules="formRules"
                    @form-submit="formSubmit"></JusForm>
        </el-dialog>
    </div>
</template>

<script>
import form from './mixin/form'
import base from './mixin/base'
import utils from './mixin/utils'
import renderCustomComponent from './component/renderCustomComponent.vue'
export default {
  name: 'JusTable',
  mixins: [ base, form, utils ],
  components: {
    renderCustomComponent
  },
  props: {

  },
  data () {
    return {

    }
  },
  computed: {
    /**
     * @description 获取查询条件数据
     * 根据表头数据中的isQuery参数
     */
    queryFormTemplate () {
      if (this.columns) {
        let data = {}
        this.columns.forEach((item) => {
          if (item.isQuery !== undefined && typeof item.isQuery === 'boolean' && item.isQuery) {
            data[item.key] = item
          }
        })
        return JSON.stringify(data) === '{}' ? null : data
      }
      return null
    }
  },
  mounted () {

  }
}
</script>
<style lang="scss">
    @import "table.scss";
</style>
