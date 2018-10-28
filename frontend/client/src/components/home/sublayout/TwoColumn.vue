<!-- Type: 0 -->
<template>
  <div class="sublayout-two-column">
    <div class="col-1-2">
      <home-layout
        :edit-mode="editMode"
        :data="getColumn1Data"/>
    </div>
    <div class="col-1-2">
      <home-layout
        :edit-mode="editMode"
        :data="getColumn2Data"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import HomeLayout from '@/components/home/Layout.vue'
export default Vue.extend({

  name: 'home-sublayout-two-column',

  components: {
    'home-layout': HomeLayout
  },

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
      return this.data && this.data.children && this.data.children.length > 0 && this.data.children ? this.data.children[0] : []
    },

    getColumn2Data () {
      return this.data && this.data.children && this.data.children.length > 1 && this.data.children ? this.data.children[1] : []
    }

  }

})
</script>

<style lang="scss" scoped>
.sublayout-two-column {
  display: block;
  width: 100%;

  & > .col-1-2 {
    background-color: red;
    float:left;
    min-height: 100px;
    width: 50%;
  }

}
</style>
