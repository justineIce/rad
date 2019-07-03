<!-- 角色管理 -->
<template>
    <d2-container class="meta-public">
        <JusLayout aside :aside-width="200" title="角色管理">

            <!-- 侧面栏 -->
            <template slot="aside">

                <div class="sub-title">角色</div>
                <!-- 菜单 -->
                <div>
                    <!--角色列表-->
                    <jus-tree style="background-color: transparent;"
                              :data="roleOption"
                              :props="props"
                              defaultExpandAll
                              icon-class="fa fa-user-o"
                              :contextmenu="contextmenu"
                              @handleEditCatalog="handleEditCatalog"
                              @node-click="nodeClick"
                              @handleRemoveCatalog="handleRemoveCatalog"></jus-tree>
                </div>
                <div class="border_"></div>
                <div class="add_" @click="addRole">
                    <span class="el-icon-plus pad"> 添加新角色</span>
                </div>
            </template>

            <!--正文部分-->
            <div class="box-content">
                <el-row :gutter="10">
                    <el-col :span="16">
                        <div class="box-content-first">
                            <el-tree
                                    ref="tree"
                                    :data="resourceOption"
                                    show-checkbox
                                    node-key="id"
                                    default-expand-all
                                    :default-checked-keys="roleForm.checkedList"
                                    @node-click="nodeClick1"
                                    @check-change="checkChange"
                                    @check="checkedNodes"
                                    :props="defaultProps">
                            </el-tree>
                        </div>
                    </el-col>
                    <el-col :span="8">
                        <div class="box_1">
                            <span class="title_">操作权限</span>
                        </div>
                        <div>
                            <!--<el-checkbox-group v-model="btnCheckList" @change="handleCheckedCitiesChange">-->
                                <!--<div style="margin-top:10px;" v-for="item in btnTree" :key="item.id">-->
                                    <!--<el-checkbox :label="item.id" :value="item.id">{{item.name}}</el-checkbox>-->
                                <!--</div>-->
                            <!--</el-checkbox-group>-->
                            <el-tree
                                    ref="tree1"
                                    :data="btnTree"
                                    show-checkbox
                                    node-key="id"
                                    default-expand-all
                                    :default-checked-keys="btnCheckList"
                                    @check-change="handleCheckedCitiesChange"
                                    :props="props">
                            </el-tree>
                        </div>
                    </el-col>
                </el-row>
            </div>
            <div class="box-btn">
                <span class="btn-float">
                    <el-button type="primary" @click="remainAll">保存权限</el-button>
                </span>
            </div>

            <!--添加新角色-->
            <el-dialog :title="title" :visible.sync="roleDialogVisible">
                <el-form :model="roleForm" ref="roleForm" :rules="roleFormRules" :label-width="labelWidth">
                    <el-form-item label="角色名称：" prop="name">
                        <el-input v-model="roleForm.name"></el-input>
                    </el-form-item>
                    <el-form-item label="角色描述：" prop="remarks">
                        <el-input v-model="roleForm.remarks"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submit">保存</el-button>
                        <el-button @click="roleDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </JusLayout>
    </d2-container>
</template>
<script>
import { GetRoleList, AddRole, DelRole, GetMenuList, GetCheckedMenuBtnLimit, GetBtnMenuList, AddRoleResource, GetBtnMenuListByIds } from '@/api/platform.manage'
export default {
  name: 'roleManage',
  data () {
    return {
      // 角色
      roleOption: [],
      props: {
        label: 'name',
        children: 'Children'
      },
      defaultProps: {
        label: 'name',
        icon: 'shortcut_icon',
        children: 'Children'
      },
      contextmenu: [
        // { text: '添加子目录', icon: 'jus-icon-file', emit: 'handleAddCatalog' },
        { text: '编辑', icon: 'jus-icon-data', emit: 'handleEditCatalog' },
        { text: '删除', icon: 'jus-icon-delete', emit: 'handleRemoveCatalog' }
      ],

      // 添加角色
      title: '添加新角色',
      roleDialogVisible: false,
      labelWidth: '120px',
      roleForm: {
        id: '',
        name: '',
        remarks: '',
        checkedList: []
      },
      roleFormRules: {
        name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }]
      },
      resourceOption: [],

      btnCheckList: [],
      btnTree: [],

      // 用于保存当前树型菜单选中状态
      checkedStatus: false,

      // 用于保存角色，资源，操作权限相对应的数据
      form: {
        roleId: '',
        resourceId: [], // 用于保存选择的菜单资源
        btnId: [] // 用于保存选择的按钮资源
      },
      // 用于保存当前选中的资源id
      currentResourceId: '',
      // 用于保存当前的角色id
      currentRoleId: '',
      // 用于保存当前选择(选中/取消)的最底层菜单资源下的所有菜单的id
      currentResourceIdArr: []
    }
  },
  methods: {
    // 编辑角色
    handleEditCatalog (node, data) {
      this.roleDialogVisible = true
      this.roleForm = {
        id: data.id,
        remarks: data.remarks,
        name: data.name,
        checkedList: []
      }
    },
    // 删除角色
    handleRemoveCatalog (node, data) {
      this.$confirm('确定删除此角色吗？', '删除提醒', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 具体操作
        DelRole({ id: data.id }).then((res) => {
          if (res.ret === 200) {
            this.$message.success('删除成功')
            this.getRoleList()
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
      if (this.currentRoleId != data.id) {
        this.currentRoleId = data.id
        this.form = {
          roleId: data.id,
          resourceId: [],
          btnId: []
        }
        this.getMenuList()
        this.btnTree = []
      }
    },
    // 添加新角色
    addRole () {
      this.roleDialogVisible = true
      this.roleForm = {
        id: '',
        name: '',
        remarks: '',
        checkedList: []
      }
    },

    // 正文内容菜单点击
    nodeClick1 (data, node, vim) {
      if (node.checked && (data.Children == null || data.Children == undefined || data.Children.length <= 0)) {
        this.currentResourceId = data.id
        // 根据资源id，从后台获取资源对应的按钮操作权限
        this.getBtnResourceList()
      } else {
        this.btnTree = []
      }
    },

    // 保存选项框选中的状态
    checkChange (data, checked, vim) {
      this.checkedStatus = checked
    },

    // 点击选项框时，当前点击所对应的数据和节点
    checkedNodes (data, nodes) {
      this.currentResourceIdArr = []
      if (data.Children == null || data.Children == undefined || data.Children.length <= 0) {
        if (this.checkedStatus) {
          this.currentResourceId = data.id
          // 根据资源id，从后台获取资源对应的按钮操作权限
          this.getBtnResourceList()
        } else {
          this.btnTree = []
        }
      } else {
        this.btnTree = []
      }

      this.form.resourceId = this.$refs.tree.getCheckedKeys() // 菜单资源权限
      if (data.Children && data.Children.length > 0) {
        this.getLastData([data])
      } else {
        this.operation(data.id)
      }

      // 获取当前选择或者取消的菜单先的子按钮资源，如果子按钮资源中的id在 btnCheckList 中存在，取消时，删除对应数据,并更新到this.form.btnId中
      var param = {
        ids: this.currentResourceIdArr.join(',')
      }
      GetBtnMenuListByIds(param).then((res) => {
        if (res.ret === 200) {
          console.log(this.btnCheckList)
          var data = res.data
          for (var i in data) {
            if (this.inArray(data[i].id, this.btnCheckList)) {
              if (!this.checkedStatus) {
                for (var j in this.btnCheckList) {
                  if (this.btnCheckList[j] == data[i].id) {
                    this.btnCheckList.splice(j, 1)
                  }
                }
              }
            }
          }
        }
      })
    },

    // 根据当前选中菜单资源获取按钮资源
    getBtnResourceList () {
      GetBtnMenuList({ parentId: this.currentResourceId }).then((res) => {
        if (res.ret === 200) {
          this.btnTree = res.data
        }
      })
    },

    // 循环资源数据，找到最底层的数据
    getLastData (data) {
      for (var i in data) {
        if (data[i].Children && data[i].Children.length > 0) {
          this.getLastData(data[i].Children)
        } else {
          this.operation(data[i].id) // 这里只添加最底层的资源id
        }
      }
    },

    operation (id) {
      // 如果存在则不在添加
      if (!this.isExistResourceId(id)) {
        this.currentResourceIdArr.push(id)
      }
    },

    // 判断资源id是否已存在
    isExistResourceId (id) {
      for (var i in this.currentResourceIdArr) {
        if (this.currentResourceIdArr[i] == id) {
          return true
        }
      }
      return false
    },

    submit () {
      this.$refs.roleForm.validate((valid) => {
        if (valid) {
          AddRole(this.roleForm).then((res) => {
            if (res.ret === 200) {
              this.$message.success('保存成功')
              this.roleDialogVisible = false
              this.getRoleList()
            }
          })
        }
      })
    },

    // 保存权限
    remainAll () {
      console.log(this.form)
      if (this.form.roleId.length <= 0) {
        this.$message.warning('请先创建新角色')
        return
      }
      var param = {
        roleId: this.form.roleId,
        resourceId: this.form.resourceId.join(','),
        btnId: this.form.btnId.join(',')
      }
      console.log(param)
      AddRoleResource(param).then((res) => {
        if (res.ret === 200) {
          this.$message.success('保存成功')
        }
      })
    },

    // 操作权限
    handleCheckedCitiesChange (data, checked, vim) {
      if (checked) { // 不存在时添加
        this.btnCheckList.push(data.id)
      } else { // 删除
        for (var i in this.btnCheckList) {
          if (this.btnCheckList[i] == data.id) {
            this.btnCheckList.splice(i, 1)
          }
        }
      }
    },

    // 判断元素在数组中是否存在
    inArray (e, arr) {
      for (var i in arr) {
        if (arr[i] == e) {
          return true
        }
      }
      return false
    },

    // 获取角色列表
    getRoleList () {
      GetRoleList().then((res) => {
        if (res.ret === 200) {
          this.roleOption = res.data

          if (res.data.length > 0) {
            var firstRole = res.data[0]
            this.currentRoleId = firstRole.id
            this.form.roleId = this.currentRoleId
          }

          this.getMenuList()
        }
      })
    },
    // 获取所有的菜单列表
    getMenuList () {
      GetMenuList().then((res) => {
        if (res.ret === 200) {
          this.resourceOption = res.data
          if (this.roleOption.length > 0) {
            var param = { roleId: this.currentRoleId, type: 'menu' }
            var btnParam = { roleId: this.currentRoleId, type: 'button' }
            this.getCheckedMenuLimit(param)
            this.getCheckedMenuLimit(btnParam)
          }
        }
      })
    },

    // 获取已拥有的菜单权限
    getCheckedMenuLimit (param) {
      GetCheckedMenuBtnLimit(param).then((res) => {
        if (res.ret === 200) {
          var data = res.data
          var arr = []

          if (data.length > 0) {
            data = data.substr(0, res.data.length - 1).split(',')
            for (var i in data) {
              arr.push(Number(data[i]))
            }
          }

          if (param.type === 'menu') {
            this.$refs.tree.setCheckedKeys(arr) // 渲染el-tree树，直接设置默认值时，清空无效
            this.form.resourceId = arr //
          } else if (param.type === 'button') {
            this.btnCheckList = arr
            this.form.btnId = arr
          }
        }
      })
    }
  },
  mounted () {
    this.getRoleList()
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
    .title_ {
        font-size: 17px;
        font-weight: 600;
    }
    .box_1 {
        padding-bottom: 15px;
        border-bottom: 1px solid #E4E7ED;
    }
    .border_ {
        height: 1px;
        border-top: 1px solid #DCDFE6;
        margin-top: 10px;
    }
    .box-content {
        height: 600px;
    }
    .box-content-first {
        height: 600px;
        overflow-y: auto;
        border-right: 1px solid #e4e7ed;
    }
    .box-btn {
        margin-top: 15px;
        padding-top: 15px;
        border-top: 1px solid #E4E7ED;
    }
    .btn-float {
        float: right;
    }
</style>
