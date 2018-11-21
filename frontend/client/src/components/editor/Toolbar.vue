<template>
  <div class="editor-toolbar">
    <div class="row-one">
    </div>
    <div class="row-two">
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
  height: 100px;
  left: 0;
  position: absolute;
  right: 0;
  top: 27px;
  z-index: 1;

  & > .row-one {
    float: left;
    height: 50px;
    width: 100%;
  }

  & > .row-two {
    float: left;
    height: 50px;
    width: 100%;
  }

}
</style>
