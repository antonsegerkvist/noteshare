<template>
  <div class="editor-navigation">
    <a href="javascript:void(0)" class="button">
      {{ getFileButtonTitle }}
    </a>
    <a href="javascript:void(0)" class="button">
      {{ getEditButtonTitle }}
    </a>
    <a href="javascript:void(0)" class="button">
      {{ getViewButtonTitle }}
    </a>
    <a href="javascript:void(0)" class="button">
      {{ getInsertButtonTitle }}
    </a>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  computed: {

    getFileButtonTitle () {
      return 'File'
    },

    getEditButtonTitle () {
      return 'Edit'
    },

    getViewButtonTitle () {
      return 'View'
    },

    getInsertButtonTitle () {
      return 'Insert'
    }

  }

})
</script>

<style lang="scss" scoped>
.editor-navigation {
  background-color: #333;
  height: 27px;
  left: 0;
  line-height: 27px;
  position: absolute;
  right: 0;
  top: 0;
  z-index: 1;

  & > .button {
    color: #eee;
    float: left;
    font-family: Arial, Helvetica, sans-serif;
    font-size: 13px;
    height: 100%;
    padding: 0 10px 0 10px;
    text-decoration: none;
    transition: background-color 100ms ease-in-out;
  }

  & > .button:hover {
    background-color: #444;
  }

}
</style>
