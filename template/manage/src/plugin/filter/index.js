import util from '@/libs/util'

export default {
  install (Vue, options) {
    Vue.filter('formatDate', util.formatDate)
    // 元数据的审核状态
    Vue.filter('metaAuditStatus', function (val) {
      if (typeof (val) === 'number') {
        switch (val) {
          case 0:
            return '未审核'
          case 1:
            return '审核未通过'
          case 2:
            return '审核已通过'
          case 3:
            return '待提交审核'
          case 4:
            return '审核完成'
          default:
            return ''
        }
      }
      return ''
    })
    // 模型类型
    Vue.filter('moduleType', function (val) {
      switch (val) {
        case 'custom':
          return '自定义模组'
        case 'php':
          return 'php'
        case 'golang':
          return 'golang'
        case 'python':
          return 'python'
        case 'sql':
          return 'sql'
        case 'shell':
          return 'shell'
        default:
          return ''
      }
    })
    // 数据类型
    Vue.filter('dataType', function (val) {
      let data = util.dataType()
      for (let i = 0; i < data.length; i++) {
        if (data[i].value === val) {
          return data[i].label
        }
      }
      return ''
    })
    // 处理时间
    Vue.filter('displayTime', function (data) {
      if (data === '' || data === null || data === undefined) {
        return ''
      }
      let str = data
      // 将字符串转换成时间格式
      let timePublish = new Date(str).getTime()

      let timeNow = new Date().getTime()

      let minute = 1000 * 60
      let hour = minute * 60
      let day = hour * 24
      let month = day * 30
      let year = day * 365

      let diffValue = timeNow - timePublish
      let diffYear = diffValue / year
      let diffMonth = diffValue / month
      let diffWeek = diffValue / (7 * day)
      let diffDay = diffValue / day
      let diffHour = diffValue / hour
      let diffMinute = diffValue / minute

      let result
      if (diffValue < 0) {
        console.log('错误时间')
        return data
      } else if (diffYear >= 1) {
        result = parseInt(diffMonth) + '月前'
      } else if (diffMonth > 1) {
        result = parseInt(diffMonth) + '月前'
      } else if (diffWeek > 1) {
        result = parseInt(diffWeek) + '周前'
      } else if (diffDay > 1) {
        result = parseInt(diffDay) + '天前'
      } else if (diffHour > 1) {
        result = parseInt(diffHour) + '小时前'
      } else if (diffMinute > 1) {
        result = parseInt(diffMinute) + '分钟前'
      } else {
        result = '刚刚发表'
      }
      return result
    })
    // 是否为空
    Vue.filter('emptyText', function (val) {
      switch (val) {
        case 0:
          return '否'
        case 1:
          return '是'
        default:
          return ''
      }
    })
    // 是否主键
    Vue.filter('primaryKeyText', function (val) {
      switch (val) {
        case 0:
          return '否'
        case 1:
          return '是'
        default:
          return ''
      }
    })
  }
}
