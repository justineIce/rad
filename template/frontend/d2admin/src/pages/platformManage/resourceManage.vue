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
                              @node-click="nodeClick"
                              :expand-on-click-node="false">
                        <span class="custom-tree-node" slot-scope="{ node, data }">
                            <span>{{ node.label }}</span>
                            <span class="hide_">
                                <span class="el-icon-edit mar_right"
                                      @click="() => handleEditCatalog(node, data)"></span>
                                <span class="el-icon-delete" style="color: #F56C6C;"
                                      @click="() => handleRemoveCatalog(node, data)"></span>
                            </span>
                        </span>
                    </jus-tree>
                </div>
                <div class="border_"></div>
                <div class="add_" @click="SaveMenu" v-if="canAdd">
                    <span class="el-icon-plus pad"> 添加新菜单 {{canAdd}}</span>
                </div>
            </template>

            <!--正文部分-->
            <div>
                <el-row :gutter="10">
                    <el-col :span="12">
                        <div>
                            <jus-tree
                                    :data="allMenuOption"
                                    :props="props"
                                    @node-click="nodeClick1"
                                    defaultExpandAll></jus-tree>
                        </div>
                    </el-col>
                    <el-col :span="12" style="border-left: 1px solid #f3f3f3;">
                        <div>
                            <div class="box_1">
                                <span class="title_">{{checkedMenu}}</span>
                                <span class="btn-float">
                                    <el-button type="primary" @click="builtChildMenu">
                                            <span class="el-icon-plus"></span>
                                        新增子菜单
                                    </el-button>
                                    <el-button type="success" @click="remainAll">
                                        <span class="el-icon-edit"></span>
                                        保存菜单
                                    </el-button>
                                    <el-button type="danger" @click="delMenu">
                                            <span class="el-icon-delete"></span>
                                        删除菜单
                                    </el-button>
                                </span>
                            </div>
                            <div class="mar_top">
                                <el-form :label-width="labelWidth" :model="menuForm1" ref="menuForm1"
                                         :rules="menuForm1Rules">
                                    <el-form-item label="图标：" prop="Icon">
                                        <el-input class="icon_box" v-model="menuForm1.Icon">
                                            <template slot="prepend">
                                                <el-icon-popover :icon="menuForm1.Icon"
                                                                 @getIcon="getIcon"></el-icon-popover>
                                            </template>
                                        </el-input>
                                    </el-form-item>
                                    <el-form-item label="资源名称：" prop="Name">
                                        <el-input v-model="menuForm1.Name"></el-input>
                                    </el-form-item>
                                    <el-form-item label="资源路径：" prop="Href">
                                        <el-input v-model="menuForm1.Href"></el-input>
                                    </el-form-item>
                                    <el-form-item label="目标地址：" prop="Target">
                                        <el-input v-model="menuForm1.Target"></el-input>
                                    </el-form-item>
                                    <el-form-item label="排序号：">
                                        <el-input placeholder="用于主菜单排序" v-model="menuForm1.Sort"></el-input>
                                    </el-form-item>
                                    <el-form-item label="资源描述：">
                                        <el-input v-model="menuForm1.Remarks" type="textarea" :rows="2"></el-input>
                                    </el-form-item>
                                </el-form>
                            </div>
                            <div>
                                <div class="title_ box_1 mar_top">操作权限</div>
                                <div class="mar_top">
                                    <div style="margin-bottom: 10px;text-align: right">
                                        <el-button type="primary" @click="addOperateLimit">
                                            <span class="el-icon-plus"></span>
                                            添加操作权限
                                        </el-button>
                                    </div>
                                    <JusTable ref="table"
                                              :columns="columns"
                                              :data="data"
                                              :row-handle="rowHandle"
                                              @row-remove="rowRemove"
                                              @editRow="editRow">
                                    </JusTable>
                                </div>
                            </div>

                        </div>
                    </el-col>
                </el-row>
            </div>

            <!--添加新菜单-->
            <el-dialog :title="title" :visible.sync="menuDialogVisible">
                <el-form :model="menuForm" ref="menuForm" :rules="menuFormRules" :label-width="labelWidth">
                    <el-form-item label="父级资源名称：" v-if="menuForm.parentName">
                        <el-input disabled v-model="menuForm.parentName"></el-input>
                    </el-form-item>
                    <el-form-item label="图标：">
                        <el-input class="icon_box" v-model="menuForm.Icon">
                            <template slot="prepend">
                                <el-icon-popover :icon="menuForm.Icon" @getIcon="getMenuIcon"></el-icon-popover>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item label="资源名称：" prop="Name">
                        <el-input v-model="menuForm.Name"></el-input>
                    </el-form-item>
                    <el-form-item label="资源路径：" prop="Href">
                        <el-input v-model="menuForm.Href"></el-input>
                    </el-form-item>
                    <el-form-item label="目标地址：" prop="Target">
                        <el-input v-model="menuForm.Target"></el-input>
                    </el-form-item>
                    <el-form-item label="排序号：">
                        <el-input placeholder="用于主菜单排序" v-model="menuForm.Sort"></el-input>
                    </el-form-item>
                    <el-form-item label="资源描述：">
                        <el-input v-model="menuForm.Remarks" type="textarea" :rows="3"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="submit">保存</el-button>
                        <el-button @click="menuDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>

            <!--添加/编辑操作权限-->
            <el-dialog :title="btnTitle" :visible.sync="addOperateLimitDialogVisible">
                <el-form :label-width="labelWidth" :model="operateLimit" ref="operateLimit" :rules="operateLimitRules">
                    <el-form-item label="所属菜单：">
                        <el-input v-model="operateLimit.SysMenuName" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="权限名称：" prop="Name">
                        <el-dropdown trigger="click" @command="handleMenuBtnCommand">
                            <el-input v-model="operateLimit.Name" placeholder="请选择或请输入权限名称"></el-input>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item icon="el-icon-plus" :command="{value:'insert',label:'添加'}">添加
                                </el-dropdown-item>
                                <el-dropdown-item icon="el-icon-delete" :command="{value:'delete',label:'删除'}">删除
                                </el-dropdown-item>
                                <el-dropdown-item icon="el-icon-search" :command="{value:'select',label:'查询'}">查询
                                </el-dropdown-item>
                                <el-dropdown-item icon="el-icon-edit" :command="{value:'update',label:'修改'}">修改
                                </el-dropdown-item>
                                <el-dropdown-item icon="el-icon-info" :command="{value:'info',label:'详情'}">详情
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-form-item>
                    <el-form-item label="权限标识：" prop="mark">
                        <el-input v-model="operateLimit.Permission" placeholder="请输入权限标识且必须唯一"></el-input>
                    </el-form-item>
                    <el-form-item label="请求方式：" prop="Method">
                        <el-select
                                v-model="operateLimit.Method"
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
                    <el-form-item label="权限Url：">
                        <el-input v-model="limitMethod" class="input-with-select" placeholder="请输入权限Url并点击添加">
                            <el-button slot="append" type="success" @click="addLimitMethod">添加</el-button>
                        </el-input>
                        <div class="method_box">
                            <el-tag size="mini" closable v-for="(item,index) in limitMethodArr"
                                    @close="delLimitMethod(index)" :key="index">{{item}}
                            </el-tag>
                        </div>
                    </el-form-item>
                    <el-form-item label="权限描述：">
                        <el-input v-model="operateLimit.Remarks" placeholder="对权限按钮的描述"></el-input>
                    </el-form-item>
                    <el-form-item class="text-right">
                        <el-button @click="submitLimit">保存</el-button>
                        <el-button @click="addOperateLimitDialogVisible=false">取消</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>

        </JusLayout>
    </d2-container>
</template>
<script>
    import {DelMenu, DelMenuBtn, GetBtnMenuList, GetMenuList, SaveMenu, SaveMenuBtn} from '@/api/platform.manage'

    export default {
        name: 'resourceMange',
        data() {
            return {
                canAdd: this.hasPermissions(["p_menu_edit"]),
                // 顶级菜单
                menuOption: [],
                props: {
                    label: 'Name',
                    icon: 'Icon',
                    children: 'Children'
                },

                // 添加新菜单
                title: '添加新菜单',
                menuDialogVisible: false,
                labelWidth: '120px',
                menuForm: {
                    ID: '',
                    Icon: '',
                    Name: '',
                    Href: '',
                    Target: '',
                    Remarks: '',
                    Sort: '0'
                },
                // 资源类型
                menuTypeOption: [
                    {value: 'menu', label: 'menu'}
                ],
                menuFormRules: {
                    Name: [{required: true, message: '请输入菜单名称', trigger: 'blur'}],
                    Href: [{required: true, message: '请输入资源路径', trigger: 'blur'}]
                },
                // 根据左侧菜单选择生成
                allMenuOption: [],
                // 顶级菜单 - 隐藏部分元素
                firstClass: true,

                // 当前被选中的数据 - 编辑删除
                checkedMenu: '',
                menuForm1: {
                    ID: '',
                    ParentId: '',
                    Icon: '',
                    Name: '',
                    Href: '',
                    Target: '',
                    Remarks: '',
                    Sort: '0'
                },
                menuForm1Rules: {
                    Name: [{required: true, message: '请输入名称', trigger: 'blur'}],
                    Icon: [{required: true, message: '请输入图标', trigger: 'blur'}],
                    Href: [{required: true, message: '请输入资源路径', trigger: 'blur'}],
                    Target: [{required: true, message: '请输入目标地址', trigger: 'blur'}]
                },
                method: '',
                MethodArr: [],
                addOperateLimitDialogVisible: false,
                // 添加操作权限
                btnTitle: '添加操作权限',
                operateLimit: {
                    ID: '',
                    SysMenuId: '',
                    SysMenuName: '',
                    Name: '',
                    Remarks: '',
                    Method: '',
                    Permission: '',
                    Path: ''
                },
                btnTypeOption: [
                    {value: 'button', label: 'button'}
                ],
                operateLimitRules: {
                    Name: [{required: true, message: '请输入操作名称', trigger: 'blur'}],
                    SysMenuId: [{required: true, message: '请选择菜单', trigger: 'blur'}],
                    Permission: [{required: true, message: '请输入权限标识且必须唯一', trigger: 'blur'}],
                    Method: [{required: true, message: '请选择请求方式', trigger: 'blur'}],
                    Path: [{required: true, message: '请添加权限Url', trigger: 'blur'}]
                },
                limitMethod: '',
                limitMethodArr: [],
                columns: [
                    {title: '名称', key: 'Name'},
                    {title: '请求方法', key: 'Method'},
                    {title: '权限标识', key: 'Permission'},
                    {title: '描述', key: 'Remarks'}
                ],
                data: [],
                rowHandle: {
                    remove: {text: '', icon: 'el-icon-remove', emit: 'rowRemove'},
                    custom: [
                        {text: '', icon: 'el-icon-edit', emit: 'editRow'}
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
                    {icon: 'el-icon-info'},
                    {icon: 'el-icon-error'},
                    {icon: 'el-icon-success'},
                    {icon: 'el-icon-warning'},
                    {icon: 'el-icon-question'},
                    {icon: 'el-icon-back'}
                ],

                baseBtnOption: [
                    {value: 'POST', label: 'POST'},
                    {value: 'GET', label: 'GET'}
                ]
            }
        },
        methods: {
            // 编辑菜单
            handleEditCatalog(node, data) {
                console.log(node, data)
                this.title = '编辑菜单'
                this.menuDialogVisible = true
                this.menuForm = {
                    ID: data.ID,
                    ParentId: data.ParentId,
                    Icon: data.Icon,
                    Href: data.Href,
                    Name: data.Name,
                    Remarks: data.Remarks,
                    Target: data.Target,
                    Sort: data.Sort
                }
            },
            // 删除菜单
            handleRemoveCatalog(node, data) {
                this.$confirm('确定删除此菜单吗？', '删除提醒', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    // 具体操作
                    DelMenu({id: data.ID}).then((res) => {
                        if (res.ret === 200) {
                            this.$message.success('删除成功')
                            this.currentTopMenuId = ''
                            this.getMenuList(0)
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
            nodeClick(data, node, vim) {
                this.currentTopMenuId = data.ID
                this.currentTopMenuId = data.ID
                // 从后台动态获取
                this.getMenuList(data.ID)
                this.checkedMenu = data.Name
                this.menuForm1 = {
                    ID: data.ID,
                    ParentId: data.ParentID,
                    Name: data.Name,
                    Href: data.Href,
                    Target: data.Target,
                    Icon: data.Icon,
                    Sort: data.Sort,
                    Method: data.Method,
                    Remarks: data.Remarks
                }
                this.operateLimit.SysMenuId = data.ID
                this.operateLimit.SysMenuName = data.Name
                this.getBtnMenuList();
            },

            // 添加新菜单
            SaveMenu() {
                this.menuDialogVisible = true
                this.title = '添加菜单'
                this.menuForm = {
                    ID: '',
                    ParentId: '0',
                    Icon: '',
                    Name: '',
                    Remarks: '',
                    Href: '',
                    Target: '',
                    Sort: '0'
                }
            },
            // 新增子菜单
            builtChildMenu() {
                this.menuDialogVisible = true
                this.title = '添加子菜单'
                this.menuForm = {
                    ID: '',
                    ParentId: this.menuForm1.ID,
                    parentName: this.menuForm1.Name,
                    Icon: '',
                    Name: '',
                    Remarks: '',
                    Href: '',
                    Target: '',
                    Sort: '0'
                }
            },
            // 保存
            submit() {
                this.$refs.menuForm.validate((valid) => {
                    if (valid) {
                        // 添加/编辑新菜单
                        SaveMenu(this.menuForm).then((res) => {
                            console.log(res)
                            if (res.ret === 200) {
                                this.menuDialogVisible = false
                                this.$message.success('保存成功')
                                this.getMenuList(0)
                                if (this.currentTopMenuId != '') {
                                    this.getMenuList(this.currentTopMenuId)
                                }
                            }
                        })
                    }
                })
            },
            getMenuIcon(icon) {
                this.menuForm.Icon = icon
            },

            // 正文内容菜单点击
            nodeClick1(data, node, vim) {
                this.checkedMenu = data.Name
                this.menuForm1 = {
                    ID: data.ID,
                    ParentId: data.ParentID,
                    Name: data.Name,
                    Icon: data.Icon,
                    Href: data.Href,
                    Target: data.Target,
                    Method: data.Method,
                    Remarks: data.Remarks,
                    Sort: data.Sort
                }
                this.operateLimit.SysMenuId = data.ID
                this.operateLimit.SysMenuName = data.Name
                this.getBtnMenuList();
            },

            // 全部保存
            remainAll() {
                this.$refs.menuForm1.validate((valid) => {
                    if (valid) {
                        SaveMenu(this.menuForm1).then((res) => {
                            console.log(res)
                            if (res.ret === 200) {
                                this.$message.success('保存成功')
                                this.getMenuList(0)
                                if (this.currentTopMenuId != '') {
                                    this.getMenuList(this.currentTopMenuId)
                                }
                            }
                        })
                    }
                })
            },
            // 删除菜单
            delMenu() {
                this.$confirm('确定删该菜单吗？', '删除提醒', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    // 具体操作
                    DelMenu({id: this.menuForm1.ID}).then((res) => {
                        if (res.ret === 200) {
                            this.$message.success('删除成功')
                            this.getMenuList(this.currentTopMenuId)
                        }
                    })
                }).catch(() => {
                    this.$message.warning('取消删除')
                })
            },

            // 添加操作权限
            addOperateLimit() {
                this.btnTitle = '添加操作权限'
                this.operateLimit.ID = ''
                this.operateLimit.Name = ''
                this.operateLimit.Remarks = ''
                this.operateLimit.Method = ''
                this.operateLimit.Permission = ''
                this.operateLimit.Path = ''
                this.addOperateLimitDialogVisible = true
            },
            // 保存操作权限
            submitLimit() {
                this.$refs.operateLimit.validate((valid) => {
                    if (valid) {
                        if (this.operateLimit.Method.length <= 0) {
                            this.$message.warning('请输入请求方法')
                        } else {
                            SaveMenuBtn(this.operateLimit).then((res) => {
                                console.log(res)
                                if (res.ret === 200) {
                                    this.addOperateLimitDialogVisible = false
                                    this.$message.success('保存成功')
                                    this.getBtnMenuList()
                                }
                            })
                        }
                    }
                })
            },
            handleMenuBtnCommand(command) {
                this.operateLimit.Name = command.label
                this.operateLimit.Permission = command.value
            },
            // 添加请求方法
            addLimitMethod() {
                if (this.limitMethod.trim().length <= 0) {
                    this.$message.warning('请输入请求方法')
                    return
                }
                this.limitMethodArr.push(this.limitMethod)
                this.limitMethod = ''
                this.operateLimit.Path = this.limitMethodArr.join(',')
            },
            // 删除请求方法
            delLimitMethod(index) {
                for (var i in this.limitMethodArr) {
                    if (index == i) {
                        this.limitMethodArr.splice(index, 1)
                    }
                }
                this.operateLimit.Path = this.limitMethodArr.join(',')
            },

            // 删除操作权限
            rowRemove({index, row}) {
                DelMenuBtn({id: row.ID}).then((res) => {
                    if (res.ret === 200) {
                        this.$message.success('删除成功')
                        this.getBtnMenuList()
                    }
                })
            },
            // 编辑操作权限
            editRow({index, row}) {
                this.addOperateLimitDialogVisible = true
                this.btnTitle = '编辑操作权限'
                this.operateLimit = {
                    ID: row.ID,
                    SysMenuId: this.menuForm1.ID,
                    SysMenuName: this.menuForm1.Name,
                    Target: this.menuForm1.Target,
                    Name: row.Name,
                    Icon: row.Icon,
                    Method: row.Method,
                    Permission: row.Permission,
                    Remarks: row.Remarks
                }
                if (row.Method.length > 0) {
                    this.limitMethodArr = row.Path.split(',')
                } else {
                    this.limitMethodArr = []
                }
            },
            getBtnMenuList() {
                // 获取所具体菜单下的所有按钮操作权限
                GetBtnMenuList({
                    menuId: this.menuForm1.ID
                }).then((res) => {
                    if (res.ret === 200) {
                        var data = res.data
                        this.data = data
                    }
                })
            },
            // 获取顶部菜单
            getMenuList(parentId) {
                var that=this
                GetMenuList({parentId: parentId}).then((res) => {
                    if (res.ret === 200) {
                        if (parentId === 0) {
                            that.menuOption = res.data
                            if (that.currentTopMenuId === '') {
                                var firtMenu = that.menuOption[0]
                                that.getMenuList(firtMenu.ID)
                                that.checkedMenu = firtMenu.Name
                                that.currentTopMenuId = firtMenu.ID
                                that.menuForm1 = {
                                    ID: firtMenu.ID,
                                    ParentId: firtMenu.ParentID,
                                    Name: firtMenu.Name,
                                    Icon: firtMenu.Icon,
                                    Href: firtMenu.Href,
                                    Target: firtMenu.Target,
                                    Method: firtMenu.Method,
                                    Remarks: firtMenu.Remarks,
                                    Sort: firtMenu.Sort
                                }
                                if (firtMenu.Method > 0) {
                                    that.MethodArr = firtMenu.Method.split(',')
                                }
                                that.operateLimit.SysMenuId = firtMenu.ID
                                that.operateLimit.SysMenuName = firtMenu.Name
                                that.getBtnMenuList();
                            }
                        } else {
                            that.allMenuOption = res.data
                        }
                    }
                })
            },
            // 图标选择
            getIcon(icon) {
                this.menuForm1.Icon = icon
            }
        },
        mounted() {
            this.getMenuList(0)
        }
    }
</script>
<style scoped>
    .el-tag + .el-tag {
        margin-left: 10px;
    }

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

    .Method_box {
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

    .is-current {
        background-color: #ffffff !important;
    }
</style>
