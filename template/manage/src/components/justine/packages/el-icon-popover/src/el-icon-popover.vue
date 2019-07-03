<template>
    <el-popover
            placement="bottom-start"
            width="200"
            v-model="visible2">
        <p>请选择菜单图标</p>
        <el-form>
            <el-form-item>
                <el-input class="input_width" size="mini" v-model="value" placeholder="查找图标" @input="searchIcon"></el-input>
            </el-form-item>
        </el-form>
        <div class="main_box">
            <div v-for="(item,index) in iconOption" :style="style['icon'+index]" :key="index" @click="choiceIcon(item, index)" class="box_">
                <span :class="item.icon"></span>
            </div>
        </div>
        <div slot="reference" class="icon_cursor">
            <span style="display: inline-block;width: 100%;height:100%;" :class="icon"></span>
        </div>
    </el-popover>
</template>
<script>
export default {
  name: 'el-icon-popover',
  props: {
    icon: ''
  },
  data () {
    return {
      visible2: false,
      iconOptions: [
        // 部分图标
        { icon: 'fa fa-cubes' },
        { icon: 'fa fa-cube' },
        { icon: 'fa fa-wrench' },
        { icon: 'fa fa-globe' },
        { icon: 'fa fa-eyedropper' },
        { icon: 'fa fa-cogs' },
        { icon: 'fa fa-database' },
        { icon: 'fa fa-file' },
        { icon: 'fa fa-file-archive-o' },
        { icon: 'fa fa-file-audio-o' },
        { icon: 'fa fa-file-code-o' },
        { icon: 'fa fa-file-excel-o' },
        { icon: 'fa fa-file-image-o' },
        { icon: 'fa fa-file-movie-o' },
        { icon: 'fa fa-file-o' },
        { icon: 'fa fa-file-pdf-o' },
        { icon: 'fa fa-file-photo-o' },
        { icon: 'fa fa-file-picture-o' },
        { icon: 'fa fa-file-powerpoint-o' },
        { icon: 'fa fa-file-sound-o' },
        { icon: 'fa fa-file-text' },
        { icon: 'fa fa-file-text-o' },
        { icon: 'fa fa-file-video-o' },
        { icon: 'fa fa-file-word-o' },
        { icon: 'fa fa-file-zip-o' },
        { icon: 'fa fa-folder' },
        { icon: 'fa fa-folder-o' },
        { icon: 'fa fa-folder-open' },
        { icon: 'fa fa-pencil' },
        { icon: 'fa fa-phone' },
        { icon: 'fa fa-mobile' },
        { icon: 'fa fa-photo' },
        { icon: 'fa fa-power-off' },
        { icon: 'fa fa-recycle' },
        { icon: 'fa fa-random' },
        { icon: 'fa fa-window-restore' },
        { icon: 'fa fa-mail-reply' },
        { icon: 'fa fa-mail-reply-all' },
        { icon: 'fa fa-quora' },
        { icon: 'fa fa-circle-o-notch' },
        { icon: 'fa fa-refresh' },
        { icon: 'fa fa-spinner' },
        { icon: 'fa fa-area-chart' },
        { icon: 'fa fa-bar-chart' },
        { icon: 'fa fa-bar-chart-o' },
        { icon: 'fa fa-line-chart' },
        { icon: 'fa fa-pie-chart' },
        { icon: 'fa fa-arrows-alt' },
        { icon: 'fa fa-backward' },
        { icon: 'fa fa-compress' },
        { icon: 'fa fa-eject' },
        { icon: 'fa fa-expand' },
        { icon: 'fa fa-fast-backward' },
        { icon: 'fa fa-fast-forward' },
        { icon: 'fa fa-forward' },
        { icon: 'fa fa-pause' },
        { icon: 'fa fa-pause-circle' },
        { icon: 'fa fa-pause-circle-o' },
        { icon: 'fa fa-play' },
        { icon: 'fa fa-play-circle' },
        { icon: 'fa fa-play-circle-o' },
        { icon: 'fa fa-random' },
        { icon: 'fa fa-step-backward' },
        { icon: 'fa fa-step-forward' },
        { icon: 'fa fa-stop' },
        { icon: 'fa fa-stop-circle' },
        { icon: 'fa fa-stop-circle-o' },
        { icon: 'fa fa-youtube-play' },
        { icon: 'fa fa-copy' },
        { icon: 'fa fa-cut' },
        { icon: 'fa fa-align-center' },
        { icon: 'fa fa-align-justify' },
        { icon: 'fa fa-align-left' },
        { icon: 'fa fa-align-right' },
        { icon: 'fa fa-bold' },
        { icon: 'fa fa-chain' },
        { icon: 'fa fa-chain-broken' },
        { icon: 'fa fa-clipboard' },
        { icon: 'fa fa-columns' },
        { icon: 'fa fa-dedent' },
        { icon: 'fa fa-eraser' },
        { icon: 'fa fa-file-text' },
        { icon: 'fa fa-floppy-o' },
        { icon: 'fa fa-font' },
        { icon: 'fa fa-header' },
        { icon: 'fa fa-indent' },
        { icon: 'fa fa-italic' },
        { icon: 'fa fa-link' },
        { icon: 'fa fa-list' },
        { icon: 'fa fa-list-alt' },
        { icon: 'fa fa-list-ol' },
        { icon: 'fa fa-list-ul' },
        { icon: 'fa fa-outdent' },
        { icon: 'fa fa-paperclip' },
        { icon: 'fa fa-paragraph' },
        { icon: 'fa fa-paste' },
        { icon: 'fa fa-repeat' },
        { icon: 'fa fa-rotate-left' },
        { icon: 'fa fa-rotate-right' },
        { icon: 'fa fa-save' },
        { icon: 'fa fa-scissors' },
        { icon: 'fa fa-strikethrough' },
        { icon: 'fa fa-subscript' },
        { icon: 'fa fa-superscript' },
        { icon: 'fa fa-table' },
        { icon: 'fa fa-text-height' },
        { icon: 'fa fa-text-width' },
        { icon: 'fa fa-th' },
        { icon: 'fa fa-th-large' },
        { icon: 'fa fa-th-list' },
        { icon: 'fa fa-underline' },
        { icon: 'fa fa-undo' },
        { icon: 'fa fa-unlink' },
        { icon: 'fa fa-angle-double-down' },
        { icon: 'fa fa-angle-double-left' },
        { icon: 'fa fa-angle-double-right' },
        { icon: 'fa fa-angle-double-up' },
        { icon: 'fa fa-angle-down' },
        { icon: 'fa fa-angle-left' },
        { icon: 'fa fa-angle-right' },
        { icon: 'fa fa-angle-up' },
        { icon: 'fa fa-arrow-circle-down' },
        { icon: 'fa fa-arrow-circle-left' },
        { icon: 'fa fa-arrow-circle-o-down' },
        { icon: 'fa fa-arrow-circle-o-left' },
        { icon: 'fa fa-arrow-circle-o-right' },
        { icon: 'fa fa-arrow-circle-o-up' },
        { icon: 'fa fa-arrow-circle-right' },
        { icon: 'fa fa-arrow-circle-up' },
        { icon: 'fa fa-arrow-down' },
        { icon: 'fa fa-arrow-left' },
        { icon: 'fa fa-arrow-right' },
        { icon: 'fa fa-arrow-up' },
        { icon: 'fa fa-arrows' },
        { icon: 'fa fa-arrows-alt' },
        { icon: 'fa fa-arrows-h' },
        { icon: 'fa fa-arrows-v' },
        { icon: 'fa fa-caret-down' },
        { icon: 'fa fa-caret-left' },
        { icon: 'fa fa-caret-right' },
        { icon: 'fa fa-caret-square-o-down' },
        { icon: 'fa fa-caret-square-o-left' },
        { icon: 'fa fa-caret-square-o-right' },
        { icon: 'fa fa-caret-square-o-up' },
        { icon: 'fa fa-caret-up' },
        { icon: 'fa fa-chevron-circle-down' },
        { icon: 'fa fa-chevron-circle-left' },
        { icon: 'fa fa-chevron-circle-right' },
        { icon: 'fa fa-chevron-circle-up' },
        { icon: 'fa fa-exchange' },
        { icon: 'fa fa-hand-o-down' },
        { icon: 'fa fa-hand-o-left' },
        { icon: 'fa fa-hand-o-right' },
        { icon: 'fa fa-hand-o-up' },
        { icon: 'fa fa-long-arrow-down' },
        { icon: 'fa fa-long-arrow-left' },
        { icon: 'fa fa-long-arrow-right' },
        { icon: 'fa fa-long-arrow-up' },
        { icon: 'fa fa-toggle-down' },
        { icon: 'fa fa-toggle-left' },
        { icon: 'fa fa-toggle-right' },
        { icon: 'fa fa-toggle-up' },
        { icon: 'fa fa-genderless' },
        { icon: 'fa fa-intersex' },
        { icon: 'fa fa-mars' },
        { icon: 'fa fa-mars-double' },
        { icon: 'fa fa-mars-stroke' },
        { icon: 'fa fa-mars-stroke-h' },
        { icon: 'fa fa-mars-stroke-v' },
        { icon: 'fa fa-mercury' },
        { icon: 'fa fa-neuter' },
        { icon: 'fa fa-transgender' },
        { icon: 'fa fa-transgender-alt' },
        { icon: 'fa fa-venus' },
        { icon: 'fa fa-venus-double' },
        { icon: 'fa fa-venus-mars' },
        { icon: 'fa fa-bitcoin' },
        { icon: 'fa fa-btc' },
        { icon: 'fa fa-cny' },
        { icon: 'fa fa-dollar' },
        { icon: 'fa fa-eur' },
        { icon: 'fa fa-euro' },
        { icon: 'fa fa-gbp' },
        { icon: 'fa fa-gg' },
        { icon: 'fa fa-gg-circle' },
        { icon: 'fa fa-ils' },
        { icon: 'fa fa-inr' },
        { icon: 'fa fa-jpy' },
        { icon: 'fa fa-krw' },
        { icon: 'fa fa-money' },
        { icon: 'fa fa-rmb' },
        { icon: 'fa fa-rouble' },
        { icon: 'fa fa-rub' },
        { icon: 'fa fa-ruble' },
        { icon: 'fa fa-rupee' },
        { icon: 'fa fa-shekel' },
        { icon: 'fa fa-sheqel' },
        { icon: 'fa fa-try' },
        { icon: 'fa fa-turkish-lira' },
        { icon: 'fa fa-usd' },
        { icon: 'fa fa-won' },
        { icon: 'fa fa-yen' },
        { icon: 'fa fa-hand-grab-o' },
        { icon: 'fa fa-hand-lizard-o' },
        { icon: 'fa fa-hand-o-down' },
        { icon: 'fa fa-hand-o-left' },
        { icon: 'fa fa-hand-o-right' },
        { icon: 'fa fa-hand-o-up' },
        { icon: 'fa fa-hand-paper-o' },
        { icon: 'fa fa-hand-peace-o' },
        { icon: 'fa fa-hand-pointer-o' },
        { icon: 'fa fa-hand-rock-o' },
        { icon: 'fa fa-hand-scissors-o' },
        { icon: 'fa fa-hand-spock-o' },
        { icon: 'fa fa-hand-stop-o' },
        { icon: 'fa fa-thumbs-down' },
        { icon: 'fa fa-thumbs-o-down' },
        { icon: 'fa fa-thumbs-o-up' },
        { icon: 'fa fa-thumbs-up' },

        { icon: 'el-icon-info' },
        { icon: 'el-icon-error' },
        { icon: 'el-icon-success' },
        { icon: 'el-icon-warning' },
        { icon: 'el-icon-question' },
        { icon: 'el-icon-back' },
        { icon: 'el-icon-arrow-left' },
        { icon: 'el-icon-arrow-down' },
        { icon: 'el-icon-arrow-right' },
        { icon: 'el-icon-arrow-up' },
        { icon: 'el-icon-caret-left' },
        { icon: 'el-icon-caret-bottom' },
        { icon: 'el-icon-caret-top' },
        { icon: 'el-icon-caret-right' },
        { icon: 'el-icon-d-arrow-left' },
        { icon: 'el-icon-d-arrow-right' },
        { icon: 'el-icon-minus' },
        { icon: 'el-icon-plus' },
        { icon: 'el-icon-remove' },
        { icon: 'el-icon-circle-plus' },
        { icon: 'el-icon-remove-outline' },
        { icon: 'el-icon-circle-plus-outline' },
        { icon: 'el-icon-close' },
        { icon: 'el-icon-check' },
        { icon: 'el-icon-circle-close' },
        { icon: 'el-icon-circle-check' },
        { icon: 'el-icon-circle-close-outline' },
        { icon: 'el-icon-circle-check-outline' },
        { icon: 'el-icon-zoom-out' },
        { icon: 'el-icon-zoom-in' },
        { icon: 'el-icon-d-caret' },
        { icon: 'el-icon-sort' },
        { icon: 'el-icon-sort-down' },
        { icon: 'el-icon-sort-up' },
        { icon: 'el-icon-tickets' },
        { icon: 'el-icon-document' },
        { icon: 'el-icon-sold-out' },
        { icon: 'el-icon-news' },
        { icon: 'el-icon-message' },
        { icon: 'el-icon-date' },
        { icon: 'el-icon-printer' },
        { icon: 'el-icon-time' },
        { icon: 'el-icon-bell' },
        { icon: 'el-icon-mobile-phone' },
        { icon: 'el-icon-service' },
        { icon: 'el-icon-view' },
        { icon: 'el-icon-menu' },
        { icon: 'el-icon-more' },
        { icon: 'el-icon-more-outline' },
        { icon: 'el-icon-star-on' },
        { icon: 'el-icon-star-off' },
        { icon: 'el-icon-location' },
        { icon: 'el-icon-location-outline' },
        { icon: 'el-icon-phone' },
        { icon: 'el-icon-phone-outline' },
        { icon: 'el-icon-picture' },
        { icon: 'el-icon-picture-outline' },
        { icon: 'el-icon-delete' },
        { icon: 'el-icon-search' },
        { icon: 'el-icon-edit' },
        { icon: 'el-icon-edit-outline' },
        { icon: 'el-icon-rank' },
        { icon: 'el-icon-refresh' },
        { icon: 'el-icon-share' },
        { icon: 'el-icon-setting' },
        { icon: 'el-icon-upload' },
        { icon: 'el-icon-upload2' },
        { icon: 'el-icon-download' },
        { icon: 'el-icon-loading' }
      ],
      iconOption: [],
      style: {
        icon0: ''
      },
      value: ''
    }
  },
  methods: {
    choiceIcon (item, index) {
      this.$emit('getIcon', item.icon)
      for (var key in this.style) {
        if (key == 'icon' + index) {
          this.style[key] = 'background-color: #F2F6FC;'
        } else {
          this.style[key] = ''
        }
      }
    },
    searchIcon () {
      if (this.value != '') {
        var arr = []
        for (var i in this.iconOptions) {
          var json = this.iconOptions[i]
          if (json.icon.indexOf(this.value) > -1) {
            arr.push({ icon: json.icon })
          }
        }
        this.iconOption = arr
      } else {
        this.iconOption = this.iconOptions
      }
    }
  },
  mounted () {
    this.iconOption = this.iconOptions
    for (var i in this.iconOption) {
      this.style['icon' + i] = ''
    }
  }
}
</script>
<style scoped>
    .main_box {
        max-height: 246px;
        overflow-y: auto;
        padding-top: 15px;
        border-top: 1px solid #e4e7ed;
    }
    .icon_cursor {
        cursor: pointer;
        padding-left: 20px;
        padding-right: 20px;
    }
    .box_ {
        width: 30px;
        height: 30px;
        display: inline-block;
        border:1px solid #e4e7ed;
        border-radius: 3px;
        margin-right: 10px;
        margin-bottom:10px;
        cursor: pointer;
        line-height: 30px;
        text-align: center;
    }
    .box_:hover {
        background-color: #F2F6FC;
    }
    .input_width {
        width:158px;
    }
</style>
