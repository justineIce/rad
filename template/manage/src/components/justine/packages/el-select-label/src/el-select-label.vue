<template>
    <div>
        <el-select :multiple="multiple" v-if="!isCreate" size="mini" v-model="form.label" @change="setLabel">
            <el-option-group
                    v-for="group in labelOptions"
                    :key="group.label"
                    :label="group.label">
                <el-option
                        v-for="item in group.options"
                        :key="item.ID"
                        :label="item.Name"
                        :value="item.ID">
                    <div class="float_left">
                        <span class="float_left1" v-if="item.ID!='1-0'&&item.ID!='1-1'" :style="{backgroundColor: item.Color}"></span>
                        <span v-if="item.ID=='1-1'" class="el-icon-plus float_left1"></span>
                        <span class="float_left">{{ item.Name }}</span>
                    </div>
                </el-option>
            </el-option-group>
        </el-select>
        <div v-if="isCreate">
            <el-input disabled></el-input>
            <el-card class="box-card pos_">
                <div slot="header" class="clearfix">
                    <span>创建标签</span>
                </div>
                <el-form>
                    <el-form-item label="标签名称">
                        <el-input v-model="form.Name"></el-input>
                    </el-form-item>
                    <el-form-item label="颜色">
                        <br/>
                        <el-row>
                            <el-col v-for="(item,index) in ColorOption" :key="index" :span="3">
                                <div class="wIDth_height" :style="{backgroundColor: item.Color}" @click="choiceColor(item.Color)"></div>
                            </el-col>
                        </el-row>
                    </el-form-item>
                    <el-form-item>
                        <div>
                            <div class="el-btn">
                                <el-button title="随机生成一个颜色" class="el-icon-refresh" :style="{backgroundColor: form.Color}" @click="refresh()"></el-button>
                            </div>
                            <div class="dis_">
                                <el-input placeholder="请输入内容" v-model="Color" @input="setColor">
                                    <template slot="prepend">#</template>
                                </el-input>
                            </div>
                        </div>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="warning" @click="isCreate=false">取消</el-button>
                        <el-button @click="submit">创建</el-button>
                    </el-form-item>
                </el-form>
            </el-card>
        </div>
    </div>
</template>
<script>
import { GetTagList, AddTags } from '@/api/data.acquisition'
export default {
  Name: 'el-select-label',
  props: {
    // 标签数据[{ID: '', Name: '', Color: ''}]
    // 是否支持多选
    multiple: {
      type: Boolean,
      default: () => { return false }
    },
    // 默认选中的数据
    data: {
      type: Array,
      default: () => { return [] }
    }
  },
  watch: {
    options (val) {
      this.labelOptions[1].options = this.options
    },
    data (val) {
      this.form.label = this.data
    }
  },
  data () {
    return {
      isCreate: false,
      form: {
        label: '', // 选择标签的标志，可以是ID
        Name: '', // 创建标签的名称
        Color: '#1273DE' // 创建标签的颜色
      },
      Color: '1273DE',
      labelOptions: [
        {
          label: '清空',
          options: [
            { ID: '1-0', Name: '无标签', Color: '' }
          ]
        },
        {
          label: '已有标签',
          options: []
        },
        {
          label: '新建',
          options: [
            { ID: '1-1', Name: '创建标签', Color: '' }
          ]
        }
      ],

      // 颜色
      ColorOption: [
        { Color: '#b80000' },
        { Color: '#db3e00' },
        { Color: '#fccb00' },
        { Color: '#008b02' },
        { Color: '#006b76' },
        { Color: '#1273de' },
        { Color: '#004dcf' },
        { Color: '#5300eb' },
        { Color: '#eb9694' },
        { Color: '#fad0c3' },
        { Color: '#fef3bd' },
        { Color: '#c1e1c5' },
        { Color: '#bedadc' },
        { Color: '#c4def6' },
        { Color: '#bed3f3' },
        { Color: '#d4c4fb' }
      ],
      backgroundColor: 'background-Color:',
      param: ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f']
    }
  },
  methods: {
    setLabel () {
      if (!this.multiple) { // 单选
        var label = this.form.label
        if (label === '1-1') {
          // 创建标签
          this.isCreate = true
        }
        if (label === '1-1' || label === '1-0') {
          this.form.label = ''
        }

        if (label === '1-0') {
          this.form.label = ''
        } else if (label === '1-1') {
          // 创建标签
          this.isCreate = true
          this.form.label = ''
          this.form.Name = ''
        }
        this.$emit('getValue', this.form.label)
      } else { // 多选
        // 获取最后一个
        var lastLabel = this.form.label[this.form.label.length - 1]
        if (lastLabel === '1-1') {
          // 创建标签
          this.isCreate = true
          this.form.Name = ''
        }
        if (lastLabel === '1-1' || lastLabel === '1-0') {
          this.form.label = []
        }

        if (lastLabel === '1-0') {
          this.form.label = []
        } else if (lastLabel === '1-1') {
          // 创建标签
          this.isCreate = true
          this.form.label = []
        }
        this.$emit('getValue', this.form.label)
      }
    },
    getValue () {
      return this.form.label.join(',')
    },
    resetLabel (val) {
      this.form.label = val.map((obj) => { return obj * 1 })
    },
    operation (label) {
      if (label === '1-1') {
        // 创建标签
        this.isCreate = true
      }
      if (label === '1-1' || label === '1-0') {
        this.form.label = ''
      }

      if (label === '1-0') {
        this.form.label = ''
      } else if (label === '1-1') {
        // 创建标签
        this.isCreate = true
        this.form.label = ''
      }
      this.$emit('getValue', this.form.label)
    },

    // 颜色选择
    choiceColor (Color) {
      this.form.Color = Color
      this.Color = Color.substring(1)
    },

    // 随机生成颜色
    refresh () {
      var str = ''
      for (var i = 0; i <= 5; i++) {
        var s = this.param[Math.floor(Math.random() * (this.param.length))]
        str = str + s
      }
      this.Color = str
      this.form.Color = '#' + str
    },
    setColor () {
      this.form.Color = '#' + this.Color
    },
    // 创建
    submit () {
      if (this.form.Name.trim().length <= 0) {
        this.$message.warning('请输入标签名称')
        return
      } else if (this.Color.trim().length <= 0) {
        this.$message.warning('请选择颜色')
        return
      }
      this.isCreate = false
      // this.$emit('builtLabel', this.form.Name, this.form.Color)
      this.builtLabel(this.form.Name, this.form.Color)
    },
    // 标签创建
    builtLabel (name, color) {
      AddTags({ name: name, color: color }).then((res) => {
        if (res.ret === 200) {
          this.$message.success('创建成功')
          this.getTagList()
        }
      })
    },
    // 获取标签
    getTagList () {
      GetTagList().then((res) => {
        if (res.ret === 200) {
          this.labelOptions[1].options = res.data
          this.form.label = this.form.label.map((obj) => { return obj * 1 })
        }
      })
    }
  },
  created () {
    this.getTagList()
  },
  mounted () {
    this.labelOptions[1].options = this.options
    this.form.label = this.data
  }
}
</script>
<style scoped>
    .pos_ {
        position: absolute;
        z-index: 99;
    }
    .wIDth_height {
        height: 30px;
        cursor: pointer;
    }
    .dis_ {
        display: inline-block;
        float: right;
        wIDth: 70%;
    }
    .el-btn {
        display: inline-block;
        float: left;
        wIDth: 20%;
    }
    .label_ {
        display: inline-block;
        wIDth: 5px;
        height: 5px;
    }
    .float_left {
        float: left;
    }
    .float_left1 {
        float: left;
        display: inline-block;
        wIDth: 12px;
        height: 12px;
        border-radius: 3px;
        margin-top: 11px;
        margin-right: 5px;
    }
</style>
