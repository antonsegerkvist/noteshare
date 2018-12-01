<template>
  <div class="editor-toolbar">
    <div class="row-one">
      <div class="button">
      </div>
      <div class="button">
      </div>
    </div>
    <div class="row-two">
      <div class="button">
      </div>
      <div class="button">
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.editor-toolbar {
  background-color: #fff;
  height: 60px;
  left: 0;
  position: absolute;
  right: 0;
  top: 27px;
  z-index: 1;

  & > .row-one {
    float: left;
    height: 30px;
    width: 100%;
  }

  & > .row-two {
    float: left;
    height: 30px;
    width: 100%;
  }

  & > .row-one > .button, .row-two > .button {
    float: left;
    height: 100%;
    width: 30px;
  }

}
</style>
