<template>
  <div class="view-subview-file">
    <div class="header">
    </div>
    <div class="content">
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({})
</script>

<style lang="scss" scoped>
.view-subview-file {
  background-color: #eee;
  height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
  width: 100%;

  & > .header {
    height: 50px;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;

    & > .row-1 {
      height: 25px;
      left: 0;
      position: absolute;
      right: 0;
      top: 0;
    }

    & > .row-2 {
      bottom: 0;
      height: 25px;
      left: 0;
      position: absolute;
      right: 0;
    }

  }

  & > .content {
    bottom: 0;
    left: 0;
    position: absolute;
    right: 0;
    top: 50px;
  }

}
</style>
