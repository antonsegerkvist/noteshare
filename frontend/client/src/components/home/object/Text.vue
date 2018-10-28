<!-- Type: 10000 -->

<template>
  <div class="object-text">
    <div v-if="editMode === false" v-html="data.content" class="view">
    </div>
    <div v-else class="edit">
      <global-editor-tinymce
        @save="save"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  name: 'home-object-text',

  props: {

    editMode: {
      type: Boolean,
      default () { return false }
    },

    data: {
      type: Object,
      default () { return {} }
    }

  },

  methods: {

    save (content) {
      this.$emit('save', content)
    }

  },

  computed: {

    getSafeContent () {
      return this.data && this.data.content ? this.data.content : ''
    }

  }

})
</script>

<style lang="scss" scoped>
.object-text {
  box-sizing: border-box;
  float: left;
  padding: 0 10px 20px 10px;
  width: 100%;

  & > .view {
    background-color: #fff;
    float: left;
    min-height: 100px;
    width: 100%;
  }

  & > .edit {
    background-color: #fff;
    float: left;
    min-height: 100px;
    width: 100%;
  }

}
</style>
