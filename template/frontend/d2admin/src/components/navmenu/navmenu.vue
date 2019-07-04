<template>
    <div class="navMenu">
        <el-menu-item @click="btn(navMenu.id, navMenu.value)" v-for="navMenu in navMenus" :index="navMenu.id" v-if="navMenu.children.length == 0" :key="navMenu.id">
            <i :class="navMenu.icon"></i>
            {{navMenu.value}}
        </el-menu-item>
        <el-submenu v-for="navMenu in navMenus" :index="navMenu.id" v-if="navMenu.children.length != 0" :key="navMenu.id">
            <template slot="title"><i :class="navMenu.icon"></i><span>{{navMenu.value}}</span></template>
            <navMenu v-if="navMenu.children != 0" :navMenus="navMenu.children" @change="btn"></navMenu>
        </el-submenu>
    </div>
</template>

<script>
export default {
  props: {
    navMenus: Array
  },
  data () {
    return {}
  },
  methods: {
    btn (param1, param2) {
      this.$emit('change', param1, param2)
    }
  }
}
</script>
<style scoped>
    element.style{}
    .el-menu-item.is-active {
        background-color: #C6E2FF;
    }
</style>
