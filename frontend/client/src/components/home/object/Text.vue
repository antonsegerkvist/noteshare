<!-- Type: 10000 -->

<template>
  <div class="object-text">
    <global-editor-tinymce
      @save="save"/>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  name: 'home-object-text',

  methods: {

    save (content) {
      this.$emit('save', content)
    }

  }

})
</script>

<style lang="scss" scoped>
.object-text {
  float: left;
  width: 100%;
}
</style>
