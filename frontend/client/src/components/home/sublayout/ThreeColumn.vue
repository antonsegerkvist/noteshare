<!-- Type: 1 -->
<template>
  <div class="sublayout-three-column">
    <div class="col-1-3">
    </div>
    <div class="col-1-3">
    </div>
    <div class="col-1-3">
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  name: 'home-sublayout-three-column',

  props: {

    editMode: {
      type: Boolean,
      default () { return false }
    }

  }

})
</script>

<style lang="scss" scoped>
.sublayout-three-column {
  display: block;
  width: 100%;

  & > .col-1-3 {
    background-color: blue;
    float: left;
    min-height: 100px;
    width: 33.33%;
  }

}
</style>
