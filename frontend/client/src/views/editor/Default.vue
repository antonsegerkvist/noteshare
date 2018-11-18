<template>
  <div class="default">
    <EditorNavigation/>
    <EditorToolbar/>
  </div>
</template>

<script>
import Vue from 'vue'
import EditorNavigation from '@/components/editor/Navigation.vue'
import EditorToolbar from '@/components/editor/Toolbar.vue'

export default Vue.extend({

  components: {
    EditorNavigation,
    EditorToolbar
  }

})
</script>

<style lang="scss" scoped>
.default {
  height: 100%;
  position: relative;
  width: 100%;
}
</style>
