<!-- Type: 10000 -->

<template>
  <div class="object-text">
    <div class="view">
      <div v-html="data.header" class="header">
      </div>
      <div v-html="data.content" class="body">
      </div>
    </div>
    <div v-if="editMode && edit" class="modal">
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

  data () {
    return {
      edit: false
    }
  },

  methods: {

    save (content) {
      this.$emit('save', content)
    }

  },

  computed: {

    getSafeHeader () {
      return this.data && this.data.header ? this.data.header : ''
    },

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
    width: 100%;

    & > .header {
      background-color: #333;
      box-sizing: border-box;
      color: #eee;
      float: left;
      padding: 10px;
      width: 100%;
    }

    & > .body {
      box-sizing: border-box;
      float: left;
      padding: 10px;
      width: 100%;
    }

  }

  & > .modal {
    background-color: rgba(0,0,0,0.5);
    bottom: 0;
    left: 0;
    position: fixed;
    right: 0;
    top: 0;
    z-index: 1000000;
  }

}
</style>
