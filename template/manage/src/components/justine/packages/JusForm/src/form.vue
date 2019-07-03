<template>
    <div class="jus-form" :class="{'justine-form-show' : type === 'show'}">
        <el-form v-bind="options" ref="form"
                 :model="formModel"
                 :rules="rules"
                 :label-position="options ? handleAttribute(options.labelPosition,null) : null"
                 :label-width="options ? handleAttribute(options.labelWidth,null) : null"
                 :inline="inline" :class="{'form-show': type === 'show'}">
            <el-row v-bind="options">
                <template v-for="(value, key) in formModel"  >
                    <el-col v-if="template[key] ? (template[key].component ? handleAttribute(template[key].component.show,true) : true) : false "
                            :span="template[key].component ? handleAttribute(template[key].component.span, inline ? null : 24) : (inline ? null : 24)"
                            :offset="template[key].component ? handleAttribute(template[key].component.offset, 0) : 0">
                        <el-form-item
                                :label="template[key].title + (type === 'show' ? '：' : '') "
                                :prop="key" >
                            <div v-if="type !== 'show'">
                                <slot :name="key" v-bind="template[key]">
                                    <!--输入框-->
                                    <el-input
                                            v-if="(!template[key].component) ||((!template[key].component.name) && (!template[key].component.render)) || template[key].component.name === 'el-input'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-input>
                                    <!--数字输入框-->
                                    <el-input-number
                                            v-else-if="template[key].component.name === 'el-input-number'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key,data: formModel})">
                                    </el-input-number>
                                    <!--单选按钮-->
                                    <el-radio-group
                                            v-else-if="template[key].component.name === 'el-radio'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                        <template v-if="template[key].component.buttonMode">
                                            <el-radio-button
                                                    v-for="option in template[key].component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-radio-button>
                                        </template>
                                        <template v-else>
                                            <el-radio
                                                    v-for="option in template[key].component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-radio>
                                        </template>
                                    </el-radio-group>
                                    <el-checkbox-group
                                            v-else-if="template[key].component.name === 'el-checkbox'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                        <template v-if="template[key].component.buttonMode">
                                            <el-checkbox-button
                                                    v-for="option in template[key].component.options"
                                                    :key="option.value"
                                                    :label="option.value"
                                            >
                                                {{option.label}}
                                            </el-checkbox-button>
                                        </template>
                                        <template v-else>
                                            <el-checkbox
                                                    v-for="option in template[key].component.options"
                                                    :key="option.value"
                                                    :label="option.value">
                                                {{option.label}}
                                            </el-checkbox>
                                        </template>
                                    </el-checkbox-group>
                                    <el-select
                                            v-else-if="template[key].component.name === 'el-select'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                        <el-option
                                                v-for="option in template[key].component.options"
                                                :key="option.value"
                                                v-bind="option"
                                        >
                                        </el-option>
                                    </el-select>
                                    <el-cascader
                                            v-else-if="template[key].component.name === 'el-cascader'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-cascader>
                                    <el-switch
                                            v-else-if="template[key].component.name === 'el-switch'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-switch>
                                    <el-slider
                                            v-else-if="template[key].component.name === 'el-slider'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-slider>
                                    <el-time-select
                                            v-else-if="template[key].component.name === 'el-time-select'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-time-select>
                                    <el-time-picker
                                            v-else-if="template[key].component.name === 'el-time-picker'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-time-picker>
                                    <el-date-picker
                                            v-else-if="template[key].component.name === 'el-date-picker'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-date-picker>
                                    <el-rate
                                            v-else-if="template[key].component.name === 'el-rate'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-rate>
                                    <el-color-picker
                                            v-else-if="template[key].component.name === 'el-color-picker'"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :disabled="handleAttributeDisabled(template[key].disabled,formModel)"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </el-color-picker>
                                    <render-custom-component
                                            v-else-if="template[key].component.name"
                                            v-model="formModel[key]"
                                            v-bind="template[key].component"
                                            :component-name="template[key].component.name"
                                            :props="template[key].component.props ? template[key].component.props : null"
                                            @change="$emit('change',{key: key, data: formModel})">
                                    </render-custom-component>
                                </slot>
                            </div>
                            <div v-if="type === 'show'">
                                <slot :name="key" v-bind="formModel">
                                    <label class="el-form-item__label">{{formModel ? (template[key].formatter ? template[key].formatter(formModel[key]) : formModel[key]): ''}}</label>
                                </slot>
                            </div>
                        </el-form-item>
                    </el-col>
                </template>
                <el-col  :span="inline ? null : 24" :offset="0">
                    <el-form-item>
                        <div class="form-buttons" style="display: flex;">
                            <!--默认的确认按钮-->
                            <el-button v-if="type !== 'show' && (options ? handleAttribute(options.saveButton, true) : true)"
                                       :style="{backgroundColor : options ? handleAttribute(options.saveButtonColor, null) : null ,'border-color' : options ? handleAttribute(options.saveButtonColor, null) : null  }"
                                       :size="options ? handleAttribute(options.saveButtonSize, null) : null"
                                       :type="options ? handleAttribute(options.saveButtonType, 'primary') : 'primary'"
                                       :icon="options ? handleAttribute(options.saveButtonIcon, null) : null"
                                       :loading="saveLoading"
                                       @click="handleFormSave"
                            >{{options ? handleAttribute(options.saveButtonText, '确定') : '确定'}}
                            </el-button>
                            <!--其他的按钮-->
                            <slot name="custom-button" v-bind="formModel"></slot>
                        </div>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
    </div>
</template>

<script>
import renderCustomComponent from './component/renderCustomComponent.vue'
import utils from './mixin/utils'
export default {
  name: 'JusForm',
  mixins: [
    utils
  ],
  components: {
    renderCustomComponent
  },
  props: {
    /**
    * @description  是否内联表单
    * */
    inline: {
      type: Boolean,
      default: false
    },
    /**
     *@description 表单类型 show显示
     **/
    type: {
      type: String,
      default: ''
    },
    /**
     * @description dialog配置
     */
    options: {
      type: Object,
      default: null
    },
    /**
    * @description data
    */
    data: {
      type: Object,
      default: null
    },
    /**
    * @description 表单模板
    */
    template: {
      type: Object,
      default: null,
      request: true
    },
    /**
     * @description 表单校验规则
     */
    rules: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      saveLoading: false,
      formModel: {}
    }
  },
  watch: {
    data (newVal, oldVal) {
      this.formModel = Object.assign(this.initFormData(), this.data)
    },
    template (newVal, oldVal) {
      if (!this.data || JSON.stringify(this.data) === '{}') {
        this.formModel = this.initFormData()
      }
    }
  },
  methods: {
    initFormData () {
      if (this.template) {
        let data = {}
        for (let prop in this.template) {
          data[prop] = this.template[prop].value ? this.template[prop].value : ''
        }
        return data
      }
    },
    /**
    * @description 保存行数据
    */
    handleFormSave () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        } else {
          this.saveLoading = this.options ? this.handleAttribute(this.options.saveLoading, false) : false
          let rowData = {}
          for (let prop in this.formModel) {
            rowData[prop] = this.formModel[prop]
          }
          // this.data = rowData
          this.$emit('form-submit', rowData, () => {
            this.saveLoading = false
          })
        }
      })
    }
  },
  created () {
    if (!this.data || JSON.stringify(this.data) === '{}') {
      this.formModel = this.initFormData()
    } else {
      this.formModel = Object.assign(this.initFormData(), this.data)
    }
  },
  mounted () {

  }
}
</script>
