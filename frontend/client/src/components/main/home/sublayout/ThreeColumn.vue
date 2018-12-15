<!-- Type: 1 -->

<template>
  <div class="sublayout-three-column">
    <div class="col-1-3">
      <home-layout
        :edit-mode="editMode"
        :data="getColumn1Data"/>
    </div>
    <div class="col-1-3">
      <home-layout
        :edit-mode="editMode"
        :data="getColumn2Data"/>
    </div>
    <div class="col-1-3">
      <home-layout
        :edit-mode="editMode"
        :data="getColumn3Data"/>
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
    },

    data: {
      type: Object,
      default () { return {} }
    }

  },

  computed: {

    getColumn1Data () {
      return this.data && this.data.children && this.data.children.length > 0 ? this.data.children[0] : []
    },

    getColumn2Data () {
      return this.data && this.data.children && this.data.children.length > 1 ? this.data.children[1] : []
    },

    getColumn3Data () {
      return this.data && this.data.children && this.data.children.length > 2 ? this.data.children[2] : []
    }

  }

})
</script>

<style lang="scss" scoped>
.sublayout-three-column {
  float: left;
  width: 100%;

  & > .col-1-3 {
    float: left;
    min-height: 1px;
    width: 33.33%;
  }

}
</style>
