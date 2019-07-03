##表格的基本封装
在element-ui上进行的封装，具体的传参可以参考element-ui传参方式
####参数设置
| 参数 | 说明 | 类型 | 默认值 |
|:-----|:-----|:-----|:-----|
| title             | 表明                | string |
| menuData          | 菜单数据            |  Array  | 
| menuDefaultProps  | 菜单取值项          | Object|
| queryFormTemplate | 查询表单设置(如果queryFormData没有值,则取自己的字段) |Object |
| queryFormData     | 查询表单数据          | Object|
| columns           | 表头数据              | Object |
| tableData         | 表的数据              | Array |
| total             | 数据总条数(计算翻页)  | Number|

####方法
| 方法名 | 说明 | 回调参数 |
|:-----|:-----|:-----|
|query-data  | 查询数据方法| 共二个参数,依次为：查询条件，查询页|
|row-edit    | 行操作      |共三个参数,依次为：类型（info 详情按钮 edit编辑按钮 delete 删除按钮）、行数、行数据|
|add-handle  | 添加按钮    |
|menu-handle | 菜单操作    |共三个参数，依次为：传递给 data 属性的数组中该节点所对应的对象、节点对应的 Node、节点组件本身|
|selection-change | 多选方法 | selection


###slot
| 名称 | 说明 |
|:----|:----|
| custom-from-button |  查询表单，自定义的其他按钮 |
| table   |  重写表格                   |
| columns的key 值| 对每列字段名进行了插槽设计，方便重写样式     |

