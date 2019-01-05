<template>
  <div class="file-small-view">
  </div>
</template>

<script>
import Vue from 'vue'
import FileSmallElement from './SmallElement.vue'
export default Vue.extend({

  props: {

    files: {
      type: Array,
      default () { return [] }
    }

  },

  data () {
    return {
      elementsPerRow: 0
    }
  },

  mounted () {
    window.addEventListener('resize', this.resize)
  },

  beforeDestroy () {
    window.removeEventListener('resize', this.resize)
  },

  methods: {

    resize () {

    },

    update () {

    }

  }

})
</script>

<style lang="scss" scoped>
.file-small-view {
  height: 100%;
  position: relative;
  width: 100%;
}
</style>