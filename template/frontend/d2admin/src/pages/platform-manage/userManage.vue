<!-- 用户管理 -->
<template>
    <d2-container class="meta-public">
        <JusLayout title="用户管理">
            <div class="mar_top">
                <el-row :gutter="20">
                    <el-form :model="searchForm" :label-width="labelWidth" label-position="left">
                        <el-col :span="4">
                            <el-form-item label="姓名：">
                                <el-input v-model="searchForm.name"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="4">
                            <el-form-item label="角色：">
                                <el-input v-model="searchForm.roleName"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="创建时间：" label-width="80px">
                                <div class="block">
                                    <el-date-picker
                                            v-model="searchForm.time"
                                            size="mini"
                                            type="daterange"
                                            @change="getValue"
                                            range-separator="至"
                                            start-placeholder="开始日期"
                                            end-placeholder="结束日期">
                                    </el-date-picker>
                                </div>
                            </el-form-item>
                        </el-col>
                        <el-col :span="4">
                            <el-button @click="searchDataList">查询</el-button>
                            <el-button @click="reSearchDataList">重置</el-button>
                        </el-col>
                    </el-form>
                </el-row>
            </div>

            <!--数据列表-->
            <JusTable  class=""
                       :options="{'row-key':'id'}"
                       :columns="column"
                       :data="data"
                       :rowHandle="rowHandle"
                       @reset="reset"
                       @changeRole="changeRole"
                       @del="del"
                       :showPagination="true"
                       @query-data="queryData"
                       @row-data-change="handleRowDataChange">
                <template slot="header">
                    <el-button class="el-icon-circle-plus" type="primary" size="small" @click="buildUser"> 创建用户</el-button>
                </template>
                <template slot="FullName" slot-scope="scope">
                    <img class="header_img" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHN2ZyB3aWR0aD0iMTk5cHgiIGhlaWdodD0iMTk5cHgiIHZpZXdCb3g9IjAgMCAxOTkgMTk5IiB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPgogICAgPCEtLSBHZW5lcmF0b3I6IFNrZXRjaCA0Ny4xICg0NTQyMikgLSBodHRwOi8vd3d3LmJvaGVtaWFuY29kaW5nLmNvbS9za2V0Y2ggLS0+CiAgICA8dGl0bGU+R3JvdXAgMjA8L3RpdGxlPgogICAgPGRlc2M+Q3JlYXRlZCB3aXRoIFNrZXRjaC48L2Rlc2M+CiAgICA8ZGVmcz48L2RlZnM+CiAgICA8ZyBpZD0i5Liq5Lq65Lit5b+DIiBzdHJva2U9Im5vbmUiIHN0cm9rZS13aWR0aD0iMSIgZmlsbD0ibm9uZSIgZmlsbC1ydWxlPSJldmVub2RkIj4KICAgICAgICA8ZyBpZD0iR3JvdXAtMjAiPgogICAgICAgICAgICA8cmVjdCBpZD0iUmVjdGFuZ2xlLTkiIGZpbGw9IiNBNkM0RDAiIHg9IjAiIHk9IjAiIHdpZHRoPSIxOTkiIGhlaWdodD0iMTk5IiByeD0iMTYiPjwvcmVjdD4KICAgICAgICAgICAgPGcgaWQ9ImJveS01IiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgzOC4wMDAwMDAsIDIyLjAwMDAwMCkiPgogICAgICAgICAgICAgICAgPHBvbHlnb24gaWQ9IkZpbGwtMSIgZmlsbD0iI0ZEQ0M5QiIgcG9pbnRzPSI1MSAxNjEgNzcgMTYxIDc3IDEzMiA1MSAxMzIiPjwvcG9seWdvbj4KICAgICAgICAgICAgICAgIDxwYXRoIGQ9Ik01MSwxMzUuNDIzMDQzIEM1MSwxMzUuNDIzMDQzIDYxLjkzOTU3MDUsMTQyLjM2NTAzOSA3NywxNDAuNzU4OCBMNzcsMTMyIEw1MSwxMzIgTDUxLDEzNS40MjMwNDMgWiIgaWQ9IkZpbGwtMiIgZmlsbD0iI0ZDQkM4NSI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTI2LDg2LjUwMjIxNTEgQzI2LDk0LjUwNTM0NjggMjAuMTgwMDUxNCwxMDEgMTIuOTk4MzAwNiwxMDEgQzUuODE5OTQ4NTgsMTAxIDAsOTQuNTA1MzQ2OCAwLDg2LjUwMjIxNTEgQzAsNzguNDkzNTQ1NyA1LjgxOTk0ODU4LDcyIDEyLjk5ODMwMDYsNzIgQzIwLjE4MDA1MTQsNzIgMjYsNzguNDkzNTQ1NyAyNiw4Ni41MDIyMTUxIEwyNiw4Ni41MDIyMTUxIFoiIGlkPSJGaWxsLTMiIGZpbGw9IiNGQ0JDODUiPjwvcGF0aD4KICAgICAgICAgICAgICAgIDxwYXRoIGQ9Ik0xMjcsODYuNTAyMjE1MSBDMTI3LDk0LjUwNTM0NjggMTIxLjE4MDE3LDEwMSAxMTQuMDAwNTY3LDEwMSBDMTA2LjgyMDk2MywxMDEgMTAxLDk0LjUwNTM0NjggMTAxLDg2LjUwMjIxNTEgQzEwMSw3OC40OTM1NDU3IDEwNi44MjA5NjMsNzIgMTE0LjAwMDU2Nyw3MiBDMTIxLjE4MDE3LDcyIDEyNyw3OC40OTM1NDU3IDEyNyw4Ni41MDIyMTUxIEwxMjcsODYuNTAyMjE1MSBaIiBpZD0iRmlsbC00IiBmaWxsPSIjRkNCQzg1Ij48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNMTEyLjk0ODcyMiw1OS4xMTIyMjY5IEMxMTIuOTQ4NzIyLDM3LjY4MjQ3NDkgOTcuMzQ1NTQ5LDIyIDYyLjQ5OTkzODMsMjIgQzI3LjY1MzIyMTcsMjIgMTIuMDUxMTU1LDM3LjY4MjQ3NDkgMTIuMDUxMTU1LDU5LjExMjIyNjkgQzEyLjA1MTE1NSw4MC41NDY0MjQ1IDguNDEyOTk1MDYsMTM3IDYyLjQ5OTkzODMsMTM3IEMxMTYuNTkwMTk5LDEzNyAxMTIuOTQ4NzIyLDgwLjU0NjQyNDUgMTEyLjk0ODcyMiw1OS4xMTIyMjY5IEwxMTIuOTQ4NzIyLDU5LjExMjIyNjkgWiIgaWQ9IkZpbGwtNSIgZmlsbD0iI0ZEQ0M5QiI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTQ0LDgzIEM0NCw4Ni4zMTM4NjMyIDQxLjUzODExMTEsODkgMzguNTAyNzgzNyw4OSBDMzUuNDY2MzQyNyw4OSAzMyw4Ni4zMTM4NjMyIDMzLDgzIEMzMyw3OS42ODYxMzY4IDM1LjQ2NjM0MjcsNzcgMzguNTAyNzgzNyw3NyBDNDEuNTM4MTExMSw3NyA0NCw3OS42ODYxMzY4IDQ0LDgzIEw0NCw4MyBaIiBpZD0iRmlsbC02IiBmaWxsPSIjM0IyNTE5Ij48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNMzgsODAuNDk5NDkyIEMzOCw4MS4zMjU0MzE4IDM3LjMyNTQzMTgsODIgMzYuNTAwNTA4LDgyIEMzNS42NzA1MDQ2LDgyIDM1LDgxLjMyNTQzMTggMzUsODAuNDk5NDkyIEMzNSw3OS42NzA1MDQ2IDM1LjY3MDUwNDYsNzkgMzYuNTAwNTA4LDc5IEMzNy4zMjU0MzE4LDc5IDM4LDc5LjY3MDUwNDYgMzgsODAuNDk5NDkyIEwzOCw4MC40OTk0OTIgWiIgaWQ9IkZpbGwtNyIgZmlsbD0iI0ZGRkZGRiI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTMwLjUzNjc5Miw2OC4zNjIzMDg1IEMzMy4xMzg1MDkzLDY5LjY5MjQyMDkgMzcuMzE2ODU2MSw2My44NDY5ODcxIDQ2Ljg4MDI1NjMsNjguOTMxNzA1OSBDNDguNjI1MTMzOSw2OS44NjA5NjI1IDQ5LjY2NTgyMDgsNjEgMzkuNDE0OTQzMyw2MSBDMzAuNTM2NzkyLDYxIDI4Ljk2NDYxOTQsNjcuNTYxNzM1OCAzMC41MzY3OTIsNjguMzYyMzA4NSBMMzAuNTM2NzkyLDY4LjM2MjMwODUgWiIgaWQ9IkZpbGwtOCIgZmlsbD0iIzUxMzYyQSI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTk1LDgzIEM5NSw4Ni4zMTM4NjMyIDkyLjUzNzExMzksODkgODkuNTAwNTU3LDg5IEM4Ni40NjI4ODYxLDg5IDg0LDg2LjMxMzg2MzIgODQsODMgQzg0LDc5LjY4NjEzNjggODYuNDYyODg2MSw3NyA4OS41MDA1NTcsNzcgQzkyLjUzNzExMzksNzcgOTUsNzkuNjg2MTM2OCA5NSw4MyBMOTUsODMgWiIgaWQ9IkZpbGwtOSIgZmlsbD0iIzNCMjUxOSI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTg4LDgwLjQ5OTQ5MiBDODgsODEuMzI1NDMxOCA4Ny4zMjk0OTU0LDgyIDg2LjQ5OTQ5Miw4MiBDODUuNjcwNTA0Niw4MiA4NSw4MS4zMjU0MzE4IDg1LDgwLjQ5OTQ5MiBDODUsNzkuNjcwNTA0NiA4NS42NzA1MDQ2LDc5IDg2LjQ5OTQ5Miw3OSBDODcuMzI5NDk1NCw3OSA4OCw3OS42NzA1MDQ2IDg4LDgwLjQ5OTQ5MiBMODgsODAuNDk5NDkyIFoiIGlkPSJGaWxsLTEwIiBmaWxsPSIjRkZGRkZGIj48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNOTYuNDYzNTI3NSw2OC4zNjIzMDg1IEM5My44NjAyNTE1LDY5LjY5MjQyMDkgODkuNjgxMTkxLDYzLjg0Njk4NzEgODAuMTE3MjcxOSw2OC45MzE3MDU5IEM3OC4zNzU0Mzk1LDY5Ljg2MDk2MjUgNzcuMzM1Njg5Miw2MSA4Ny41ODcyMDMxLDYxIEM5Ni40NjM1Mjc1LDYxIDk4LjAzNDg1NDIsNjcuNTYxNzM1OCA5Ni40NjM1Mjc1LDY4LjM2MjMwODUgTDk2LjQ2MzUyNzUsNjguMzYyMzA4NSBaIiBpZD0iRmlsbC0xMSIgZmlsbD0iIzUxMzYyQSI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTYzLjk5ODI3MzMsMTA0LjUzMjUyMyBDNTguMzgxOTc4NiwxMDQuNTMyNTIzIDU1LDEwMC42MjQzNzQgNTUsMTAyLjUxNzg2OSBDNTUsMTA0LjQwNjAwNCA1Ni42MjMwNzM1LDEwOCA2My45OTgyNzMzLDEwOCBDNzEuMzc2OTI2NSwxMDggNzMsMTA0LjQwNjAwNCA3MywxMDIuNTE3ODY5IEM3MywxMDAuNjI0Mzc0IDY5LjYxNDU2OCwxMDQuNTMyNTIzIDYzLjk5ODI3MzMsMTA0LjUzMjUyMyBMNjMuOTk4MjczMywxMDQuNTMyNTIzIFoiIGlkPSJGaWxsLTEyIiBmaWxsPSIjRkNCQzg1Ij48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNNjQsMTIxLjg0NTEzOCBDNjIuMTI2NTI2MSwxMjEuODQ1MTM4IDYxLDEyMC41NDI1NjEgNjEsMTIxLjE3MTY0IEM2MSwxMjEuODAwNzIgNjEuNTQxNjIwNCwxMjMgNjQsMTIzIEM2Ni40NjE3MDkyLDEyMyA2NywxMjEuODAwNzIgNjcsMTIxLjE3MTY0IEM2NywxMjAuNTQyNTYxIDY1Ljg3MjM2NCwxMjEuODQ1MTM4IDY0LDEyMS44NDUxMzggTDY0LDEyMS44NDUxMzggWiIgaWQ9IkZpbGwtMTMiIGZpbGw9IiNGQ0JDODUiPjwvcGF0aD4KICAgICAgICAgICAgICAgIDxwYXRoIGQ9Ik02NC4wMDA1NjQ4LDExNi40Mzc0NjIgQzU1LjI2MzYyNSwxMTYuNDM3NDYyIDUwLDExMy4zNDM4MzIgNTAsMTE0LjEyNzA2NyBDNTAsMTE0LjkxMzI1IDUyLjUyNzg5NTQsMTE4IDY0LjAwMDU2NDgsMTE4IEM3NS40NzY2MjI3LDExOCA3OCwxMTQuOTEzMjUgNzgsMTE0LjEyNzA2NyBDNzgsMTEzLjM0MzgzMiA3Mi43MzYzNzUsMTE2LjQzNzQ2MiA2NC4wMDA1NjQ4LDExNi40Mzc0NjIgTDY0LjAwMDU2NDgsMTE2LjQzNzQ2MiBaIiBpZD0iRmlsbC0xNCIgZmlsbD0iI0Y3OTQ1RSI+PC9wYXRoPgogICAgICAgICAgICAgICAgPHBhdGggZD0iTTY0LDE0NCBMNjQsMTc3IEwwLDE3NyBDMCwxNjMuNjQwOTc1IDMzLjcxMDg1NTYsMTQ0IDY0LDE0NCBMNjQsMTQ0IFoiIGlkPSJGaWxsLTE1IiBmaWxsPSIjRjc5NDFFIj48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNNjQsMTQ0IEw2NCwxNzcgTDEyOCwxNzcgQzEyOCwxNjMuNjQwOTc1IDk0LjI4OTY3NDUsMTQ0IDY0LDE0NCBMNjQsMTQ0IFoiIGlkPSJGaWxsLTE2IiBmaWxsPSIjRjc5NDFFIj48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNNjMuMTU5OTA5NCwxNDQgQzU2LjgxMDgxNTQsMTQ0IDUwLjMwNjAzMDYsMTQ0Ljg1NTcyMSA0NCwxNDYuMzE3MzU2IEM0NC4wNDMwNjM0LDE0Ni4zNjYyNTQgNTAuOTk2MTQ5NSwxNTUgNjMuMTU5OTA5NCwxNTUgQzc0LjQyNTk2MjYsMTU1IDgxLjkxNzg5MzUsMTQ3LjYwNzg0NyA4MywxNDYuNDcwNDI5IEM3Ni40ODMwNjkxLDE0NC45MTMxMjMgNjkuNzQxOTg3NSwxNDQgNjMuMTU5OTA5NCwxNDQgTDYzLjE1OTkwOTQsMTQ0IFoiIGlkPSJGaWxsLTE3IiBmaWxsPSIjRjM2QzIxIj48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNNTEsMTQ1LjIxNDI1MyBDNTEsMTQ1LjIxNDI1MyA1NC44MDQ2MzUzLDE1MCA2NC4wMDA1NTI4LDE1MCBDNzMuMTk5Nzg3NCwxNTAgNzcsMTQ1LjIxNDI1MyA3NywxNDUuMjE0MjUzIEM3NywxNDUuMjE0MjUzIDYzLjUxMTg0MzUsMTQwLjIzMjE4NCA1MSwxNDUuMjE0MjUzIEw1MSwxNDUuMjE0MjUzIFoiIGlkPSJGaWxsLTE4IiBmaWxsPSIjRkRDQzlCIj48L3BhdGg+CiAgICAgICAgICAgICAgICA8cGF0aCBkPSJNMTA1LjE2NzUxOCwxNy4wNjU1NDU2IEMxMDUuMTY3NTE4LDE3LjA2NTU0NTYgODYuMzE4OTM1OSwtNy4wNDE4MjcyNyA0Ny4xMzY1MDgzLDIuMDQ4NzE5NjIgQzcuOTU1MTk0MSwxMS4xMzkyNjY1IDUuMDE5MTg0MTcsMzQuODUzMzQ2OSA2LjE5ODI2NDM4LDU3LjkyNDI2NTMgQzcuMzc2MjMxMiw4MC45OTk2NjU2IDEwLjUyMTU1ODUsOTQuNTMyOTYxIDE0LjIzNTgyODIsOTIuODYxMTg3OCBDMTcuOTQ4OTg0NSw5MS4xODk0MTQ2IDE2LjAxODM2NTksNzUuOTU3NDU0MiAxNi40MTAyNzkyLDcyLjQ0MjQ3MjYgQzE2LjgwMTA3OTIsNjguOTI3NDkxIDIwLjMzNzIwNjQsNTMuMDM1NTYxMiA0NC4xNjU5ODMzLDU1LjQwNzY0MTYgQzY3Ljk5NTg3MzYsNTcuNzc1MjM5OSA4OC4wNzU4NjU3LDQ5LjA4MjQ2NzUgODguMDc1ODY1Nyw0OS4wODI0Njc1IEM4OC4wNzU4NjU3LDQ5LjA4MjQ2NzUgOTYuMTYxMzA1Miw1OS4zNjQwOTY4IDEwMS40Nzk5Nyw2MS4zMzM5MjA3IEMxMTIuMTA5NTA2LDY1LjI2OTA4NjYgMTA4LjYyMTI1NSw5MS41MTc3MTg2IDExMi43ODc1NjEsOTEuNTE3NzE4NiBDMTE2Ljk0OTQxMyw5MS41MTc3MTg2IDEzNy4zNzM0NDMsMzAuOTAxMzczNyAxMDUuMTY3NTE4LDE3LjA2NTU0NTYgTDEwNS4xNjc1MTgsMTcuMDY1NTQ1NiBaIiBpZD0iRmlsbC0yMCIgZmlsbD0iIzUxMzYyQSI+PC9wYXRoPgogICAgICAgICAgICA8L2c+CiAgICAgICAgPC9nPgogICAgPC9nPgo8L3N2Zz4=" alt="">
                    <span>{{scope.FullName}}</span>
                </template>
                <template slot="Resources" slot-scope="scope">
                    <el-tag class="mar-right" size="mini" type="info" v-for="(item,index) in scope.Resources" :key="index">{{item}}</el-tag>
                </template>
            </JusTable>

            <!--创建用户-->
            <el-dialog :title="title" :visible.sync="buildDialogVisible">
                <el-form :model="form" :rules="formRules" ref="form" label-width="120px">
                    <div v-if="isBuild">
                        <el-form-item label="账号：" prop="account">
                            <el-input v-model="form.account"></el-input>
                        </el-form-item>
                        <el-form-item label="密码：" prop="password">
                            <el-input v-model="form.password"></el-input>
                        </el-form-item>
                        <el-form-item label="姓名：" prop="fullName">
                            <el-input v-model="form.fullName"></el-input>
                        </el-form-item>
                        <el-form-item label="手机：" prop="mobile">
                            <el-input v-model="form.mobile" placeholder="手机号确认后不能再修改"></el-input>
                        </el-form-item>
                    </div>

                    <div v-if="!isBuild">
                        <el-form-item label="姓名：">
                            <span>{{form.name}}</span>
                        </el-form-item>
                        <el-form-item label="当前角色：">
                            <span>{{form.roleName}}</span>
                        </el-form-item>
                    </div>

                    <el-form-item label="用户角色：" prop="roleId">
                        <el-select v-model="form.roleId">
                            <el-option v-for="item in roleOption" :key="item.id" :label="item.name" :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item v-if="isBuild" label="邮件：">
                        <el-input v-model="form.email"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="buildDialogVisible=false">取消</el-button>
                        <el-button @click="submit('form')">确定</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </JusLayout>
    </d2-container>
</template>
<script>
import { GetRoleList, AddSystemManage, GetAuthorityList, DelSystemManage, ResetPassword } from '@/api/platform.manage'
import util from '@/libs/util'
export default {
  name: 'userManage',
  data () {
    return {
      searchForm: {
        name: '',
        roleName: '',
        time: [],
        time1: '',
        time2: '',
        pageIndex: 1,
        pageSize: 10
      },
      pagination: {
        pageSize: 10,
        currentPage: 0,
        total: 0
      },
      labelWidth: '60px',
      column: [
        { title: '姓名', key: 'FullName', width: '150' },
        { title: '账号', key: 'Account', width: '100' },
        { title: '手机号', key: 'Mobile', width: '150' },
        { title: '角色', key: 'RoleName', width: '100' },
        { title: '已开通权限模块', key: 'Resources' },
        { title: '创建时间',
          key: 'CreatedAt',
          width: '200',
          formatter: (row, column, value, index) => {
            return this.$formatDate(value)
          }
        },
        { title: '邮箱', key: 'Email', width: '200' }
      ],
      rowHandle: {
        width: '250',
        custom: [
          { text: '重置密码', emit: 'reset' },
          { text: '变更角色', emit: 'changeRole' },
          { text: '删除', emit: 'del' }
        ]
      },
      data: [],

      buildDialogVisible: false,
      form: {
        id: '',
        account: '',
        password: '',
        fullName: '',
        mobile: '',
        roleId: '',
        email: '',
        roleName: '',
        authorityId: ''
      },
      formRules: {
        name: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        roleId: [{ required: true, message: '请选择角色', trigger: 'blur' }],
        mobile: [{ required: true, message: '请输入11位数手机号', trigger: 'blur' }],
        fullName: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
      },
      roleOption: [],
      title: '创建用户',
      isBuild: true
    }
  },
  methods: {
    handleRowDataChange (row) {

    },
    queryData (data, currentPage) {
      this.searchForm.pageIndex = currentPage.pageIndex
      this.searchForm.pageSize = currentPage.pageSize
      this.getDataList()
    },

    // 重置密码 - 密码忘记时
    reset ({ index, row }) {
      this.$confirm('确定要重置用户 ' + row.FullName + ' 的密码吗？', '', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 具体操作
        ResetPassword({ id: row.ID }).then((res) => {
          if (res.ret === 200) {
            this.$message.success('重置密码成功，密码为：123456')
          }
        })
      }).catch(() => {
        this.$message.warning('取消重置')
      })
    },

    // 创建时间范围
    getValue (val) {
      if (this.searchForm.time && this.searchForm.time.length > 0) {
        this.searchForm.time1 = util.formatDate(this.searchForm.time[0], 'yyyy-MM-dd')
        this.searchForm.time2 = util.formatDate(this.searchForm.time[1], 'yyyy-MM-dd')
      } else {
        this.searchForm.time1 = ''
        this.searchForm.time2 = ''
      }
    },

    // 变更角色
    changeRole ({ index, row }) {
      this.title = '变更用户角色'
      this.buildDialogVisible = true
      this.isBuild = false
      this.form = {
        id: row.ID,
        name: row.FullName,
        email: row.Email,
        roleId: row.RoleId,
        roleName: row.RoleName,
        authorityId: row.AuthorityId,
        mobile: row.Mobile
      }
    },

    // 删除
    del ({ index, row }) {
      this.$confirm('确定删除用户 ' + row.FullName + ' 吗？', '删除用户', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 具体操作
        DelSystemManage({ id: row.ID }).then((res) => {
          if (res.ret === 200) {
            this.$message.success('删除成功')
            this.getDataList()
          }
        })
      }).catch(() => {
        this.$message.warning('取消删除')
      })
    },

    // 创建用户
    buildUser () {
      this.title = '创建用户'
      this.buildDialogVisible = true
      this.isBuild = true
      this.form = {
        id: '',
        account: '',
        password: '',
        fullName: '',
        mobile: '',
        roleId: '',
        email: '',
        roleName: ''
      }
    },

    // 创建用户
    submit (form) {
      var reg = new RegExp('^[1][34578][0-9]{9}$')
      var emailReg = new RegExp('^[a-z0-9A-Z]+[- | a-z0-9A-Z . _]+@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-z]{2,}$')
      this.$refs[form].validate((valid) => {
        if (valid) {
          if (!reg.test(this.form.mobile)) {
            this.$message.warning('请输入正确的手机号格式')
            return
          }
          if (this.form.email != '') {
            if (!emailReg.test(this.form.email)) {
              this.$message.warning('请输入正确的邮箱格式')
              return
            }
          }

          AddSystemManage(this.form).then((res) => {
            if (res.ret === 200) {
              this.$message.success('保存成功')
              this.buildDialogVisible = false
              this.getDataList()
            }
          })
        }
      })
    },

    // 获取角色
    getRoleList () {
      GetRoleList().then((res) => {
        if (res.ret === 200) {
          this.roleOption = res.data
        }
      })
    },
    // 获取数据列表
    getDataList () {
      // 这是获取有权限的管理员
      GetAuthorityList(this.searchForm).then((res) => {
        if (res.ret === 200) {
          var data = res.data
          this.data = data.Data
          this.pagination = {
            currentPage: data.pageIndex,
            pageSize: data.pageSize,
            total: data.Count
          }
        }
      })
    },

    // 查询
    searchDataList () {
      this.getDataList()
    },
    // 重置
    reSearchDataList () {
      this.searchForm = {
        name: '',
        roleName: '',
        time: '',
        time1: '',
        time2: '',
        pageIndex: 1,
        pageSize: 10
      }
      this.getDataList()
    }
  },
  mounted () {
    this.getRoleList()
    this.getDataList()
  }
}
</script>
<style scoped>
    .mar_top {
        margin-top: 0px;
    }
    .mar-right {
        margin-right: 5px;
        margin-bottom: 5px;
    }
    .el-form .el-form-item {
        margin-bottom: 20px;
    }
    .header_img {
        height: 32px;
        width: 32px;
        vertical-align: middle;
        border-radius: 50%;
        margin-right: 10px;
    }
</style>
