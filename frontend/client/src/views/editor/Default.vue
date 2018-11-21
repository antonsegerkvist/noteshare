<template>
  <div class="default">
    <EditorNavigation/>
    <EditorToolbar/>
    <EditorEditor/>
  </div>
</template>

<script>
import Vue from 'vue'
import EditorEditor from '@/components/editor/Editor.vue'
import EditorNavigation from '@/components/editor/Navigation.vue'
import EditorToolbar from '@/components/editor/Toolbar.vue'

export default Vue.extend({

  components: {
    EditorEditor,
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
