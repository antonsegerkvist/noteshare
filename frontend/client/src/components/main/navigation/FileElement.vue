<template>
  <ul class="main-navigation-file-element">
    <li class="folder">
    </li>
    <li class="children">
    </li>
  </ul>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  name: 'main-navigation-file-element',

  components: {
    'main-navigation-file-element': () => import('@/components/main/navigation/FileElement.vue')
  },

  props: {

    folderTree: {
      type: Object,
      default () { return {} }
    },

    folders: {
      type: Array,
      default () { return {} }
    }

  }

})
</script>

<style lang="scss" scoped>
.main-navigation-file-element {
  box-sizing: border-box;
  display: block;
  width: 100%;
}
</style>