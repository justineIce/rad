import request from '@/plugin/axios'

/* ---------------抽取源----------------- */

// 获取抽取源列表
export function GetExtractSourceList (data) {
  return request({
    url: '/acquisition/extract/source/list',
    method: 'post',
    data
  })
}

// 添加抽取源
export function AddExtractSource (data) {
  return request({
    url: '/acquisition/extract/source/add',
    method: 'post',
    data
  })
}

// 删除抽取源
export function DelExtractSource (data) {
  return request({
    url: '/acquisition/extract/source/del',
    method: 'post',
    data
  })
}

// 获取单条抽取源
export function GetExtractSource (data) {
  return request({
    url: '/acquisition/extract/source/get',
    method: 'post',
    data
  })
}

// 检测抽取源
export function TestContact (data) {
  return request({
    url: '/acquisition/extract/source/test',
    method: 'post',
    data
  })
}

/* --------------数据同步---------------- */

// 添加/编辑抽取目录
export function AddExtractDir (data) {
  return request({
    url: '/acquisition/sync/dir/add',
    method: 'post',
    data
  })
}

// 删除抽取目录
export function DelExtractDir (data) {
  return request({
    url: '/acquisition/sync/dir/del',
    method: 'post',
    data
  })
}

// 获取抽取目录列表
export function GetExtractDirList (data) {
  return request({
    url: '/acquisition/sync/dir/list',
    method: 'post',
    data
  })
}

// 获取集成任务列表
export function GetExtractTaskList (data) {
  return request({
    url: '/acquisition/sync/task/list',
    method: 'post',
    data
  })
}

// 删除集成任务
export function DelExtractTask (data) {
  return request({
    url: '/acquisition/sync/task/del',
    method: 'post',
    data
  })
}

// 数据同步任务 - 上线任务
export function AddTask (data) {
  return request({
    url: '/acquisition/sync/task/add',
    method: 'post',
    data
  })
}

// 数据同步任务 - 获取节点列表
export function GetNodeList (data) {
  return request({
    url: '/acquisition/node/list',
    method: 'post',
    data
  })
}

// 根据抽取任务id获取详情数据
export function GetExtractTaskById (data) {
  return request({
    url: '/acquisition/sync/task/get',
    method: 'post',
    data
  })
}

// 数据同步任务 - 获取标签列表
export function GetTagList (data) {
  return request({
    url: '/acquisition/tag/list',
    method: 'post',
    data
  })
}

// 创建标签
export function AddTags (data) {
  return request({
    url: '/acquisition/tag/add',
    method: 'post',
    data
  })
}

// 移动同步任务
export function TaskMoveTo (data) {
  return request({
    url: '/acquisition/sync/task/move',
    method: 'post',
    data
  })
}

/* ---------------单列抽取----------------- */

// 根据抽取源获取表
export function GetTableByExtractResource (data) {
  return request({
    url: '/acquisition/single/extract/resource/table/list',
    method: 'post',
    data
  })
}

// 根据表名和抽取源获取字段
export function GetExtractColumnByTable (data) {
  return request({
    url: '/acquisition/single/extract/resource/column/list',
    method: 'post',
    data
  })
}

// 获取存储源
export function GetStorageResourceList (data) {
  return request({
    url: '/acquisition/single/storage/database/list',
    method: 'post',
    data
  })
}

// 根据存储源获取表
export function GetTableByStorageResource (data) {
  return request({
    url: '/acquisition/single/storage/database/table/list',
    method: 'post',
    data
  })
}

// 根据存储源表和表名获取字段
export function GetStorageColumnByTable (data) {
  return request({
    url: '/acquisition/single/storage/database/column/list',
    method: 'post',
    data
  })
}

// 获取资源层信息
export function GetResourceLayer (data) {
  return request({
    url: '/acquisition/single/storage/database/resource/layer/v/list',
    method: 'post',
    data
  })
}

// 保存抽取任务
export function AddExtractTask (data) {
  return request({
    url: '/acquisition/single/extract/task/add',
    method: 'post',
    data
  })
}

/* -------------批量抽取----------------- */

// 获取抽取源下所有的表
export function GetAllExtractTable (data) {
  return request({
    url: '/acquisition/extract/sync/batch/table/list',
    method: 'post',
    data
  })
}

// 获取存储模型下所有的表
export function GetAllStorageTable (data) {
  return request({
    url: '/acquisition/storage/sync/batch/table/list',
    method: 'post',
    data
  })
}

// 添加/编辑批量抽取
export function AddBatchExtractTask (data) {
  return request({
    url: '/acquisition/extract/sync/batch/task/add',
    method: 'post',
    data
  })
}

// 根据集成任务id获取批量抽取详情
export function GetExtractBatchTaskById (data) {
  return request({
    url: '/acquisition/extract/sync/batch/task/get',
    method: 'post',
    data
  })
}

// 根据id获取选中的批量抽取源
export function GetBatchResourceById (data) {
  return request({
    url: '/acquisition/extract/sync/batch/resource/get',
    method: 'post',
    data
  })
}

// 根据id获取选中的批量抽取源
export function GetBatchStorageById (data) {
  return request({
    url: '/acquisition/extract/sync/batch/storage/get',
    method: 'post',
    data
  })
}

/* ------------------自定义采集------------------------ */

// 获取任务目录和其下任务
export function GetUserDefinedDirList (data) {
  return request({
    url: '/acquisition/extract/userDined/dir/list',
    method: 'post',
    data
  })
}

// 获取任务目录和其下任务
export function GetExtractRuleByTaskId (data) {
  return request({
    url: '/acquisition/extract/rule/by/task/get',
    method: 'post',
    data
  })
}

// 添加啊自定义规则
export function AddExtractUserDefinedTask (data) {
  return request({
    url: '/acquisition/extract/userDefined/rule/add',
    method: 'post',
    data
  })
}

// 获取任务
export function GetExtractTask (data) {
  return request({
    url: '/acquisition/extract/task/get',
    method: 'post',
    data
  })
}

/* -----------------fInput--------------------- */

// 查询表
export function GetTablesByStorageDatabase (data) {
  return request({
    url: '/acquisition/database/table/list',
    method: 'post',
    data
  })
}
export function CheckDatabaseWhere (data) {
  return request({
    url: '/acquisition/database/where/judge',
    method: 'post',
    data
  })
}
export function GetDatabasePreview (data) {
  return request({
    url: '/acquisition/database/data/preview',
    method: 'post',
    data
  })
}

/* -----------------上传对象---------------------- */

// 获取文件
export function GetFieldFolder (data) {
  return request({
    url: '/acquisition/extract/upload/field/list',
    method: 'post',
    data
  })
}
