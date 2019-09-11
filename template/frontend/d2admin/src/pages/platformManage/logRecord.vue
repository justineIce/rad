<!-- 日志审计 -->
<template>
    <d2-container class="meta-public">
        <JusLayout title="日志审计">
            <div>
                <el-row :gutter="20">
                    <el-form :model="searchForm" :label-width="labelWidth" label-position="left">
                        <el-col :span="6">
                            <el-form-item label="操作人：">
                                <el-input v-model="searchForm.name" placeholder="请输入操作人姓名"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="10">
                            <el-form-item label="操作时间：" label-width="80px">
                                <div class="block">
                                    <el-date-picker
                                            v-model="searchForm.operaTime"
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
                            <el-button @click="searchLog()">查询</el-button>
                            <el-button @click="reSearchLog">重置</el-button>
                        </el-col>
                    </el-form>
                </el-row>
            </div>
            <div>
                <!--数据列表-->
                <JusTable class="mT20"
                          :options="{'row-key':'ID'}"
                          :columns="column"
                          :data="data"
                          :pagination="pagination"
                          :showPagination="true"
                          @query-data="queryData"
                          expand>
                    <template slot="ResultStatus" slot-scope="scope">
                        <el-tag v-if="scope.ResultStatus=='200' || scope.ResultStatus=='404'" class="el-icon-check"
                                size="mini" type="success"> 成功
                        </el-tag>
                        <el-tag v-else-if="scope.ResultStatus=='401'" class="el-icon-s-release" size="mini" type="info">
                            授权失败
                        </el-tag>
                        <el-tag v-else-if="scope.ResultStatus=='403'" class="el-icon-remove-outline" size="mini" type="warning"> 权限不足
                        </el-tag>
                        <el-tag v-else class="el-icon-error" size="mini" type="danger"> 失败</el-tag>
                    </template>
                    <template slot="expand" slot-scope="props">
                        <el-form :data="props" label-position="left" inline class="demo-table-expand">
                            <el-form-item label="method">
                                <span>{{ props.Method }}</span>
                            </el-form-item>
                            <el-form-item label="提交参数">
                                <span>{{ props.Params }}</span>
                            </el-form-item>
                            <el-form-item label="UserAgent">
                                <span>{{ props.UserAgent }}</span>
                            </el-form-item>
                            <el-form-item label="IP地址">
                                <span>{{ props.RemoteAddr }}</span>
                            </el-form-item>
                            <el-form-item label="操作状态">
                                <span>{{ props.ResultStatus }}</span>
                            </el-form-item>
                            <el-form-item label="操作描述">
                                <span>{{ props.ResultMsg }}</span>
                            </el-form-item>
                        </el-form>
                    </template>

                </JusTable>
            </div>
        </JusLayout>
    </d2-container>
</template>
<script>
    import {GetAuthorityList, GetManageActionLog, GetTopMenuList} from '@/api/platform.manage'
    import util from '@/libs/util'

    export default {
        name: 'logRecord',
        data() {
            return {
                searchForm: {
                    resourceId: '',
                    systemManageId: '',
                    resultStatus: '',
                    operaTime: [],
                    time1: '',
                    time2: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                labelWidth: '80px',
                modelCategoryOption: [],
                column: [
                    {title: '标题', key: 'Title'},
                    {title: '请求地址', key: 'RequestURL'},
                    {
                        title: '操作时间',
                        key: 'CreatedAt',
                        formatter: (row, column, value, index) => {
                            return this.$formatDate(value)
                        }
                    },
                    {title: '操作人', key: 'Name'},
                    {title: '操作结果', key: 'ResultStatus'}
                ],
                data: [],
                pagination: {
                    currentPage: 1,
                    pageSize: 10,
                    total: 0
                }
            }
        },
        methods: {

            // 获取数据
            searchLog() {
                // 需将时间格式化
                // 时间范围
                this.getDataList()
            },
            // 重置
            reSearchLog() {
                this.searchForm = {
                    resourceId: '',
                    systemManageId: '',
                    resultStatus: '',
                    operaTime: [],
                    time1: '',
                    time2: '',
                    pageIndex: 1,
                    pageSize: 10
                }
                this.getDataList()
            },

            // 操作时间
            getValue(val) {
                if (this.searchForm.operaTime && this.searchForm.operaTime.length > 0) {
                    var date1 = this.searchForm.operaTime[0]
                    var date2 = this.searchForm.operaTime[1]
                    this.searchForm.time1 = util.formatDate(date1, 'yyyy-MM-dd')
                    this.searchForm.time2 = util.formatDate(date2, 'yyyy-MM-dd')
                } else {
                    this.searchForm.time1 = ''
                    this.searchForm.time2 = ''
                }
            },

            queryData(data, currentPage) {
                this.searchForm.pageIndex = currentPage.pageIndex
                this.searchForm.pageSize = currentPage.pageSize
                this.getDataList()
            },

            // 获取数据
            getDataList() {
                GetManageActionLog(this.searchForm).then((res) => {
                    if (res.ret === 200) {
                        var data = res.data
                        this.data = data.data
                        this.pagination = {
                            currentPage: data.PageIndex,
                            pageSize: data.PageSize,
                            total: data.Count
                        }
                    }
                })
            }
        },
        mounted() {
            this.getDataList()
        }
    }
</script>
<style scoped>
    .success {
        color: #67C23A;
    }

    .info {
        color: #909399;
    }

    .el-form .el-input__inner {
        width: 100%;
    }
</style>
<style>
    .demo-table-expand {
        font-size: 0;
    }

    .demo-table-expand label {
        width: 90px;
        color: #99a9bf !important;
        font-weight: bold;
    }

    .demo-table-expand .el-form-item {
        margin-right: 0;
        margin-bottom: 0;
        width: 100%;
    }

    .demo-table .demo-table-expand label {
        width: 90px;
        color: #99a9bf;
    }
</style>
