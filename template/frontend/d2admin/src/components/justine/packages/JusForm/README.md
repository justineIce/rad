##表单的简单封装
在element-ui上进行的封装，具体的传参可以参考element-ui传参方式
####参数设置
| 参数 | 说明 | 类型 | 默认值 |
|:-----|:-----|:-----|:-----|
|inline | 行内表单模式 | boolean | false |
|formOptions | 表单样式及保存按钮样式 |  Object   |
|formData | 表单模板| Object |
|formTemplate| 表单模板|Object |
|formRules |表单验证规则，具体的使用可以参照eleemnt-ui表单的规则验证 | Object|

####formOptions 参数设置
labelWidth
说明: 表单 label 宽度
类型: String
可选值: 无
默认值: 无

 labelPosition
说明: 表单 label 位置
类型: String
可选值: 无
默认值: 无

 saveButtonSize
说明: 保存按钮尺寸
类型: String
可选值: medium / small / mini
默认值: 无

 saveButtonType
说明: 保存按钮类型
类型: String
可选值: primary / success / warning / danger / info / text
默认值: 无

 saveButtonIcon
说明: 保存按钮图标类名
类型: String
可选值: 无
默认值: 无

 saveButtonText
说明: 保存按钮文字
类型: String
可选值: 无
默认值: 保存

 saveButtonColor
说明: 按钮颜色
类型: String
可选值: 无
默认值: 保存

# saveLoading
说明: 保存按钮是否加载中状态
类型: Boolean
可选值: 无
默认值: false



####formTemplate参数设置

title
说明: 表单 label 的名称
类型: String

value
说明: 表单模板的默认值

component.name
说明: 表单渲染的组件名，组件请参考 组件，自定义组件写法请参考自定义组件示例
类型: String
可选值: el-input / el-input-number / el-radio / el-checkbox / el-select / el-cascader / el-switch / el-slider / el-time-select / el-time-picker / el-date-picker / el-rate / el-color-picker / 自定义组件
            
component.options
说明: el-radio / el-checkbox / el-select 中的 options
类型: Array

component.show
说明: 是否显示此项
类型: Boolean

component.disabled
说明: 是否禁用此项
类型: Boolean

component.span
说明: 表单布局所占栅格数，栅格布局请参考 Layout 布局
类型: Number


component.offset
说明: 表单栅格布局偏移量，栅格布局请参考 Layout 布局
类型: Number
