<!--菜单管理-->
<template>
    <d2-container class="meta-public">
        <JusLayout aside :aside-width="200" title="菜单管理">
            <!-- 侧面栏 -->
            <template slot="aside">

                <div class="sub-title">顶部菜单</div>
                <!-- 菜单 -->
                <div>
                    <!--这里只显示顶级菜单-->
                    <jus-tree style="background-color: transparent;"
                              :data="menuOption"
                              :props="props"
                              defaultExpandAll
                              @node-click="nodeClick">
                        <span class="custom-tree-node" slot-scope="{ node, data }">
                            <span>{{ node.label }}</span>
                            <span class="hide_">
                                <span class="el-icon-edit mar_right" @click="() => handleEditCatalog(node, data)"></span>
                                <span class="el-icon-delete" @click="() => handleRemoveCatalog(node, data)"></span>
                            </span>
                        </span>
                    </jus-tree>
                </div>
                <div class="border_"></div>
                <div class="add_" @click="addMenu">
                    <span class="el-icon-plus pad"> 添加新菜单</span>
                </div>
            </template>

            <!--正文部分-->
            <div class="box-content">
                <el-row :gutter="10">
                    <el-col :span="14">
                        <div class="box-content-first">
                            <jus-tree
                                    :data="allMenuOption"
                                    :props="props"
                                    @node-click="nodeClick1"
                                    defaultExpandAll></jus-tree>
                        </div>
                    </el-col>
                    <el-col :span="10">
                        <div class="box-content-second">
                            <div class="box_1">
                                <span class="title_">{{checkedMenu}}</span>
                                <span class="el-icon-delete el-float" @click="delMenu"></span>
                            </div>
                            <div class="mar_top">
                                <el-form :label-width="labelWidth" :model="menuForm1" ref="menuForm1" :rules="menuForm1Rules">
                                    <el-form-item label="图标：" prop="shortcutIcon">
                                        <el-input class="icon_box" v-model="menuForm1.shortcutIcon">
                                            <template slot="prepend">
                                                <el-icon-popover :icon="menuForm1.shortcutIcon" @getIcon="getIcon"></el-icon-popover>
                                            </template>
                                        </el-input>
                                    </el-form-item>
                                    <el-form-item label="资源名称：" prop="name">
                                        <el-input v-model="menuForm1.name"></el-input>
                                    </el-form-item>
                                    <el-form-item label="资源类型：" prop="type">
                                        <el-select v-model="menuForm1.type">
                                            <el-option v-for="item in menuTypeOption" :key="item.value" :value="item.value" :label="item.label"></el-option>
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="资源路径：" prop="url">
                                        <el-input v-model="menuForm1.url"></el-input>
                                    </el-form-item>
                                    <el-form-item label="请求方法：">
                                        <el-input v-model="method" class="input-with-select">
                                            <el-button slot="append" type="success" @click="addMethod">添加</el-button>
                                        </el-input>
                                        <div class="method_box">
                                            <div class="m_box" v-for="(item,index) in methodArr" :key="index">
                                                <span>{{item}}</span>
                                                <span class="el-icon-delete del_" @click="delMethod(index)"></span>
                                            </div>
                                        </div>
                                    </el-form-item>
                                    <el-form-item label="描述：">
                                        <el-input v-model="menuForm1.remarks" type="textarea" :rows="2"></el-input>
                                    </el-form-item>
                                </el-form>
                            </div>
                            <div>
                                <div class="title_ box_1 mar_top">操作权限</div>
                                <div class="mar_top">
                                    <el-button type="warning" @click="lookOperateLimit">操作权限</el-button>
                                    <el-button type="primary" @click="addOperateLimit">添加操作权限</el-button>
                                </div>
                            </div>

                        </div>
                    </el-col>
                </el-row>
            </div>
            <div class="box-btn">
                <span class="btn-float">
                    <el-button type="primary" @click="builtChildMenu">新增子菜单</el-button>
                    <el-button type="success" @click="remainAll">保存</el-button>
                </span>
            </div>

            <!--添加新菜单-->
            <el-dialog :title="title" :visible.sync="menuDialogVisible">
                <el-form :model="menuForm" ref="menuForm" :rules="menuFormRules" :label-width="labelWidth">
                    <el-form-item label="图标：">
                        <el-input class="icon_box" v-model="menuForm.shortcutIcon">
                            <template slot="prepend">
                                <el-icon-popover :icon="menuForm.shortcutIcon" @getIcon="getMenuIcon"></el-icon-popover>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item label="资源名称：" prop="name">
                        <el-input v-model="menuForm.name"></el-input>
                    </el-form-item>
                    <el-form-item label="资源类型：" prop="type">
                        <el-select v-model="menuForm.type">
                            <el-option v-for="item in menuTypeOption" :key="item.value" :value="item.value" :label="item.label"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="资源路径：" prop="url">
                        <el-input v-model="menuForm.url"></el-input>
                    </el-form-item>
                    <el-form-item label="资源描述：">
                        <el-input v-model="menuForm.remarks" type="textarea" :rows="3"></el-input>
                    </el-form-item>
                    <el-form-item label="排序号：">
                        <el-input placeholder="用于主菜单排序" v-model="menuForm.sort"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submit">保存</el-button>
                        <el-button @click="menuDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>

            <!--新增子菜单-->
            <el-dialog title="新增菜单" :visible.sync="childMenuDialogVisible">
                <el-form :label-width="labelWidth" :model="childMenuForm" ref="childMenuForm" :rules="childMenuFormRules">
                    <el-form-item label="父级资源名称：">
                        <el-input disabled v-model="childMenuForm.parentMenu"></el-input>
                    </el-form-item>
                    <el-form-item label="图标：" prop="shortcutIcon">
                        <el-input class="icon_box" v-model="childMenuForm.shortcutIcon">
                            <template slot="prepend">
                                <el-icon-popover :icon="childMenuForm.shortcutIcon" @getIcon="getChildMenuIcon"></el-icon-popover>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item label="资源名称：" prop="name">
                        <el-input v-model="childMenuForm.name"></el-input>
                    </el-form-item>
                    <el-form-item label="资源类型：" prop="type">
                        <el-select v-model="childMenuForm.type">
                            <el-option v-for="item in menuTypeOption" :key="item.value" :value="item.value" :label="item.label"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="资源路径：" prop="url">
                        <el-input v-model="childMenuForm.url"></el-input>
                    </el-form-item>
                    <el-form-item label="请求方法：">
                        <el-input v-model="childMenuForm.method" type="textarea" :rows="3" placeholder="请输入请求方法，多个用英文逗号隔开"></el-input>
                    </el-form-item>
                    <el-form-item label="资源描述：">
                        <el-input v-model="childMenuForm.remarks" type="textarea" :rows="2"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submitChildMenu">保存</el-button>
                        <el-button @click="childMenuDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>

            <!--添加/编辑操作权限-->
            <el-dialog :title="btnTitle" :visible.sync="addOperateLimitDialogVisible">
                <el-form :label-width="labelWidth" :model="operateLimit" ref="operateLimit" :rules="operateLimitRules">
                    <el-form-item label="父级名称：">
                        <el-input v-model="operateLimit.parentName" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="图标：">
                        <el-input class="icon_box" v-model="operateLimit.shortcutIcon">
                            <template slot="prepend">
                                <el-icon-popover :icon="operateLimit.shortcutIcon" @getIcon="getOperateLimitIcon"></el-icon-popover>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item label="名称：" prop="name">
                        <el-select
                                v-model="operateLimit.name"
                                filterable
                                allow-create
                                default-first-option
                                placeholder="请选择或请输入">
                            <el-option
                                    v-for="item in baseBtnOption"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="类型：" prop="type">
                        <el-select v-model="operateLimit.type">
                            <el-option v-for="item in btnTypeOption" :key="item.value" :value="item.value" :label="item.label"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="请求方法：">
                        <el-input v-model="limitMethod" class="input-with-select">
                            <el-button slot="append" type="success" @click="addLimitMethod">添加</el-button>
                        </el-input>
                        <div class="method_box">
                            <div class="m_box" v-for="(item,index) in limitMethodArr" :key="index">
                                <span>{{item}}</span>
                                <span class="el-icon-delete del_" @click="delLimitMethod(index)"></span>
                            </div>
                        </div>
                    </el-form-item>
                    <el-form-item label="权限标识：" prop="mark">
                        <el-input v-model="operateLimit.mark"></el-input>
                    </el-form-item>
                    <el-form-item label="描述：">
                        <el-input v-model="operateLimit.remarks"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submitLimit">保存</el-button>
                        <el-button @click="addOperateLimitDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>

            <!--操作权限详情-->
            <el-dialog title="操作权限" :visible.sync="OperateLimitDialogVisible">
                <JusTable ref="table"
                          :columns="columns"
                          :data="data"
                          :row-handle="rowHandle"
                          :pagination="pagination"
                          @row-remove="rowRemove"
                          @editRow="editRow"
                          @query-data="handleQueryData">
                </JusTable>
                <div class="align">
                    <el-button type="primary" @click="OperateLimitDialogVisible=false">关闭</el-button>
                </div>
            </el-dialog>

        </JusLayout>
    </d2-container>
</template>
<script>
import { GetTopMenuList, AddTopLevelMenu, DelTopLevelMenu, GetAllMenuList, GetBtnMenuList } from '@/api/platform.manage'
export default {
  name: 'resourceMange',
  data () {
    return {
      // 顶级菜单
      menuOption: [],
      props: {
        label: 'name',
        icon: 'shortcut_icon',
        children: 'Children'
      },

      // 添加新菜单
      title: '添加新菜单',
      menuDialogVisible: false,
      labelWidth: '120px',
      menuForm: {
        id: '',
        shortcutIcon: '',
        name: '',
        url: '',
        type: 'menu',
        remarks: '',
        sort: '0'
      },
      // 资源类型
      menuTypeOption: [
        { value: 'menu', label: 'menu' }
      ],
      menuFormRules: {
        name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
        type: [{ required: true, message: '请选择资源类型', trigger: 'blur' }],
        url: [{ required: true, message: '请输入资源路径', trigger: 'blur' }]
      },
      // 根据左侧菜单选择生成
      allMenuOption: [],
      // 顶级菜单 - 隐藏部分元素
      firstClass: true,

      // 当前被选中的数据 - 编辑删除
      checkedMenu: '',
      menuForm1: {
        id: '',
        parentId: '',
        name: '',
        shortcutIcon: '',
        url: '',
        type: '',
        method: '',
        remarks: '',
        sort: ''
      },
      menuForm1Rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        shortcutIcon: [{ required: true, message: '请输入图标', trigger: 'blur' }],
        url: [{ required: true, message: '请输入资源路径', trigger: 'blur' }],
        type: [{ required: true, message: '请选择资源类型', trigger: 'blur' }]
      },
      method: '',
      methodArr: [],

      // 新增菜单
      childMenuDialogVisible: false,
      childMenuForm: {
        name: '',
        url: '',
        parentMenu: '',
        parentId: '',
        shortcutIcon: 'fa fa-file-o',
        remarks: '',
        method: '',
        type: 'menu'
      },
      childMenuFormRules: {
        name: [{ required: true, message: '请输入资源名称', trigger: 'blur' }],
        url: [{ required: true, message: '请输入资源路径', trigger: 'blur' }],
        type: [{ required: true, message: '请选择资源类型', trigger: 'blur' }]
      },

      // 添加操作权限
      btnTitle: '添加操作权限',
      addOperateLimitDialogVisible: false,
      operateLimit: {
        id: '',
        parentId: '',
        parentName: '',
        name: '',
        type: 'button',
        shortcutIcon: '',
        method: '',
        mark: '',
        remarks: ''
      },
      btnTypeOption: [
        { value: 'button', label: 'button' }
      ],
      operateLimitRules: {
        name: [{ required: true, message: '请输入操作名称', trigger: 'blur' }],
        type: [{ required: true, message: '请选择资源类型', trigger: 'blur' }],
        mark: [{ required: true, message: '请输入权限标识', trigger: 'blur' }]
      },
      limitMethod: '',
      limitMethodArr: [],

      // 操作权限详情
      OperateLimitDialogVisible: false,
      columns: [
        { title: '名称', key: 'name' },
        { title: '图标', key: 'shortcut_icon' },
        { title: '类型', key: 'type' },
        { title: '请求方法', key: 'method' },
        { title: '权限标识', key: 'mark' },
        { title: '描述', key: 'remarks' }
      ],
      data: [],
      rowHandle: {
        remove: {},
        custom: [
          { text: '编辑', icon: '', emit: 'editRow' }
        ]
      },
      pagination: {
        pageIndex: 1,
        pageSize: 10,
        total: 0
      },
      page: {
        pageIndex: '',
        pageSize: ''
      },

      // 用于当前最顶级菜单的id
      currentTopMenuId: '',

      visible2: false,
      iconOption: [
        { icon: 'el-icon-info' },
        { icon: 'el-icon-error' },
        { icon: 'el-icon-success' },
        { icon: 'el-icon-warning' },
        { icon: 'el-icon-question' },
        { icon: 'el-icon-back' }
      ],

      baseBtnOption: [
        { value: '增加', label: '增加' },
        { value: '删除', label: '删除' },
        { value: '编辑', label: '编辑' }
      ]
    }
  },
  methods: {
    // 编辑菜单
    handleEditCatalog (node, data) {
      console.log(node, data)
      this.title = '编辑菜单'
      this.menuDialogVisible = true
      this.menuForm = {
        id: data.id,
        parentId: data.parentId,
        shortcutIcon: data.shortcut_icon,
        url: data.url,
        name: data.name,
        type: data.type,
        remarks: data.remarks,
        sort: data.sort
      }
    },
    // 删除菜单
    handleRemoveCatalog (node, data) {
      this.$confirm('确定删除此菜单吗？', '删除提醒', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 具体操作
        DelTopLevelMenu({ id: data.id }).then((res) => {
          if (res.ret === 200) {
            this.$message.success('删除成功')
            this.getTopMenuList()
          }
        })
      }).catch(() => {
        this.$message.warning('取消删除')
      })
    },
    /**
       * 数据， 节点， 当前对象
       * @param data
       * @param node
       * @param vim
       */
    nodeClick (data, node, vim) {
      // if (data.parent_id == 0) {
      //   this.firstClass = true
      // } else {
      //   this.firstClass = false
      // }
      this.currentTopMenuId = data.id
      // 从后台动态获取
      this.getAllMenuList(data.id)
      this.checkedMenu = data.name
      this.menuForm1 = {
        id: data.id,
        parentId: data.parent_id,
        name: data.name,
        url: data.url,
        shortcutIcon: data.shortcut_icon,
        type: data.type,
        method: data.method,
        remarks: data.remarks
      }
      if (data.method.length > 0) {
        this.methodArr = data.method.split(',')
      } else {
        this.methodArr = []
      }
      this.operateLimit.parentId = data.id
      this.operateLimit.parentName = data.name
    },

    // 添加新菜单
    addMenu () {
      this.menuDialogVisible = true
      this.title = '添加菜单'
      this.menuForm = {
        id: '',
        parentId: '0',
        shortcutIcon: '',
        name: '',
        type: 'menu',
        remarks: '',
        sort: '0'
      }
    },
    // 保存
    submit () {
      this.$refs.menuForm.validate((valid) => {
        if (valid) {
          // 添加/编辑新菜单
          AddTopLevelMenu(this.menuForm).then((res) => {
            console.log(res)
            if (res.ret === 200) {
              this.menuDialogVisible = false
              this.$message.success('保存成功')
              this.getTopMenuList()
            }
          })
        }
      })
    },
    getMenuIcon (icon) {
      this.menuForm.shortcutIcon = icon
    },

    // 正文内容菜单点击
    nodeClick1 (data, node, vim) {
      // if (data.parent_id == 0) {
      //   this.firstClass = true
      // } else {
      //   this.firstClass = false
      // }
      this.checkedMenu = data.name
      this.menuForm1 = {
        id: data.id,
        parentId: data.parent_id,
        name: data.name,
        shortcutIcon: data.shortcut_icon,
        url: data.url,
        type: data.type,
        method: data.method,
        remarks: data.remarks,
        sort: data.sort
      }
      if (data.method.length > 0) {
        this.methodArr = data.method.split(',')
      } else {
        this.methodArr = []
      }
      this.operateLimit.parentId = data.id
      this.operateLimit.parentName = data.name
    },

    // 新增子菜单
    builtChildMenu () {
      this.childMenuForm = {
        name: '',
        url: '',
        parentMenu: '',
        parentId: '',
        shortcutIcon: 'fa fa-file-o',
        remarks: '',
        method: '',
        type: 'menu'
      }
      this.childMenuForm.parentId = this.menuForm1.id
      this.childMenuForm.parentMenu = this.menuForm1.name

      this.childMenuDialogVisible = true
    },

    // 全部保存
    remainAll () {
      this.$refs.menuForm1.validate((valid) => {
        if (valid) {
          AddTopLevelMenu(this.menuForm1).then((res) => {
            console.log(res)
            if (res.ret === 200) {
              this.$message.success('保存成功')
              this.getTopMenuList()
            }
          })
        }
      })
    },

    // 保存菜单
    submitChildMenu () {
      this.$refs.childMenuForm.validate((valid) => {
        if (valid) {
          AddTopLevelMenu(this.childMenuForm).then((res) => {
            console.log(res)
            if (res.ret === 200) {
              this.childMenuDialogVisible = false
              this.$message.success('保存成功')
              this.getTopMenuList()
            }
          })
        }
      })
    },
    getChildMenuIcon (icon) {
      this.childMenuForm.shortcutIcon = icon
    },

    // 删除菜单
    delMenu () {
      this.$confirm('确定删该菜单吗？', '删除提醒', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 具体操作
        DelTopLevelMenu({ id: this.menuForm1.id }).then((res) => {
          if (res.ret === 200) {
            this.$message.success('删除成功')
            this.getAllMenuList(this.currentTopMenuId)
          }
        })
      }).catch(() => {
        this.$message.warning('取消删除')
      })
    },

    // 添加操作权限
    addOperateLimit () {
      this.btnTitle = '添加操作权限'
      this.addOperateLimitDialogVisible = true
      this.operateLimit.name = ''
      this.operateLimit.type = 'button'
      this.operateLimit.shortcutIcon = ''
      this.operateLimit.method = ''
      this.operateLimit.remarks = ''
    },
    // 保存操作权限
    submitLimit () {
      this.$refs.operateLimit.validate((valid) => {
        if (valid) {
          if (this.operateLimit.method.length <= 0) {
            this.$message.warning('请输入请求方法')
          } else {
            AddTopLevelMenu(this.operateLimit).then((res) => {
              console.log(res)
              if (res.ret === 200) {
                this.addOperateLimitDialogVisible = false
                this.$message.success('保存成功')
                this.lookOperateLimit()
              }
            })
          }
        }
      })
    },
    getOperateLimitIcon (icon) {
      this.operateLimit.shortcutIcon = icon
    },
    // 添加请求方法
    addLimitMethod () {
      if (this.limitMethod.trim().length <= 0) {
        this.$message.warning('请输入请求方法')
        return
      }
      this.limitMethodArr.push(this.limitMethod)
      this.limitMethod = ''
      this.operateLimit.method = this.limitMethodArr.join(',')
    },
    // 删除请求方法
    delLimitMethod (index) {
      for (var i in this.limitMethodArr) {
        if (index == i) {
          this.limitMethodArr.splice(index, 1)
        }
      }
      this.operateLimit.method = this.limitMethodArr.join(',')
    },

    // 操作权限详情
    handleQueryData (data, currentPage) {
      this.page = {
        pageIndex: currentPage.pageIndex,
        pageSize: currentPage.pageSize
      }
      this.lookOperateLimit()
    },
    // 删除操作权限
    rowRemove ({ index, row }) {
      DelTopLevelMenu({ id: row.id }).then((res) => {
        if (res.ret === 200) {
          this.$message.success('删除成功')
          this.lookOperateLimit()
        }
      })
    },
    // 编辑操作权限
    editRow ({ index, row }) {
      console.log(index, row)
      this.btnTitle = '编辑操作权限'
      this.addOperateLimitDialogVisible = true
      this.operateLimit = {
        id: row.id,
        parentId: row.parent_id,
        parentName: this.menuForm1.name,
        name: row.name,
        type: 'button',
        shortcutIcon: row.short_icon,
        method: row.method,
        mark: row.mark,
        remarks: row.remarks
      }
      if (row.method.length > 0) {
        this.limitMethodArr = row.method.split(',')
      } else {
        this.limitMethodArr = []
      }
    },
    lookOperateLimit () {
      this.OperateLimitDialogVisible = true
      // 获取所具体菜单下的所有按钮操作权限
      GetBtnMenuList({
        parentId: this.menuForm1.id,
        page: 'page',
        pageIndex: this.page.pageIndex,
        pageSize: this.page.pageSize
      }).then((res) => {
        if (res.ret === 200) {
          var data = res.data
          this.data = data.data
          this.pagination = {
            currentPage: data.pageIndex,
            pageSize: data.pageSize,
            total: data.Count
          }
        }
      })
    },
    addMethod () {
      if (this.method.trim().length <= 0) {
        this.$message.warning('请输入请求方法')
        return
      }
      this.methodArr.push(this.method)
      this.method = ''
      this.menuForm1.method = this.methodArr.join(',')
      console.log(this.menuForm1.method)
    },
    delMethod (index) {
      for (var i in this.methodArr) {
        if (index == i) {
          this.methodArr.splice(index, 1)
        }
      }
      this.menuForm1.method = this.methodArr.join(',')
      console.log(this.menuForm1.method)
    },

    // 获取顶部菜单
    getTopMenuList () {
      GetTopMenuList().then((res) => {
        if (res.ret === 200) {
          this.menuOption = res.data
          var firtMenu = this.menuOption[0]
          this.getAllMenuList(firtMenu.id)
          this.checkedMenu = firtMenu.name
          this.currentTopMenuId = firtMenu.id
          this.menuForm1 = {
            id: firtMenu.id,
            parentId: firtMenu.parent_id,
            name: firtMenu.name,
            shortcutIcon: firtMenu.shortcut_icon,
            url: firtMenu.url,
            type: firtMenu.type,
            method: firtMenu.method,
            remarks: firtMenu.remarks,
            sort: firtMenu.sort
          }
          if (firtMenu.method > 0) {
            this.methodArr = firtMenu.method.split(',')
          }
          this.operateLimit.parentId = firtMenu.id
          this.operateLimit.parentName = firtMenu.name
        }
      })
    },
    // 获取全部菜单结构
    getAllMenuList (id) {
      GetAllMenuList({ id: id }).then((res) => {
        if (res.ret === 200) {
          console.log(res.data)
          this.allMenuOption = res.data
          var firtMenu = this.allMenuOption[0]
          this.checkedMenu = firtMenu.name
          this.menuForm1 = {
            id: firtMenu.id,
            parentId: firtMenu.parent_id,
            name: firtMenu.name,
            shortcutIcon: firtMenu.shortcut_icon,
            url: firtMenu.url,
            type: firtMenu.type,
            method: firtMenu.method,
            remarks: firtMenu.remarks,
            sort: firtMenu.sort
          }
          if (firtMenu.method.length > 0) {
            this.methodArr = firtMenu.method
          }
        }
      })
    },

    // 图标选择
    getIcon (icon) {
      this.menuForm1.shortcutIcon = icon
    }
  },
  mounted () {
    this.getTopMenuList()
  }
}
</script>
<style scoped>
    .sub-title {
        padding: 10px 20px;
        margin: 20px 0 10px;
        font-weight: 700;
        font-size: 15px;
        color: #526069;
        cursor: default;
    }
    .pad {
        margin-left: 10px;
        cursor: pointer;
        display: inline-block;
        font-size: 15px;
    }
    .add_ {
        padding: 10px 0;
        cursor: pointer;
        margin-top: 20px;
    }
    .add_:hover {
        color: #409EFF;
        background-color: #ecf5fe;
    }
    .border_ {
        height: 1px;
        border-top: 1px solid #DCDFE6;
        margin-top: 10px;
    }
    .el-float {
        cursor: pointer;
        float: right;
    }
    .title_ {
        font-size: 17px;
        font-weight: 600;
    }
    .box_1 {
        padding-bottom: 15px;
        border-bottom: 1px solid #E4E7ED;
    }
    .box-content {
        min-height: 600px;
    }
    .box-content-first {
        height: 600px;
        overflow-y: auto;
        border-right: 1px solid #e4e7ed;
    }
    .box-content-second {
        height: 600px;
        overflow-y: auto;
    }
    .box-btn {
        margin-top: 15px;
        padding-top: 15px;
        border-top: 1px solid #E4E7ED;
    }
    .btn-float {
        float: right;
    }
    .mar_top {
        margin-top: 15px;
    }
    .align {
        text-align: center;
    }

    .custom-tree-node {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: 14px;
        padding-right: 8px;
    }
    .mar_right {
        margin-right: 5px;
    }

    .hide_ {
        display: none;
    }
    .jus-tree-node__content:hover .hide_ {
        display: inline-block;
    }

    .del_ {
        margin-left: 10px;
        display: none;
    }
    .del_:hover {
        cursor: pointer;
        color: #F56C6C;
    }
    .m_box {
        border-left: 1px solid #e4e7ed;
        border-right: 1px solid #e4e7ed;
        border-top: 1px solid #e4e7ed;
    }
    .m_box:hover .del_ {
        display: inline-block;
    }
    .m_box:hover {
        background-color: #409eff2b;
        border-radius: 3px;
    }
    .m_box span:first-child {
        display: inline-block;
        margin-left: 10px;
    }
    .method_box {
        margin-top: 5px;
        border-radius: 3px;
    }
    .m_box:not(:first-child) {
    }
    .m_box:last-child {
        border-bottom: 1px solid #e4e7ed;
    }
</style>
<style>
    .icon_box .el-input-group__prepend {
        padding: 0 !important;
    }
</style>
